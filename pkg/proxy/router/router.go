package router

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"slices"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	proxyv1 "github.com/rancher/opni/pkg/apis/proxy/v1"
	"github.com/rancher/opni/pkg/logger"
	"github.com/rancher/opni/pkg/proxy"
	"github.com/rancher/opni/pkg/proxy/backend"
	"github.com/rancher/opni/pkg/storage"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	pathParam = "proxyPath"
)

type Router interface {
	SetRoutes(*gin.RouterGroup)
}

type RouterConfig struct {
	Store  storage.RoleBindingStore
	Logger *slog.Logger
	Client proxyv1.RegisterProxyClient
}

func NewRouter(config RouterConfig) (Router, error) {
	backend, err := backend.NewBackend(config.Logger.WithGroup("backend"), config.Client)
	if err != nil {
		return nil, err
	}

	proxyPath, err := config.Client.Path(context.TODO(), &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	path := proxyPath.GetPath()
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DialContext = (&net.Dialer{
		Timeout: 5 * time.Second,
	}).DialContext
	transport.TLSHandshakeTimeout = 5 * time.Second

	return &implRouter{
		backend:   backend,
		store:     config.Store,
		path:      path,
		logger:    config.Logger,
		transport: transport,
		client:    config.Client,
	}, nil
}

type implRouter struct {
	backend   backend.Backend
	store     storage.RoleBindingStore
	path      string
	logger    *slog.Logger
	transport *http.Transport
	client    proxyv1.RegisterProxyClient
}

func (r *implRouter) handle(c *gin.Context) {
	subject, ok := c.Get(proxy.SubjectKey)
	var roleList *corev1.ReferenceList
	var err error
	if ok {
		userID, ok := subject.(string)
		if !ok {
			r.logger.With(logger.Err(err)).Error("failed to get user")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		roleList, err = r.fetchRoles(userID)
		if err != nil {
			r.logger.With(logger.Err(err)).Error("failed to fetch roles")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}

	path, ok := c.Params.Get(pathParam)
	if !ok {
		path = ""
	}

	backendInfo, err := r.client.Backend(c, &emptypb.Empty{})
	if err != nil {
		r.logger.With(logger.Err(err)).Error("failed to get backend info function")
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	target, err := url.Parse(backendInfo.GetBackend())
	if err != nil {
		r.logger.With(logger.Err(err)).Error("failed to get parse backend url")
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	if backendInfo.BackendCAData != nil {
		pool := x509.NewCertPool()
		ok := pool.AppendCertsFromPEM([]byte(backendInfo.GetBackendCAData()))
		if !ok {
			r.logger.Warn("no certs added from pem data")
		}
		r.transport.TLSClientConfig = &tls.Config{
			RootCAs: pool,
		}
	}

	userID, _ := subject.(string)
	rewrite, err := r.backend.RewriteProxyRequest(c, path, target, roleList, userID)
	if err != nil {
		r.logger.With(logger.Err(err)).Error("failed to get rewrite function")
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	proxy := httputil.ReverseProxy{
		Rewrite:   rewrite,
		Transport: r.transport,
		ErrorLog:  slog.NewLogLogger(r.logger.Handler(), slog.LevelDebug), // TODO
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func (r *implRouter) fetchRoles(userID string) (*corev1.ReferenceList, error) {
	if userID == "opni.io_admin" {
		return &corev1.ReferenceList{
			Items: []*corev1.Reference{
				{Id: "opni.io_admin"},
			},
		}, nil
	}
	bindings, err := r.store.ListRoleBindings(context.Background())
	if err != nil {
		return nil, err
	}
	roleList := &corev1.ReferenceList{}
	for _, binding := range bindings.GetItems() {
		if slices.Contains(binding.GetSubjects(), userID) {
			roleList.Items = append(roleList.Items, &corev1.Reference{
				Id: binding.RoleId,
			})
		}
	}
	return roleList, nil
}

func (r *implRouter) SetRoutes(router *gin.RouterGroup) {
	router.Any(fmt.Sprintf("%s/*%s", r.path, pathParam), r.handle)
}
