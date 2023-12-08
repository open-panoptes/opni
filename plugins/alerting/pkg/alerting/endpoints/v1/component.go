package endpoints

import (
	"context"
	"sync"

	"log/slog"

	"github.com/open-panoptes/opni/pkg/alerting/server"
	alertingSync "github.com/open-panoptes/opni/pkg/alerting/server/sync"
	"github.com/open-panoptes/opni/pkg/alerting/storage/spec"
	alertingv1 "github.com/open-panoptes/opni/pkg/apis/alerting/v1"
	"github.com/open-panoptes/opni/pkg/util"
	"github.com/open-panoptes/opni/pkg/util/future"
	notifications "github.com/open-panoptes/opni/plugins/alerting/pkg/alerting/notifications/v1"
)

type manualSync func(ctx context.Context, hashRing spec.HashRing, routers spec.RouterStorage) error

type EndpointServerComponent struct {
	alertingv1.UnsafeAlertEndpointsServer

	ctx context.Context
	util.Initializer

	mu sync.Mutex
	server.Config

	notifications *notifications.NotificationServerComponent

	logger *slog.Logger

	endpointStorage  future.Future[spec.EndpointStorage]
	conditionStorage future.Future[spec.ConditionStorage]
	routerStorage    future.Future[spec.RouterStorage]
	hashRing         future.Future[spec.HashRing]
}

var _ server.ServerComponent = (*EndpointServerComponent)(nil)

func NewEndpointServerComponent(
	ctx context.Context,
	logger *slog.Logger,
	notifications *notifications.NotificationServerComponent,
) *EndpointServerComponent {
	return &EndpointServerComponent{
		ctx:              ctx,
		logger:           logger,
		notifications:    notifications,
		endpointStorage:  future.New[spec.EndpointStorage](),
		conditionStorage: future.New[spec.ConditionStorage](),
		routerStorage:    future.New[spec.RouterStorage](),
		hashRing:         future.New[spec.HashRing](),
	}
}

type EndpointServerConfiguration struct {
	spec.EndpointStorage
	spec.ConditionStorage
	spec.RouterStorage
	spec.HashRing
}

func (e *EndpointServerComponent) Name() string {
	return "endpoint"
}

func (e *EndpointServerComponent) Status() server.Status {
	return server.Status{
		Running: e.Initialized(),
	}
}

func (e *EndpointServerComponent) Ready() bool {
	return e.Initialized()
}

func (e *EndpointServerComponent) Healthy() bool {
	return e.Initialized()
}

func (e *EndpointServerComponent) SetConfig(conf server.Config) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.Config = conf
}

func (e *EndpointServerComponent) Sync(_ context.Context, _ alertingSync.SyncInfo) error {
	return nil
}

func (e *EndpointServerComponent) Initialize(conf EndpointServerConfiguration) {
	e.InitOnce(func() {
		e.endpointStorage.Set(conf.EndpointStorage)
		e.conditionStorage.Set(conf.ConditionStorage)
		e.routerStorage.Set(conf.RouterStorage)
		e.hashRing.Set(conf.HashRing)
	})
}
