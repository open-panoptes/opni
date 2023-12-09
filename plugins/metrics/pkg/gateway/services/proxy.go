package services

import (
	"context"
	"fmt"

	proxyv1 "github.com/rancher/opni/pkg/apis/proxy/v1"
	"github.com/rancher/opni/pkg/plugins/apis/proxy"
	"github.com/rancher/opni/pkg/plugins/driverutil"
	"github.com/rancher/opni/pkg/plugins/meta"
	"github.com/rancher/opni/plugins/metrics/pkg/types"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WebProxyService struct {
	Context types.ManagementServiceContext `option:"context"`
}

func (s *WebProxyService) Activate() error {
	return nil
}

func (s *WebProxyService) AddToScheme(scheme meta.Scheme) {
	scheme.Add(proxy.ProxyPluginID, proxy.NewPlugin(s))
}

// AuthHeaders implements v1.RegisterProxyServer.
func (s *WebProxyService) AuthHeaders(ctx context.Context, req *proxyv1.HeaderRequest) (*proxyv1.HeaderResponse, error) {
	return &proxyv1.HeaderResponse{
		Headers: []*proxyv1.Header{
			{
				Key:    "X-WEBAUTH-USER",
				Values: []string{req.GetUser()},
			},
			{
				Key:    "X-WEBAUTH-ROLE",
				Values: []string{"Admin"}, // TODO
			},
		},
	}, nil
}

// Backend implements v1.RegisterProxyServer.
func (s *WebProxyService) Backend(ctx context.Context, _ *emptypb.Empty) (*proxyv1.BackendInfo, error) {
	svc := s.Context.ClusterDriver().GetGrafanaServiceConfig()
	return &proxyv1.BackendInfo{
		Backend: fmt.Sprintf("http://%s", svc.HTTPAddress),
	}, nil
}

// Path implements v1.RegisterProxyServer.
func (s *WebProxyService) Path(ctx context.Context, _ *emptypb.Empty) (*proxyv1.PathInfo, error) {
	return &proxyv1.PathInfo{
		Path: "/grafana",
	}, nil
}

func init() {
	types.Services.Register("Web Proxy Service", func(_ context.Context, opts ...driverutil.Option) (types.Service, error) {
		svc := &WebProxyService{}
		driverutil.ApplyOptions(svc, opts...)
		return svc, nil
	})
}
