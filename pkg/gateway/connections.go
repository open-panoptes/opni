package gateway

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	sync "sync"
	"time"

	"github.com/google/uuid"
	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	"github.com/rancher/opni/pkg/auth/cluster"
	"github.com/rancher/opni/pkg/config/reactive"
	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/logger"
	"github.com/rancher/opni/pkg/storage"
	"github.com/rancher/opni/pkg/util"
	"github.com/rancher/opni/pkg/util/streams"
	"github.com/rancher/opni/pkg/versions"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type TrackedConnectionListener interface {
	// Called when a new agent connection to any gateway instance is tracked.
	// The provided context will be canceled when the tracked connection is deleted.
	// If a tracked connection is updated, this method will be called again with
	// the same context, agentId, and leaseId, but updated instanceInfo.
	// Implementations of this method MUST NOT block.
	HandleTrackedConnection(ctx context.Context, agentId string, holder string, instanceInfo *corev1.InstanceInfo)
}

const controllingInstanceAnnotation = "controlling-instance"

type activeTrackedConnection struct {
	agentId               string
	holder                string
	revision              int64
	timestamp             time.Time
	instanceInfo          *corev1.InstanceInfo
	trackingContext       context.Context
	cancelTrackingContext context.CancelFunc
}

// ConnectionTracker is a locally replicated view into the current state of
// all live gateway/agent connections. It can be used to determine which
// agents are currently connected, and to which gateway instance.
// All gateway instances track the same state and update in real time at the
// speed of the underlying storage backend.
type ConnectionTracker struct {
	rootContext context.Context
	mgr         *configv1.GatewayConfigManager
	logger      *slog.Logger

	connectionsKv storage.KeyValueStore
	connectionsLm storage.LockManager

	listenersMu   sync.Mutex
	connListeners []TrackedConnectionListener

	mu                sync.RWMutex
	activeConnections map[string]*activeTrackedConnection

	whoami string

	localInstanceInfoMu sync.Mutex
	localInstanceInfo   *corev1.InstanceInfo

	activeStreams sync.WaitGroup
}

func NewConnectionTracker(
	rootContext context.Context,
	mgr *configv1.GatewayConfigManager,
	connectionsKv storage.KeyValueStore,
	connectionsLm storage.LockManager,
	lg *slog.Logger,
) *ConnectionTracker {
	hostname, _ := os.Hostname()
	ct := &ConnectionTracker{
		localInstanceInfo: &corev1.InstanceInfo{
			Annotations: map[string]string{
				"hostname": hostname,
				"pid":      fmt.Sprint(os.Getpid()),
				"version":  versions.Version,
			},
		},
		rootContext:       rootContext,
		mgr:               mgr,
		connectionsKv:     connectionsKv,
		connectionsLm:     connectionsLm,
		logger:            lg,
		activeConnections: make(map[string]*activeTrackedConnection),
		whoami:            uuid.New().String(),
	}
	return ct
}

func NewReadOnlyConnectionTracker(
	rootContext context.Context,
	connectionsKv storage.KeyValueStore,
) *ConnectionTracker {
	return &ConnectionTracker{
		localInstanceInfo: nil,
		rootContext:       rootContext,
		connectionsKv:     connectionsKv,
		logger:            slog.New(slog.NewJSONHandler(io.Discard, nil)),
		activeConnections: make(map[string]*activeTrackedConnection),
	}
}

func (ct *ConnectionTracker) Lookup(agentId string) *corev1.InstanceInfo {
	ct.mu.RLock()
	defer ct.mu.RUnlock()
	if conn, ok := ct.activeConnections[agentId]; ok {
		return conn.instanceInfo
	}
	return nil
}

func (ct *ConnectionTracker) ListActiveConnections() []string {
	ct.mu.RLock()
	defer ct.mu.RUnlock()
	agents := make([]string, 0, len(ct.activeConnections))
	for agentId := range ct.activeConnections {
		agents = append(agents, agentId)
	}
	return agents
}

func (ct *ConnectionTracker) AddTrackedConnectionListener(listener TrackedConnectionListener) {
	ct.listenersMu.Lock()
	defer ct.listenersMu.Unlock()
	ct.mu.RLock()
	defer ct.mu.RUnlock()

	ct.connListeners = append(ct.connListeners, listener)
	for _, conn := range ct.activeConnections {
		listener.HandleTrackedConnection(conn.trackingContext, conn.agentId, conn.holder, conn.instanceInfo)
	}
}

func (ct *ConnectionTracker) LocalInstanceInfo() *corev1.InstanceInfo {
	ct.localInstanceInfoMu.Lock()
	defer ct.localInstanceInfoMu.Unlock()
	return util.ProtoClone(ct.localInstanceInfo)
}

// Starts the connection tracker. This will block until the context is canceled
// and the underlying kv store watcher is closed.
func (ct *ConnectionTracker) Run(ctx context.Context) error {
	// var wg sync.WaitGroup
	eg, ctx := errgroup.WithContext(ctx)

	reactive.Bind(ctx,
		func(v []protoreflect.Value) {
			ct.localInstanceInfoMu.Lock()
			defer ct.localInstanceInfoMu.Unlock()
			ct.localInstanceInfo.RelayAddress = os.ExpandEnv(v[0].String())
			ct.localInstanceInfo.ManagementAddress = os.ExpandEnv(v[1].String())
			ct.localInstanceInfo.GatewayAddress = os.ExpandEnv(v[2].String())
			ct.localInstanceInfo.WebAddress = os.ExpandEnv(v[3].String())
			go ct.updateInstanceInfo(util.Must(proto.Marshal(ct.localInstanceInfo)))
		},
		ct.mgr.Reactive(configv1.ProtoPath().Relay().AdvertiseAddress()),
		ct.mgr.Reactive(configv1.ProtoPath().Management().AdvertiseAddress()),
		ct.mgr.Reactive(configv1.ProtoPath().Server().AdvertiseAddress()),
		ct.mgr.Reactive(configv1.ProtoPath().Dashboard().AdvertiseAddress()),
	)

	connectionsWatcher, err := ct.connectionsKv.Watch(ctx, "", storage.WithPrefix(), storage.WithRevision(0))
	if err != nil {
		return err
	}

	eg.Go(func() error {
		for event := range connectionsWatcher {
			ct.mu.Lock()
			ct.handleConnEvent(event)
			ct.mu.Unlock()
		}
		return nil
	})

	err = eg.Wait()
	ct.logger.Info("releasing all active connection locks")
	// release all active connections
	// ct.mu.Lock()
	// for _, conn := range ct.activeConnections {
	// 	if err := ct.connectionsKv.Delete(conn.trackingContext, ct.connKey(conn.agentId)); err != nil {
	// 		ct.logger.With(
	// 			logger.Err(err),
	// 			"agentId", conn.agentId,
	// 		).Error("failed to release connection lock")
	// 	}
	// 	conn.cancelTrackingContext()
	// }
	// ct.mu.Unlock()
	ct.activeStreams.Wait()
	return err
}

func (ct *ConnectionTracker) updateInstanceInfo(instanceInfoData []byte) {
	ct.mu.RLock()
	defer ct.mu.RUnlock()
	numConnections := len(ct.activeConnections)
	ct.logger.With(
		"activeConnections", numConnections,
	).Debug("updating instance info")

	for agentId, conn := range ct.activeConnections {
		if err := ct.connectionsKv.Put(conn.trackingContext, ct.connKey(agentId), instanceInfoData); err != nil {
			ct.logger.With(
				logger.Err(err),
				"agentId", agentId,
			).Warn("failed to update instance info")
		}
	}
}

type instanceInfoKeyType struct{}

var instanceInfoKey = instanceInfoKeyType{}

func (ct *ConnectionTracker) connKey(agentId string) string {
	return agentId
}

func (ct *ConnectionTracker) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ct.activeStreams.Add(1)
		defer ct.activeStreams.Done()

		agentId := cluster.StreamAuthorizedID(ss.Context())
		instanceInfo := ct.LocalInstanceInfo()
		key := agentId
		lock := ct.connectionsLm.NewLock(key)
		ct.logger.With("agentId", agentId).Debug("attempting to acquire connection lock")
		acquired, done, err := lock.TryLock(ss.Context())
		if !acquired {
			ct.logger.With(
				logger.Err(err),
				"agentId", agentId,
			).Error("failed to acquire lock on agent connection")
			if err == nil {
				return status.Errorf(codes.FailedPrecondition, "already connected to a different gateway instance")
			}
			return status.Errorf(codes.Internal, "failed to acquire lock on agent connection: %v", err)
		}
		ct.logger.With("agentId", agentId).Debug("acquired agent connection lock")

		defer func() {
			ct.logger.With("agentId", agentId).Debug("releasing agent connection lock")
			if err := lock.Unlock(); err != nil {
				ct.logger.With(
					logger.Err(err),
					"agentId", agentId,
				).Error("failed to release lock on agent connection")
			}
			<-done
		}()

		instanceInfo.Acquired = true
		instanceInfo.Annotations = map[string]string{
			controllingInstanceAnnotation: ct.whoami,
		}
		if err := ct.connectionsKv.Put(ss.Context(), ct.connKey(agentId), util.Must(proto.Marshal(instanceInfo))); err != nil {
			ct.logger.Warn("failed to persist instance info in the connections KV")
		}

		stream := streams.NewServerStreamWithContext(ss)
		stream.Ctx = context.WithValue(stream.Ctx, instanceInfoKey, util.ProtoClone(instanceInfo))
		return handler(srv, stream)
	}
}

// func decodeConnKey(key string) (agentId string, leaseId string, ok bool) {
// 	parts := strings.Split(key, "/")
// 	if len(parts) != 2 {
// 		// only handle keys of the form <agentId>/<leaseId>
// 		return "", "", false
// 	}

// 	agentId = parts[0]
// 	leaseId = parts[1]
// 	ok = true
// 	return
// }

func (ct *ConnectionTracker) handleConnUpdate(
	event storage.WatchEvent[storage.KeyRevision[[]byte]],
	agentId string,
	conn *activeTrackedConnection,
	holder string,
	info *corev1.InstanceInfo,
) (invalidated bool) {
	key := event.Current.Key()
	lg := ct.logger.With("key", key, "agentId", agentId)
	if !info.GetAcquired() {
		// a different instance is only attempting to acquire the lock,
		// ignore the event
		lg.With("instance", info.GetRelayAddress()).Debug("tracked connection is still being acquired...")
		return false
	}
	if conn.holder == holder {
		// update revision and instance info, and notify listeners
		if holder != ct.connKey(agentId) {
			// don't log this for local instances, it would be too noisy
			ct.logger.With("agentId", agentId, "revision", event.Current.Revision()).Debug("tracked connection updated")
		}
		conn.revision = event.Current.Revision()
		conn.instanceInfo = info
		for _, listener := range ct.connListeners {
			listener.HandleTrackedConnection(conn.trackingContext, agentId, holder, info)
		}
		return false
	}
	// a different instance has acquired the lock, invalidate
	// the current tracked connection
	lg.Debug("tracked connection invalidated")
	conn.cancelTrackingContext()
	delete(ct.activeConnections, agentId)
	return true
}

func (ct *ConnectionTracker) handleConnCreate(
	event storage.WatchEvent[storage.KeyRevision[[]byte]],
	agentId, holder string,
	info *corev1.InstanceInfo,
) {
	if !info.GetAcquired() {
		return // ignore unacquired connections
	}
	if ct.IsLocalInstance(info) {
		ct.logger.With("agentId", agentId).Debug("tracking new connection (local)")
	} else {
		ct.logger.With("agentId", agentId).Debug("tracking new connection")
	}
	ctx, cancel := context.WithCancel(ct.rootContext)
	conn := &activeTrackedConnection{
		agentId:               agentId,
		holder:                holder,
		revision:              event.Current.Revision(),
		instanceInfo:          info,
		trackingContext:       ctx,
		cancelTrackingContext: cancel,
	}
	ct.activeConnections[agentId] = conn
	for _, listener := range ct.connListeners {
		listener.HandleTrackedConnection(ctx, agentId, holder, info)
	}
}

func (ct *ConnectionTracker) handleConnEvent(event storage.WatchEvent[storage.KeyRevision[[]byte]]) {
	agentId := event.Current.Key()
	if strings.Contains(agentId, "/") {
		return // todo: hack to ignore old-style connection keys
	}
	switch event.EventType {
	case storage.WatchEventPut:
		ct.listenersMu.Lock()
		defer ct.listenersMu.Unlock()
		lg := ct.logger.With(
			"agentId", agentId,
			"key", event.Current.Key(),
		)
		instanceInfo := sync.OnceValues(func() (*corev1.InstanceInfo, error) {
			var info corev1.InstanceInfo
			if err := proto.Unmarshal(event.Current.Value(), &info); err != nil {
				return nil, err
			}
			return &info, nil
		})
		if conn, ok := ct.activeConnections[agentId]; ok {
			info, err := instanceInfo()
			if err != nil {
				lg.With(logger.Err(err)).Error("failed to unmarshal instance info")
				return
			}
			var holder string
			if info.Annotations != nil {
				holder = info.Annotations[controllingInstanceAnnotation]
			}
			invalidated := ct.handleConnUpdate(
				event,
				agentId,
				conn,
				holder,
				info,
			)
			if !invalidated {
				return
			}
		}
		info, err := instanceInfo()
		if err != nil {
			lg.With(logger.Err(err)).Error("failed to unmarshal instance info")
			return
		}
		var holder string
		if info.Annotations != nil {
			holder = info.Annotations[controllingInstanceAnnotation]
		}
		ct.handleConnCreate(
			event,
			agentId,
			holder,
			info,
		)
	case storage.WatchEventDelete:
		// make sure the previous revision of the deleted key is the same as the
		// revision of the tracked connection.
		lg := ct.logger.With("key", event.Previous.Key(), "rev", event.Previous.Revision())
		if conn, ok := ct.activeConnections[agentId]; ok {
			prev := &corev1.InstanceInfo{}
			if err := proto.Unmarshal(event.Previous.Value(), prev); err == nil {
				if ct.IsLocalInstance(prev) {
					lg.Debug("tracked connection deleted (local)")
				} else {
					lg.With(
						"prevAddress", prev.GetRelayAddress(),
					).Debug("tracked connection deleted")
				}
			} else {
				lg.With(
					"prevAddress", "(unknown)",
				).Debug("tracked connection deleted")
			}
			conn.cancelTrackingContext()
			delete(ct.activeConnections, agentId)
		} else {
			lg.With("key", event.Previous.Key()).Debug("ignoring untracked key deletion event")
		}
	}
}

func (ct *ConnectionTracker) IsLocalInstance(instanceInfo *corev1.InstanceInfo) bool {
	if ct.localInstanceInfo == nil || instanceInfo.Annotations == nil {
		return false
	}

	// Note: this assumes that if the relay address is the same, then the
	// management address is also the same. If we ever decide to allow
	// standalone management servers, this will need to be updated.
	return instanceInfo.Annotations[controllingInstanceAnnotation] == ct.whoami
}

func (ct *ConnectionTracker) IsTrackedLocal(agentId string) bool {
	if ct.localInstanceInfo == nil {
		return false
	}

	ct.mu.RLock()
	defer ct.mu.RUnlock()
	info, ok := ct.activeConnections[agentId]
	if ok && ct.IsLocalInstance(info.instanceInfo) {
		return true
	}
	return false
}

func (ct *ConnectionTracker) IsTrackedRemote(agentId string) bool {
	ct.mu.RLock()
	defer ct.mu.RUnlock()
	info, ok := ct.activeConnections[agentId]
	if ok && !ct.IsLocalInstance(info.instanceInfo) {
		return true
	}
	return false
}
