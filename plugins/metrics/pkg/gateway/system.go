package gateway

import (
	"context"
	"os"

	"github.com/open-panoptes/opni/plugins/metrics/apis/cortexops"
	"github.com/open-panoptes/opni/plugins/metrics/apis/node"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	corev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	managementv1 "github.com/open-panoptes/opni/pkg/apis/management/v1"
	"github.com/open-panoptes/opni/pkg/config/v1beta1"
	"github.com/open-panoptes/opni/pkg/logger"
	"github.com/open-panoptes/opni/pkg/machinery"
	"github.com/open-panoptes/opni/pkg/plugins/apis/system"
	"github.com/open-panoptes/opni/pkg/task"
	"github.com/open-panoptes/opni/plugins/metrics/pkg/backend"
	"github.com/open-panoptes/opni/plugins/metrics/pkg/cortex"

	_ "github.com/open-panoptes/opni/pkg/storage/etcd"
	_ "github.com/open-panoptes/opni/pkg/storage/jetstream"
	"github.com/open-panoptes/opni/pkg/storage/kvutil"
)

func (p *Plugin) UseManagementAPI(client managementv1.ManagementClient) {
	p.mgmtClient.Set(client)
	cfg, err := client.GetConfig(context.Background(), &emptypb.Empty{}, grpc.WaitForReady(true))
	if err != nil {
		p.logger.With(
			logger.Err(err),
		).Error("failed to get config")
		os.Exit(1)
	}
	objectList, err := machinery.LoadDocuments(cfg.Documents)
	if err != nil {
		p.logger.With(
			logger.Err(err),
		).Error("failed to load config")
		os.Exit(1)
	}
	machinery.LoadAuthProviders(p.ctx, objectList)
	objectList.Visit(func(config *v1beta1.GatewayConfig) {
		backend, err := machinery.ConfigureStorageBackend(p.ctx, &config.Spec.Storage)
		if err != nil {
			p.logger.With(
				logger.Err(err),
			).Error("failed to configure storage backend")
			os.Exit(1)
		}
		p.storageBackend.Set(backend)
		p.config.Set(config)
		tlsConfig := p.loadCortexCerts()
		p.cortexTlsConfig.Set(tlsConfig)
		clientset, err := cortex.NewClientSet(p.ctx, &config.Spec.Cortex, tlsConfig)
		if err != nil {
			p.logger.With(
				logger.Err(err),
			).Error("failed to configure cortex clientset")
			os.Exit(1)
		}
		p.cortexClientSet.Set(clientset)
	})

	p.authMw.Set(machinery.LoadAuthProviders(p.ctx, objectList))
	<-p.ctx.Done()
}

func (p *Plugin) UseKeyValueStore(client system.KeyValueStoreClient) {
	ctrl, err := task.NewController(p.ctx, "uninstall", system.NewKVStoreClient[*corev1.TaskStatus](client), &p.uninstallRunner)
	if err != nil {
		p.logger.With(
			logger.Err(err),
		).Error("failed to create task controller")
		os.Exit(1)
	}
	p.uninstallController.Set(ctrl)

	p.backendKvClients.Set(&backend.KVClients{
		DefaultClusterConfigurationSpec: kvutil.WithKey(system.NewKVStoreClient[*cortexops.CapabilityBackendConfigSpec](client), "/config/cluster/default"),
		DefaultCapabilitySpec:           kvutil.WithKey(system.NewKVStoreClient[*node.MetricsCapabilitySpec](client), "/config/capability/default"),
		NodeCapabilitySpecs:             kvutil.WithPrefix(system.NewKVStoreClient[*node.MetricsCapabilitySpec](client), "/config/capability/nodes/"),
	})
	<-p.ctx.Done()
}
