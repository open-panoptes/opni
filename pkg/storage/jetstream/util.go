package jetstream

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/lestrrat-go/backoff/v2"
	"github.com/nats-io/nats.go"
	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/logger"
)

func AcquireJetstreamConn(ctx context.Context, conf *configv1.JetStreamSpec, lg *slog.Logger) (nats.JetStreamContext, error) {
	nkeyOpt, err := nats.NkeyOptionFromSeed(conf.GetNkeySeedPath())
	if err != nil {
		return nil, err
	}
	nc, err := nats.Connect(conf.GetEndpoint(),
		nkeyOpt,
		nats.MaxReconnects(-1),
		nats.RetryOnFailedConnect(true),
		nats.DisconnectErrHandler(func(c *nats.Conn, err error) {
			if err == nil {
				lg.Debug("jetstream client closed")
				return
			}
			lg.With(
				logger.Err(err),
			).Warn("disconnected from jetstream")
		}),
		nats.ReconnectHandler(func(c *nats.Conn) {
			lg.With(
				"server", c.ConnectedAddr(),
				"id", c.ConnectedServerId(),
				"name", c.ConnectedServerName(),
				"version", c.ConnectedServerVersion(),
			).Info("reconnected to jetstream")
		}),
	)
	if err != nil {
		return nil, err
	}

	ctrl := backoff.Exponential(
		backoff.WithMaxRetries(0),
		backoff.WithMinInterval(10*time.Millisecond),
		backoff.WithMaxInterval(10*time.Millisecond<<9),
		backoff.WithMultiplier(2.0),
	).Start(ctx)
	for {
		if rtt, err := nc.RTT(); err == nil {
			lg.With("rtt", rtt).Info("nats server connection is healthy")
			break
		}
		select {
		case <-ctrl.Done():
			return nil, ctx.Err()
		case <-ctrl.Next():
		}
	}

	js, err := nc.JetStream(nats.Context(ctx))
	if err != nil {
		return nil, err
	}
	return js, nil
}

// Takes a prefix path and replaces invalid elements for jetstream with their valid identifiers
func sanitizePrefix(prefix string) string {
	return strings.ReplaceAll(strings.ReplaceAll(prefix, "/", "-"), ".", "_")
}
