package backend

import (
	"context"
	"slices"
	"sync"

	"log/slog"

	"github.com/open-panoptes/opni/pkg/agent"
	capabilityv1 "github.com/open-panoptes/opni/pkg/apis/capability/v1"
	opnicorev1 "github.com/open-panoptes/opni/pkg/apis/core/v1"
	managementv1 "github.com/open-panoptes/opni/pkg/apis/management/v1"
	"github.com/open-panoptes/opni/pkg/capabilities/wellknown"
	"github.com/open-panoptes/opni/pkg/logger"
	"github.com/open-panoptes/opni/pkg/management"
	streamext "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/stream"
	"github.com/open-panoptes/opni/pkg/storage"
	"github.com/open-panoptes/opni/pkg/task"
	"github.com/open-panoptes/opni/pkg/util"
	"github.com/open-panoptes/opni/plugins/logging/apis/node"
	driver "github.com/open-panoptes/opni/plugins/logging/pkg/gateway/drivers/backend"
	"github.com/open-panoptes/opni/plugins/logging/pkg/opensearchdata"
	loggingutil "github.com/open-panoptes/opni/plugins/logging/pkg/util"
	"google.golang.org/protobuf/types/known/emptypb"
)

type LoggingBackend struct {
	capabilityv1.UnsafeBackendServer
	node.UnsafeNodeLoggingCapabilityServer
	LoggingBackendConfig
	util.Initializer

	nodeStatusMu      sync.RWMutex
	desiredNodeSpecMu sync.RWMutex
	watcher           *management.ManagementWatcherHooks[*managementv1.WatchEvent]
}

type LoggingBackendConfig struct {
	Logger              *slog.Logger                              `validate:"required"`
	StorageBackend      storage.Backend                           `validate:"required"`
	MgmtClient          managementv1.ManagementClient             `validate:"required"`
	Delegate            streamext.StreamDelegate[agent.ClientSet] `validate:"required"`
	UninstallController *task.Controller                          `validate:"required"`
	OpensearchManager   *opensearchdata.Manager                   `validate:"required"`
	ClusterDriver       driver.ClusterDriver                      `validate:"required"`
}

var _ node.NodeLoggingCapabilityServer = (*LoggingBackend)(nil)

// TODO: set up watches on underlying k8s objects to dynamically request a sync
func (b *LoggingBackend) Initialize(conf LoggingBackendConfig) {
	b.InitOnce(func() {
		if err := loggingutil.Validate.Struct(conf); err != nil {
			panic(err)
		}
		b.LoggingBackendConfig = conf

		b.watcher = management.NewManagementWatcherHooks[*managementv1.WatchEvent](context.TODO())
		b.watcher.RegisterHook(func(event *managementv1.WatchEvent) bool {
			return event.Type == managementv1.WatchEventType_Put && slices.ContainsFunc(event.Cluster.Metadata.Capabilities, func(c *opnicorev1.ClusterCapability) bool {
				return c.Name == wellknown.CapabilityLogs
			})
		}, b.updateClusterMetadata)

		go func() {
			clusters, err := b.MgmtClient.ListClusters(context.Background(), &managementv1.ListClustersRequest{})
			if err != nil {
				b.Logger.With(
					logger.Err(err),
				).Error("could not list clusters for reconciliation")
				return
			}

			if err := b.reconcileClusterMetadata(context.Background(), clusters.Items); err != nil {
				b.Logger.With(logger.Err(err)).Error("could not reconcile opni agents with metadata index, some agents may not be included")
				return
			}

			b.watchClusterEvents(context.Background())
		}()
	})
}

func (b *LoggingBackend) Info(context.Context, *emptypb.Empty) (*capabilityv1.Details, error) {
	return &capabilityv1.Details{
		Name:    wellknown.CapabilityLogs,
		Source:  "plugin_logging",
		Drivers: driver.Drivers.List(),
	}, nil
}
