package backend

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"net/http/httputil"

	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	proxyv1 "github.com/rancher/opni/pkg/apis/proxy/v1"
	"github.com/rancher/opni/pkg/logger"
)

type Backend interface {
	RewriteProxyRequest(context.Context, string, *url.URL, *corev1.ReferenceList, string) (func(*httputil.ProxyRequest), error)
}

type implBackend struct {
	pluginClient proxyv1.RegisterProxyClient
	logger       *slog.Logger
}

func NewBackend(logger *slog.Logger, client proxyv1.RegisterProxyClient) (Backend, error) {
	return &implBackend{
		pluginClient: client,
		logger:       logger,
	}, nil
}

func (b *implBackend) RewriteProxyRequest(
	ctx context.Context,
	requestPath string,
	target *url.URL,
	roleList *corev1.ReferenceList,
	user string,
) (func(*httputil.ProxyRequest), error) {
	targetPath, err := url.Parse(requestPath)
	if err != nil {
		b.logger.With(logger.Err(err)).Error(fmt.Sprintf("failed to parse path: %s", requestPath))
		return nil, err
	}

	var extraHeaders *proxyv1.HeaderResponse
	if roleList != nil {
		extraHeaders, err = b.pluginClient.AuthHeaders(context.TODO(), &proxyv1.HeaderRequest{
			User:     user,
			Bindings: roleList,
		})
		if err != nil {
			b.logger.Error("failed to fetch additional headers")
			return nil, err
		}
	}

	return func(r *httputil.ProxyRequest) {
		r.SetURL(target)
		r.Out.URL.Path = targetPath.Path
		r.Out.URL.RawPath = targetPath.RawPath
		r.Out.Host = r.In.Host

		headers := r.In.Header
		for _, header := range extraHeaders.GetHeaders() {
			headers[header.GetKey()] = header.GetValues()
		}
		r.Out.Header = headers
		r.SetXForwarded()

		query := r.In.URL.Query()
		r.Out.URL.RawQuery = query.Encode()
	}, nil
}
