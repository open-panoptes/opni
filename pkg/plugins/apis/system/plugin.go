package system

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"github.com/hashicorp/go-plugin"
	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	managementv1 "github.com/rancher/opni/pkg/apis/management/v1"
	"github.com/rancher/opni/pkg/caching"
	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/plugins"
	"github.com/rancher/opni/pkg/storage"
	"github.com/rancher/opni/pkg/util"
	"github.com/samber/lo"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SystemPluginClient interface {
	UseConfigAPI(configv1.GatewayConfigClient)
	UseManagementAPI(managementv1.ManagementClient)
	UseCachingProvider(caching.CachingProvider[proto.Message])
	UseKeyValueStore(KeyValueStoreClient)
	UseAPIExtensions(ExtensionClientInterface)
	mustEmbedUnimplementedSystemPluginClient()
}

// UnimplementedSystemPluginClient must be embedded to have forward compatible implementations.
type UnimplementedSystemPluginClient struct{}

func (UnimplementedSystemPluginClient) UseConfigAPI(configv1.GatewayConfigClient)                 {}
func (UnimplementedSystemPluginClient) UseManagementAPI(managementv1.ManagementClient)            {}
func (UnimplementedSystemPluginClient) UseKeyValueStore(KeyValueStoreClient)                      {}
func (UnimplementedSystemPluginClient) UseAPIExtensions(ExtensionClientInterface)                 {}
func (UnimplementedSystemPluginClient) UseCachingProvider(caching.CachingProvider[proto.Message]) {}
func (UnimplementedSystemPluginClient) mustEmbedUnimplementedSystemPluginClient()                 {}

type SystemPluginServer interface {
	ServeConfigAPI(configv1.GatewayConfigServer) error
	ServeManagementAPI(server managementv1.ManagementServer) error
	ServeKeyValueStore(namespace string, backend storage.Backend) error
	ServeCachingProvider() error
}

const (
	SystemPluginID  = "opni.System"
	KVServiceID     = "system.KeyValueStore"
	SystemServiceID = "system.System"
)

type systemPlugin struct {
	plugin.NetRPCUnsupportedPlugin

	client SystemPluginClient
}

var _ plugin.Plugin = (*systemPlugin)(nil)

func NewPlugin(client SystemPluginClient) plugin.Plugin {
	return &systemPlugin{
		client: client,
	}
}

// Plugin side

func (p *systemPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	RegisterSystemServer(s, &systemPluginClientImpl{
		broker: broker,
		server: s,
		client: p.client,
		cache:  caching.NewClientGrpcTtlCacher(),
	})
	return nil
}

type systemPluginClientImpl struct {
	broker *plugin.GRPCBroker
	server *grpc.Server
	client SystemPluginClient
	cache  caching.GrpcCachingInterceptor
}

func (c *systemPluginClientImpl) UseConfigAPI(ctx context.Context, in *BrokerID) (*emptypb.Empty, error) {
	cc, err := c.broker.Dial(in.Id)
	if err != nil {
		return nil, err
	}
	defer cc.Close()
	client := configv1.NewGatewayConfigClient(cc)
	ctx, ca := context.WithCancel(ctx)
	defer ca()
	clientHandler := lo.Async(func() error {
		return c.runAPIExtensionClientHandler(ctx, client)
	})
	clientRpc := lo.Async0(func() {
		c.client.UseConfigAPI(client)
	})

	select {
	case err := <-clientHandler:
		// don't wait for clientRpc to finish, the peer will be blocked on
		// a lower level context
		if err != nil {
			return nil, err
		}
	case <-clientRpc:
		ca()
		<-clientHandler
	}
	return &emptypb.Empty{}, nil
}

func (c *systemPluginClientImpl) UseManagementAPI(_ context.Context, in *BrokerID) (*emptypb.Empty, error) {
	cc, err := c.broker.Dial(
		in.Id,
		grpc.WithChainUnaryInterceptor(
			c.cache.UnaryClientInterceptor(),
		),
	)
	if err != nil {
		return nil, err
	}
	defer cc.Close()
	client := managementv1.NewManagementClient(cc)
	c.client.UseManagementAPI(client)
	return &emptypb.Empty{}, nil
}

func (c *systemPluginClientImpl) UseKeyValueStore(_ context.Context, in *BrokerID) (*emptypb.Empty, error) {
	cc, err := c.broker.Dial(in.Id)
	if err != nil {
		return nil, err
	}
	defer cc.Close()
	c.client.UseKeyValueStore(NewKeyValueStoreClient(cc))
	return &emptypb.Empty{}, nil
}

func (c *systemPluginClientImpl) runAPIExtensionClientHandler(ctx context.Context, client configv1.GatewayConfigClient) error {
	stream, err := client.WatchReactive(ctx, &corev1.ReactiveWatchRequest{
		Paths: []string{
			configv1.ProtoPath().Management().GrpcListenAddress()[1:].String()[1:],
		},
	})
	if err != nil {
		return err
	}
	for {
		events, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		addr := events.Items[0].Value.ToValue().String()

		dialOpts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
			grpc.WithContextDialer(util.DialProtocol),
			grpc.WithConnectParams(grpc.ConnectParams{
				Backoff: backoff.Config{
					BaseDelay:  1 * time.Millisecond,
					Multiplier: 2,
					Jitter:     0.2,
					MaxDelay:   1 * time.Second,
				},
			}),
			grpc.WithChainUnaryInterceptor(
				c.cache.UnaryClientInterceptor(),
			),
		}
		cc, err := grpc.DialContext(ctx, addr, dialOpts...)
		if err == nil {
			c.client.UseAPIExtensions(&apiExtensionInterfaceImpl{
				managementClientConn: cc,
			})
			cc.Close()
		}
	}
}

func (c *systemPluginClientImpl) UseCachingProvider(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	// no external caching provider served
	c.client.UseCachingProvider(c.cache)
	return &emptypb.Empty{}, nil
}

// Gateway side

func (p *systemPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	systemErr := plugins.CheckAvailability(ctx, c, SystemServiceID)
	kvErr := plugins.CheckAvailability(ctx, c, KVServiceID)
	if systemErr != nil && kvErr != nil {
		return nil, systemErr
	}

	return &systemPluginHandler{
		ctx:    ctx,
		broker: broker,
		client: NewSystemClient(c),
	}, nil
}

type systemPluginHandler struct {
	ctx    context.Context
	broker *plugin.GRPCBroker
	client SystemClient
}

var _ SystemPluginServer = (*systemPluginHandler)(nil)

func (s *systemPluginHandler) ServeConfigAPI(api configv1.GatewayConfigServer) error {
	return s.serveSystemApi(
		func(srv *grpc.Server) {
			configv1.RegisterGatewayConfigServer(srv, api)
		},
		func(id uint32) error {
			_, err := s.client.UseConfigAPI(s.ctx, &BrokerID{
				Id: id,
			})
			return err
		},
	)
}

func (s *systemPluginHandler) ServeManagementAPI(api managementv1.ManagementServer) error {
	return s.serveSystemApi(
		func(srv *grpc.Server) {
			managementv1.RegisterManagementServer(srv, api)
		},
		func(id uint32) error {
			_, err := s.client.UseManagementAPI(s.ctx, &BrokerID{
				Id: id,
			})
			return err
		},
	)
}

func (s *systemPluginHandler) ServeKeyValueStore(namespace string, backend storage.Backend) error {
	store := backend.KeyValueStore(namespace)
	var lockManager storage.LockManager
	if lmb, ok := backend.(storage.LockManagerBroker); ok {
		lockManager = lmb.LockManager(namespace)
	}

	kvStoreSrv := NewKVStoreServer(store, lockManager)
	return s.serveSystemApi(
		func(srv *grpc.Server) {
			RegisterKeyValueStoreServer(srv, kvStoreSrv)
		},
		func(id uint32) error {
			_, err := s.client.UseKeyValueStore(s.ctx, &BrokerID{
				Id: id,
			})
			return err
		},
	)
}

func (s *systemPluginHandler) ServeCachingProvider() error {
	_, err := s.client.UseCachingProvider(s.ctx, &emptypb.Empty{})
	return err
}

func init() {
	plugins.GatewayScheme.Add(SystemPluginID, NewPlugin(nil))
}

func (s *systemPluginHandler) serveSystemApi(regCallback func(*grpc.Server), useCallback func(uint32) error) error {
	id := s.broker.NextId()
	var srv *grpc.Server
	srvLock := make(chan struct{})
	once := sync.Once{}
	go s.broker.AcceptAndServe(id, func(so []grpc.ServerOption) *grpc.Server {
		so = append(so,
			grpc.StatsHandler(otelgrpc.NewServerHandler()),
		)
		srv = grpc.NewServer(so...)
		close(srvLock)
		go func() {
			<-s.ctx.Done()
			once.Do(srv.Stop)
		}()
		regCallback(srv)
		return srv
	})
	done := make(chan error)
	go func() {
		done <- useCallback(id)
	}()
	var err error
	select {
	case <-s.ctx.Done():
		err = s.ctx.Err()
	case err = <-done:
	}
	<-srvLock
	if srv != nil {
		once.Do(srv.Stop)
	}
	return err
}
