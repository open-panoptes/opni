package gateway

/*
Declares the topology plugin's gateway meta scheme.
*/

import (
	"context"

	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/open-panoptes/opni/pkg/agent"
	managementv1 "github.com/open-panoptes/opni/pkg/apis/management/v1"
	"github.com/open-panoptes/opni/pkg/config/v1beta1"
	"github.com/open-panoptes/opni/pkg/logger"
	managementext "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/management"
	streamext "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/stream"
	"github.com/open-panoptes/opni/pkg/plugins/apis/capability"
	"github.com/open-panoptes/opni/pkg/plugins/apis/system"
	"github.com/open-panoptes/opni/pkg/plugins/meta"
	"github.com/open-panoptes/opni/pkg/storage"
	"github.com/open-panoptes/opni/pkg/task"
	"github.com/open-panoptes/opni/pkg/util"
	"github.com/open-panoptes/opni/pkg/util/future"
	"github.com/open-panoptes/opni/plugins/topology/apis/orchestrator"
	"github.com/open-panoptes/opni/plugins/topology/apis/representation"
	"github.com/open-panoptes/opni/plugins/topology/pkg/backend"
	"github.com/open-panoptes/opni/plugins/topology/pkg/topology/gateway/drivers"
	"github.com/open-panoptes/opni/plugins/topology/pkg/topology/gateway/stream"
	"google.golang.org/protobuf/proto"
)

type Plugin struct {
	representation.UnsafeTopologyRepresentationServer
	system.UnimplementedSystemPluginClient
	uninstallRunner TopologyUninstallTaskRunner

	ctx    context.Context
	logger *slog.Logger

	topologyRemoteWrite stream.TopologyStreamWriter
	topologyBackend     backend.TopologyBackend

	nc      future.Future[*nats.Conn]
	storage future.Future[ConfigStorageAPIs]

	mgmtClient future.Future[managementv1.ManagementClient]

	storageBackend      future.Future[storage.Backend]
	gatewayConfig       future.Future[*v1beta1.GatewayConfig]
	delegate            future.Future[streamext.StreamDelegate[agent.ClientSet]]
	uninstallController future.Future[*task.Controller]
	clusterDriver       future.Future[drivers.ClusterDriver]
}

func NewPlugin(ctx context.Context) *Plugin {
	p := &Plugin{
		ctx:                 ctx,
		logger:              logger.NewPluginLogger().WithGroup("topology"),
		nc:                  future.New[*nats.Conn](),
		storage:             future.New[ConfigStorageAPIs](),
		mgmtClient:          future.New[managementv1.ManagementClient](),
		storageBackend:      future.New[storage.Backend](),
		delegate:            future.New[streamext.StreamDelegate[agent.ClientSet]](),
		uninstallController: future.New[*task.Controller](),
		clusterDriver:       future.New[drivers.ClusterDriver](),
		topologyBackend:     backend.TopologyBackend{},
		gatewayConfig:       future.New[*v1beta1.GatewayConfig](),
	}
	future.Wait1(p.nc, func(nc *nats.Conn) {
		p.topologyRemoteWrite.Initialize(stream.TopologyStreamWriteConfig{
			Logger: p.logger.With("component", "stream"),
			Nc:     nc,
		})
	})

	future.Wait2(p.storageBackend, p.uninstallController,
		func(storageBackend storage.Backend, uninstallController *task.Controller) {
			p.uninstallRunner.logger = p.logger.WithGroup("topology-uninstall-runner")
			p.uninstallRunner.storageBackend = storageBackend
		})

	p.logger.Debug("waiting for async requirements for starting topology backend")
	future.Wait5(p.storageBackend, p.mgmtClient, p.delegate, p.uninstallController, p.clusterDriver,
		func(
			storageBackend storage.Backend,
			mgmtClient managementv1.ManagementClient,
			delegate streamext.StreamDelegate[agent.ClientSet],
			uninstallController *task.Controller,
			clusterDriver drivers.ClusterDriver,
		) {
			p.logger.With(
				"storageBackend", storageBackend,
				"mgmtClient", mgmtClient,
				"delegate", delegate,
				"uninstallController", uninstallController,
				"clusterDriver", clusterDriver,
			).Debug("async requirements for starting topology backend are ready")
			p.topologyBackend.Initialize(backend.TopologyBackendConfig{
				Logger:              p.logger.WithGroup("topology-backend"),
				StorageBackend:      storageBackend,
				MgmtClient:          mgmtClient,
				Delegate:            delegate,
				UninstallController: uninstallController,
				ClusterDriver:       clusterDriver,
			})
			p.logger.Debug("initialized topology backend")
		})

	return p
}

var _ representation.TopologyRepresentationServer = (*Plugin)(nil)

type ConfigStorageAPIs struct {
	//FIXME: to rename if we find a use
	Placeholder storage.KeyValueStoreT[proto.Message]
}

func Scheme(ctx context.Context) meta.Scheme {
	scheme := meta.NewScheme(meta.WithMode(meta.ModeGateway))
	p := NewPlugin(ctx)
	scheme.Add(system.SystemPluginID, system.NewPlugin(p))
	scheme.Add(managementext.ManagementAPIExtensionPluginID,
		managementext.NewPlugin(
			util.PackService(
				&representation.TopologyRepresentation_ServiceDesc,
				p,
			),
			util.PackService(
				&orchestrator.TopologyOrchestrator_ServiceDesc,
				&p.topologyBackend,
			),
		),
	)
	scheme.Add(streamext.StreamAPIExtensionPluginID, streamext.NewGatewayPlugin(p))
	scheme.Add(capability.CapabilityBackendPluginID, capability.NewPlugin(&p.topologyBackend))
	return scheme
}
