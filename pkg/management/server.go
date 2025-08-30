package management

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jhump/protoreflect/desc"
	capabilityv1 "github.com/rancher/opni/pkg/apis/capability/v1"
	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	managementv1 "github.com/rancher/opni/pkg/apis/management/v1"
	"github.com/rancher/opni/pkg/auth/local"
	authutil "github.com/rancher/opni/pkg/auth/util"
	"github.com/rancher/opni/pkg/caching"
	"github.com/rancher/opni/pkg/capabilities"
	"github.com/rancher/opni/pkg/config"
	"github.com/rancher/opni/pkg/config/reactive"
	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/health"
	"github.com/rancher/opni/pkg/logger"
	"github.com/rancher/opni/pkg/pkp"
	"github.com/rancher/opni/pkg/plugins"
	"github.com/rancher/opni/pkg/plugins/apis/apiextensions"
	"github.com/rancher/opni/pkg/plugins/hooks"
	"github.com/rancher/opni/pkg/plugins/meta"
	"github.com/rancher/opni/pkg/plugins/types"
	"github.com/rancher/opni/pkg/storage"
	"github.com/rancher/opni/pkg/util"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	channelzservice "google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CoreDataSource provides a way to obtain data which the management
// server needs to serve its core API
type CoreDataSource interface {
	StorageBackend() storage.Backend
	TLSConfig() *tls.Config
}

// CapabilitiesDataSource provides a way to obtain data which the management
// server needs to serve capabilities-related endpoints
type CapabilitiesDataSource interface {
	capabilityv1.BackendServer
}

type HealthStatusDataSource interface {
	GetClusterHealthStatus(ref *corev1.Reference) (*corev1.HealthStatus, error)
	WatchClusterHealthStatus(ctx context.Context) <-chan *corev1.ClusterHealthStatus
}

type apiExtension struct {
	client      apiextensions.ManagementAPIExtensionClient
	clientConn  *grpc.ClientConn
	status      *health.ServingStatus
	serviceDesc *desc.ServiceDescriptor
	httpRules   []*managementv1.HTTPRuleDescriptor
}

type Server struct {
	managementv1.UnsafeManagementServer
	managementv1.UnimplementedLocalPasswordServer
	managementServerOptions
	mgr               *configv1.GatewayConfigManager
	rbacManagerStore  capabilities.RBACManagerStore
	logger            *slog.Logger
	coreDataSource    CoreDataSource
	dashboardSettings *DashboardSettingsManager
	localAuth         local.LocalAuthenticator
	director          StreamDirector
	pluginsLoaded     chan struct{}

	apiExtMu      sync.RWMutex
	apiExtensions []apiExtension
}

var _ managementv1.ManagementServer = (*Server)(nil)

type managementServerOptions struct {
	lifecycler             config.Lifecycler
	capabilitiesDataSource CapabilitiesDataSource
	healthStatusDataSource HealthStatusDataSource
}

type ManagementServerOption func(*managementServerOptions)

func (o *managementServerOptions) apply(opts ...ManagementServerOption) {
	for _, op := range opts {
		op(o)
	}
}

func WithCapabilitiesDataSource(src CapabilitiesDataSource) ManagementServerOption {
	return func(o *managementServerOptions) {
		o.capabilitiesDataSource = src
	}
}

func WithHealthStatusDataSource(src HealthStatusDataSource) ManagementServerOption {
	return func(o *managementServerOptions) {
		o.healthStatusDataSource = src
	}
}

func NewServer(
	ctx context.Context,
	cds CoreDataSource,
	mgr *configv1.GatewayConfigManager,
	pluginLoader plugins.LoaderInterface,
	opts ...ManagementServerOption,
) *Server {
	lg := logger.New().WithGroup("mgmt")
	options := managementServerOptions{}
	options.apply(opts...)

	m := &Server{
		managementServerOptions: options,
		mgr:                     mgr,
		logger:                  lg,
		coreDataSource:          cds,
		rbacManagerStore:        capabilities.NewRBACManagerStore(lg),
		dashboardSettings: &DashboardSettingsManager{
			kv:     cds.StorageBackend().KeyValueStore("dashboard"),
			logger: lg,
		},
		pluginsLoaded: make(chan struct{}),
	}
	m.director = m.configureApiExtensionDirector(ctx, pluginLoader)

	pluginLoader.Hook(hooks.OnLoadM(func(sp types.SystemPlugin, md meta.PluginMeta) {
		go func() {
			if err := sp.ServeConfigAPI(m.mgr.AsServer()); err != nil {
				lg.With(
					"plugin", md.Module,
					logger.Err(err),
				).Error("failed to serve config API to plugin")
			}
		}()
		go func() {
			if err := sp.ServeManagementAPI(m); err != nil {
				lg.With(
					"plugin", md.Module,
					logger.Err(err),
				).Error("failed to serve management API to plugin")
			}
		}()
	}))

	pluginLoader.Hook(hooks.OnLoadM(func(p types.CapabilityRBACPlugin, md meta.PluginMeta) {
		list, err := p.List(ctx, &emptypb.Empty{})
		if err != nil {
			lg.With(
				"plugin", md.Module,
				logger.Err(err),
			).Error("failed to list capabilities")
			return
		}
		for _, cap := range list.GetItems() {
			info, err := p.Info(ctx, &corev1.Reference{Id: cap.GetName()})
			if err != nil {
				lg.With(
					"plugin", md.Module,
				).Error("failed to get capability info")
				return
			}
			if err := m.rbacManagerStore.Add(info.Name, p); err != nil {
				lg.With(
					"plugin", md.Module,
					logger.Err(err),
				).Error("failed to add capability backend rbac")
			}
			lg.With(
				"plugin", md.Module,
				"capability", info.Name,
			).Info("added capability rbac backend")
		}
	}))

	pluginLoader.Hook(hooks.OnLoadingCompleted(func(int) {
		close(m.pluginsLoaded)
	}))

	m.localAuth = local.NewLocalAuthenticator(m.coreDataSource.StorageBackend().KeyValueStore(authutil.AuthNamespace))

	return m
}

type managementApiServer interface {
	ServeManagementAPI(managementv1.ManagementServer)
}

func (m *Server) ListenAndServe(ctx context.Context) error {
	eg := errgroup.Group{}
	eg.Go(func() error {
		err := m.listenAndServeGrpc(ctx)
		if err != nil {
			if !errors.Is(err, context.Canceled) {
				return fmt.Errorf("management grpc server exited with error: %w", err)
			}
		}
		m.logger.Info("management grpc server stopped")
		return nil
	})
	eg.Go(func() error {
		err := m.listenAndServeHttp(ctx)
		if err != nil {
			if !errors.Is(err, context.Canceled) {
				return fmt.Errorf("management http server exited with error: %w", err)
			}
		}
		m.logger.Info("management http server stopped")
		return nil
	})

	return eg.Wait()
}

func (m *Server) listenAndServeGrpc(ctx context.Context) error {
	grpcListenAddr := m.mgr.Reactive(configv1.ProtoPath().Management().GrpcListenAddress())
	var mu sync.Mutex
	var serveError error
	var done chan struct{}
	var server *grpc.Server
	grpcListenAddr.WatchFunc(ctx, func(v protoreflect.Value) {
		mu.Lock()
		defer mu.Unlock()

		if server != nil {
			server.Stop()
		}

		server = grpc.NewServer(
			grpc.Creds(insecure.NewCredentials()),
			grpc.UnknownServiceHandler(unknownServiceHandler(m.director)),
			grpc.StatsHandler(otelgrpc.NewServerHandler()),
			grpc.ChainUnaryInterceptor(caching.NewClientGrpcTtlCacher().UnaryServerInterceptor()),
		)
		managementv1.RegisterManagementServer(server, m)
		configv1.RegisterGatewayConfigServer(server, m.mgr.AsServer())
		managementv1.RegisterLocalPasswordServer(server, m)
		channelzservice.RegisterChannelzServiceToServer(server)

		addr := v.String()
		listener, err := util.NewProtocolListener(addr)
		if err != nil {
			m.logger.With(
				"address", addr,
				logger.Err(err),
			).Error("failed to start management gRPC server")
			return
		}
		m.logger.With(
			"address", listener.Addr().String(),
		).Info("management gRPC server starting")
		done = make(chan struct{})
		go func(done chan struct{}) {
			err := server.Serve(listener)
			mu.Lock()
			serveError = err
			close(done)
			mu.Unlock()
		}(done)
	})
	<-ctx.Done()
	mu.Lock()
	if server != nil {
		server.Stop()
	}
	mu.Unlock()
	d := done
	if d != nil {
		<-d
	}
	return serveError
}

func (m *Server) listenAndServeHttp(ctx context.Context) error {
	httpListenAddr := m.mgr.Reactive(configv1.ProtoPath().Management().HttpListenAddress())
	grpcListenAddr := m.mgr.Reactive(configv1.ProtoPath().Management().GrpcListenAddress())

	var mu sync.Mutex
	var stopServer func()

	reactive.Bind(ctx, func(v []protoreflect.Value) {
		mu.Lock()
		defer mu.Unlock()

		if stopServer != nil {
			stopServer()
		}

		httpAddr := v[0].String()
		grpcAddr := v[1].String()

		listener, err := util.NewProtocolListener(httpAddr)
		if err != nil {
			m.logger.With(
				"address", httpAddr,
				logger.Err(err),
			).Error("failed to start management HTTP server")
			return
		}
		m.logger.With(
			"address", listener.Addr().String(),
		).Info("management HTTP server starting")
		gwmux := runtime.NewServeMux(
			runtime.WithMarshalerOption("application/json", &LegacyJsonMarshaler{}),
			runtime.WithMarshalerOption("application/octet-stream", &DynamicV1Marshaler{}),
			runtime.WithMarshalerOption(runtime.MIMEWildcard, &DynamicV1Marshaler{}),
		)
		router := gin.New()
		server := &http.Server{
			Addr: httpAddr,
			BaseContext: func(net.Listener) context.Context {
				return ctx
			},
			Handler: router.Handler(),
		}
		stopServer = sync.OnceFunc(func() {
			server.Close()
		})
		router.Any("/*any", gin.WrapF(gwmux.ServeHTTP))
		m.configureManagementHttpApi(ctx, server, grpcAddr, gwmux)
		go func(server *http.Server) {
			// important: wait for plugins that may have api extensions to load,
			// otherwise the server will be started before adding them to the router
			select {
			case <-m.pluginsLoaded:
			case <-ctx.Done():
			}
			m.configureHttpApiExtensions(server, gwmux)
			if err := server.Serve(listener); err != nil {
				if !errors.Is(err, http.ErrServerClosed) {
					m.logger.With(logger.Err(err)).Warn("management HTTP server exited with error")
				}
			}
			m.logger.Info("management HTTP server stopped")
		}(server)
	}, httpListenAddr, grpcListenAddr)
	<-ctx.Done()
	mu.Lock()
	if stopServer != nil {
		stopServer()
	}
	mu.Unlock()
	return ctx.Err()
}

func (m *Server) CertsInfo(_ context.Context, _ *emptypb.Empty) (*managementv1.CertsInfoResponse, error) {
	resp := &managementv1.CertsInfoResponse{
		Chain: []*corev1.CertInfo{},
	}
	tlsConfig := m.coreDataSource.TLSConfig()
	if tlsConfig == nil || len(tlsConfig.Certificates) == 0 {
		return resp, nil
	}
	for _, tlsCert := range tlsConfig.Certificates[:1] {
		for _, der := range tlsCert.Certificate {
			cert, err := x509.ParseCertificate(der)
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			resp.Chain = append(resp.Chain, &corev1.CertInfo{
				Issuer:      cert.Issuer.String(),
				Subject:     cert.Subject.String(),
				IsCA:        cert.IsCA,
				NotBefore:   cert.NotBefore.Format(time.RFC3339),
				NotAfter:    cert.NotAfter.Format(time.RFC3339),
				Fingerprint: pkp.NewSha256(cert).Encode(),
				Raw:         cert.Raw,
			})
		}
	}
	return resp, nil
}

func (m *Server) ListCapabilities(ctx context.Context, in *emptypb.Empty) (*managementv1.CapabilityList, error) {
	if m.capabilitiesDataSource == nil {
		return nil, status.Error(codes.Unavailable, "capability backend store not configured")
	}

	clusters, err := m.ListClusters(ctx, &managementv1.ListClustersRequest{})
	if err != nil {
		return nil, err
	}
	counts := make(map[string]int32)
	for _, cluster := range clusters.Items {
		for _, cap := range cluster.GetCapabilities() {
			counts[cap.Name]++
		}
	}

	list, err := m.capabilitiesDataSource.List(ctx, in)
	if err != nil {
		return nil, err
	}
	var items []*managementv1.CapabilityInfo
	for _, details := range list.GetItems() {
		items = append(items, &managementv1.CapabilityInfo{
			Details:   details,
			NodeCount: counts[details.GetName()],
		})
	}

	return &managementv1.CapabilityList{
		Items: items,
	}, nil
}
