package dashboard

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"log/slog"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	"github.com/rancher/opni/pkg/auth"
	"github.com/rancher/opni/pkg/auth/local"
	"github.com/rancher/opni/pkg/auth/middleware"
	authutil "github.com/rancher/opni/pkg/auth/util"
	"github.com/rancher/opni/pkg/config/reactive"
	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/logger"
	"github.com/rancher/opni/pkg/plugins"
	"github.com/rancher/opni/pkg/plugins/hooks"
	"github.com/rancher/opni/pkg/plugins/types"
	"github.com/rancher/opni/pkg/proxy"
	proxyrouter "github.com/rancher/opni/pkg/proxy/router"
	"github.com/rancher/opni/pkg/util"
	"github.com/rancher/opni/web"
	ginoauth2 "github.com/zalando/gin-oauth2"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"golang.org/x/oauth2"
	"google.golang.org/protobuf/reflect/protopath"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Server struct {
	ServerOptions
	mgr         *configv1.GatewayConfigManager
	pl          plugins.LoaderInterface
	oauthConfig *oauth2.Config
	ds          AuthDataSource
}

type extraHandler struct {
	method  string
	prefix  string
	handler []gin.HandlerFunc
}

type ServerOptions struct {
	logger        *slog.Logger
	extraHandlers []extraHandler
	assetsFS      fs.FS
	authenticator local.LocalAuthenticator
}

type ServerOption func(*ServerOptions)

func (o *ServerOptions) apply(opts ...ServerOption) {
	for _, op := range opts {
		op(o)
	}
}

func WithHandler(method, prefix string, handler ...gin.HandlerFunc) ServerOption {
	return func(o *ServerOptions) {
		o.extraHandlers = append(o.extraHandlers, extraHandler{
			method:  method,
			prefix:  prefix,
			handler: handler,
		})
	}
}

func WithAssetsFS(fs fs.FS) ServerOption {
	return func(o *ServerOptions) {
		o.assetsFS = fs
	}
}

func WithLocalAuthenticator(authenticator local.LocalAuthenticator) ServerOption {
	return func(o *ServerOptions) {
		o.authenticator = authenticator
	}
}

func WithLogger(logger *slog.Logger) ServerOption {
	return func(o *ServerOptions) {
		o.logger = logger
	}
}

func NewServer(
	ctx context.Context,
	mgr *configv1.GatewayConfigManager,
	pl plugins.LoaderInterface,
	ds AuthDataSource,
	opts ...ServerOption,
) (*Server, error) {
	options := ServerOptions{
		assetsFS:      web.DistFS,
		authenticator: local.NewLocalAuthenticator(ds.StorageBackend().KeyValueStore(authutil.AuthNamespace)),
		logger:        logger.NewNop(),
	}
	options.apply(opts...)

	if !web.WebAssetsAvailable(options.assetsFS) {
		return nil, errors.New("web assets not available")
	}

	return &Server{
		ServerOptions: options,
		mgr:           mgr,
		pl:            pl,
		ds:            ds,
	}, nil
}

func (ws *Server) ListenAndServe(ctx context.Context) error {
	lg := ws.logger

	var mu sync.Mutex
	var cancel context.CancelFunc
	var done chan struct{}

	warn := time.AfterFunc(10*time.Second, func() {
		lg.Warn("dashboard server taking longer than expected to start, check for missing configuration")
	})
	stopOnce := sync.OnceFunc(func() {
		warn.Stop()
	})

	reactive.Bind(ctx, func(v []protoreflect.Value) {
		mu.Lock()
		defer mu.Unlock()
		stopOnce()

		if cancel != nil {
			cancel()
			<-done
		}

		dc := v[0].Message().Interface().(*configv1.DashboardServerSpec)
		mgmtHttpAddr := v[1].String()

		httpListenAddr := dc.GetHttpListenAddress()
		var tlsConfig *tls.Config
		var err error
		if dc.GetWebCerts() == nil {
			err = configv1.ErrInsecure
		} else {
			tlsConfig, err = dc.GetWebCerts().AsTlsConfig(tls.NoClientCert)
		}
		if err != nil {
			if errors.Is(err, configv1.ErrInsecure) {
				lg.With(
					logger.Err(err),
				).Warn("dashboard serving in insecure mode")
			} else {
				lg.With(
					logger.Err(err),
				).Error("failed to configure TLS")
				return
			}
		}
		hostname := dc.GetHostname()
		var listener net.Listener
		if tlsConfig == nil {
			listener, err = net.Listen("tcp4", httpListenAddr)
		} else {
			listener, err = tls.Listen("tcp4", httpListenAddr, tlsConfig)
		}
		if err != nil {
			lg.With(
				"address", httpListenAddr,
				logger.Err(err),
			).Error("failed to start dashboard server")
			return
		}
		lg.With(
			"address", listener.Addr(),
			"hostname", hostname,
		).Info("dashboard server starting")
		webFsTracer := otel.Tracer("webfs")
		router := gin.New()
		router.Use(
			gin.Recovery(),
			logger.GinLogger(ws.logger),
			otelgin.Middleware("opni-ui"),
		)
		middleware := ws.configureAuth(ctx, router)
		// Static assets
		sub, err := fs.Sub(ws.assetsFS, "dist")
		if err != nil {
			lg.With(
				logger.Err(err),
			).Error("assets filesystem missing dist directory")
			return
		}

		webfs := http.FS(sub)
		router.Use(middleware.Handler(ws.checkAdminAccess))

		router.NoRoute(func(c *gin.Context) {
			ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
			defer cancel()
			_, span := webFsTracer.Start(ctx, c.Request.URL.Path)
			defer span.End()
			path := c.Request.URL.Path
			if path[0] == '/' {
				path = path[1:]
			}
			if _, err := fs.Stat(sub, path); err == nil {
				c.FileFromFS(path, webfs)
				return
			}

			c.FileFromFS("/", webfs) // serve index.html
		})

		mgmtUrl, err := url.Parse("http://" + mgmtHttpAddr)
		if err != nil {
			lg.With(
				"url", mgmtHttpAddr,
				"error", err,
			).Error("failed to parse management API URL")
			return
		}
		apiGroup := router.Group("/opni-api")
		apiGroup.Any("/*any", gin.WrapH(http.StripPrefix("/opni-api", httputil.NewSingleHostReverseProxy(mgmtUrl))))

		proxy := router.Group("/proxy")
		proxy.Use(middleware.Handler())
		ws.pl.Hook(hooks.OnLoad(func(p types.ProxyPlugin) {
			log := lg.WithGroup("proxy")
			pluginRouter, err := proxyrouter.NewRouter(proxyrouter.RouterConfig{
				Store:  ws.ds.StorageBackend(),
				Logger: log,
				Client: p,
			})
			if err != nil {
				log.With(
					logger.Err(err),
				).Error("failed to create plugin router")
				return
			}
			pluginRouter.SetRoutes(proxy)
		}))

		for _, h := range ws.extraHandlers {
			router.Handle(h.method, h.prefix, h.handler...)
		}

		var serveContext context.Context
		serveContext, cancel = context.WithCancel(ctx)
		done = make(chan struct{})
		go func() {
			defer close(done)
			if err := util.ServeHandler(serveContext, router, listener); err != nil {
				if !errors.Is(err, context.Canceled) {
					lg.With(logger.Err(err)).Warn("dashboard server exited with error")
					return
				}
			}
			lg.Info("dashboard server stopped")
		}()
	},
		ws.mgr.Reactive(protopath.Path(configv1.ProtoPath().Dashboard())),
		ws.mgr.Reactive(configv1.ProtoPath().Management().HttpListenAddress()),
	)
	<-ctx.Done()
	mu.Lock()
	if cancel != nil {
		cancel()
		<-done
	}
	mu.Unlock()
	return ctx.Err()
}

func (ws *Server) configureAuth(ctx context.Context, router *gin.Engine) *middleware.MultiMiddleware {
	issuer := ""             // TODO: Load this from config
	clientID := ""           // TODO: Load this from config
	clientSecret := ""       //TODO: Load this from config
	scopes := []string{}     //TODO: Load this from config
	localServerAddress := "" // TODO: Load this from config
	var useOIDC bool         //TODO: Load this from config

	if useOIDC {
		provider, err := oidc.NewProvider(ctx, issuer)
		if err != nil {
			panic(err)
		}
		ws.oauthConfig = &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Endpoint:     provider.Endpoint(),
			Scopes:       scopes,
			RedirectURL:  fmt.Sprintf("%s/auth/oidc/callback", localServerAddress),
		}
		handler := oidcHandler{
			logger:   ws.logger.WithGroup("oidc_auth"),
			config:   ws.oauthConfig,
			provider: provider,
		}
		router.Any("/auth/oidc/redirect", handler.handleRedirect)
		router.Any("/auth/oidc/callback", handler.handleCallback)
	}

	middleware := &middleware.MultiMiddleware{
		Logger:             ws.logger.WithGroup("auth_middleware"),
		Config:             ws.oauthConfig,
		IdentifyingClaim:   "user", //TODO: load this from config
		UseOIDC:            useOIDC,
		LocalAuthenticator: ws.authenticator,
	}
	router.GET("/auth/type", middleware.GetAuthType)
	return middleware
}

func (ws *Server) checkAdminAccess(_ *ginoauth2.TokenContainer, ctx *gin.Context) bool {
	lg := ws.logger.WithGroup("auth")
	user, ok := ctx.Get(proxy.SubjectKey)
	if !ok {
		lg.Warn("no user in context")
		return false
	}
	userID, ok := user.(string)
	if !ok {
		lg.Warn("could not find user string in context")
		return false
	}

	if userID == "opni.io_admin" {
		return true
	}

	rb, err := ws.ds.StorageBackend().GetRoleBinding(ctx, &corev1.Reference{
		Id: auth.AdminRoleBindingName,
	})
	if err != nil {
		lg.With(logger.Err(err)).Error("failed to fetch admin user list")
	}

	for _, subject := range rb.Subjects {
		if userID == subject {
			return true
		}
	}
	return false
}
