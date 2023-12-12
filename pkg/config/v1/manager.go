package configv1

import (
	context "context"

	"github.com/rancher/opni/pkg/config/reactive"
	driverutil "github.com/rancher/opni/pkg/plugins/driverutil"
	storage "github.com/rancher/opni/pkg/storage"
)

type GatewayConfigManagerOptions struct {
	controllerOptions []reactive.ControllerOption
}

type GatewayConfigManagerOption func(*GatewayConfigManagerOptions)

func (o *GatewayConfigManagerOptions) apply(opts ...GatewayConfigManagerOption) {
	for _, op := range opts {
		op(o)
	}
}

func WithControllerOptions(controllerOptions ...reactive.ControllerOption) GatewayConfigManagerOption {
	return func(o *GatewayConfigManagerOptions) {
		o.controllerOptions = append(o.controllerOptions, controllerOptions...)
	}
}

type GatewayConfigManager struct {
	*driverutil.BaseConfigServer[
		*driverutil.GetRequest,
		*SetRequest,
		*ResetRequest,
		*driverutil.ConfigurationHistoryRequest,
		*HistoryResponse,
		*GatewayConfigSpec,
	]

	*reactive.Controller[*GatewayConfigSpec]
}

type GatewayConfigManagerServer struct {
	*GatewayConfigManager
	reactive.ControllerServer[GatewayConfig_WatchReactiveServer, *GatewayConfigSpec]
}

func (m *GatewayConfigManager) AsServer() *GatewayConfigManagerServer {
	s := &GatewayConfigManagerServer{
		GatewayConfigManager: m,
	}
	s.Build(m.Controller)
	return s
}

func NewGatewayConfigManager(
	defaultStore, activeStore storage.ValueStoreT[*GatewayConfigSpec],
	loadDefaultsFunc driverutil.DefaultLoaderFunc[*GatewayConfigSpec],
	opts ...GatewayConfigManagerOption,
) *GatewayConfigManager {
	options := GatewayConfigManagerOptions{}
	options.apply(opts...)

	g := &GatewayConfigManager{}
	g.BaseConfigServer = g.Build(defaultStore, activeStore, loadDefaultsFunc)
	g.Controller = reactive.NewController(g.Tracker(), options.controllerOptions...)
	return g
}

func (s *GatewayConfigManager) DryRun(ctx context.Context, req *DryRunRequest) (*DryRunResponse, error) {
	res, err := s.Tracker().DryRun(ctx, req)
	if err != nil {
		return nil, err
	}
	return &DryRunResponse{
		Current:          res.Current,
		Modified:         res.Modified,
		ValidationErrors: res.ValidationErrors.ToProto(),
	}, nil
}

var ProtoPath = (&GatewayConfigSpec{}).ProtoPath
