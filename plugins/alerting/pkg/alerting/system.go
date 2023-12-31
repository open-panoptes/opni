package alerting

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rancher/opni/pkg/caching"
	"github.com/rancher/opni/pkg/logger"
	"github.com/rancher/opni/pkg/management"
	"github.com/rancher/opni/plugins/alerting/apis/alertops"
	"github.com/rancher/opni/plugins/alerting/pkg/alerting/alarms/v1"
	"github.com/rancher/opni/plugins/alerting/pkg/node_backend"
	"github.com/rancher/opni/plugins/metrics/apis/cortexadmin"
	"github.com/rancher/opni/plugins/metrics/apis/cortexops"

	"github.com/nats-io/nats.go"
	alertingClient "github.com/rancher/opni/pkg/alerting/client"
	"github.com/rancher/opni/pkg/alerting/shared"
	alertingStorage "github.com/rancher/opni/pkg/alerting/storage"

	"github.com/rancher/opni/pkg/alerting/server"
	managementv1 "github.com/rancher/opni/pkg/apis/management/v1"
	"github.com/rancher/opni/pkg/config/adapt"
	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/config/v1beta1"
	"github.com/rancher/opni/pkg/machinery"
	"github.com/rancher/opni/pkg/plugins/apis/system"
	"github.com/rancher/opni/pkg/plugins/driverutil"
	natsutil "github.com/rancher/opni/pkg/util/nats"
	"github.com/rancher/opni/plugins/alerting/pkg/apis/node"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	_ "github.com/rancher/opni/pkg/storage/etcd"
	_ "github.com/rancher/opni/pkg/storage/jetstream"
	"github.com/rancher/opni/pkg/storage/kvutil"
)

func (p *Plugin) UseManagementAPI(client managementv1.ManagementClient) {
	p.mgmtClient.Set(client)

	tlsConfig := p.loadCerts()
	p.configureDriver(
		p.ctx,
		driverutil.NewOption("logger", p.logger.WithGroup("alerting-manager")),
		driverutil.NewOption("subscribers", []chan alertingClient.AlertingClient{p.clusterNotifier}),
		driverutil.NewOption("tlsConfig", tlsConfig),
	)
	p.alertingTLSConfig.Set(tlsConfig)
	go p.handleDriverNotifications()
	go p.runSync()
	p.useWatchers(client)
	<-p.ctx.Done()
}

func (p *Plugin) UseConfigAPI(client configv1.GatewayConfigClient) {
	config, err := client.GetConfiguration(p.ctx, &driverutil.GetRequest{})
	if err != nil {
		p.logger.With(logger.Err(err)).Error("failed to get gateway configuration")
		return
	}
	p.gatewayConfig.Set(&v1beta1.GatewayConfig{
		Spec: *adapt.V1BetaConfigOf[*v1beta1.GatewayConfigSpec](config),
	})
	backend, err := machinery.ConfigureStorageBackendV1(p.ctx, config.Storage)
	if err != nil {
		p.logger.With(
			"err", err,
		).Error("failed to configure storage backend")
		os.Exit(1)
	}
	p.storageBackend.Set(backend)

	<-p.ctx.Done()
}

// UseKeyValueStore Alerting Condition & Alert Endpoints are stored in K,V stores
func (p *Plugin) UseKeyValueStore(client system.KeyValueStoreClient) {
	p.capabilitySpecStore.Set(node_backend.CapabilitySpecKV{
		DefaultCapabilitySpec: kvutil.WithKey(system.NewKVStoreClient[*node.AlertingCapabilitySpec](client), "/alerting/config/capability/default"),
		NodeCapabilitySpecs:   kvutil.WithPrefix(system.NewKVStoreClient[*node.AlertingCapabilitySpec](client), "/alerting/config/capability/nodes"),
	})

	var (
		nc  *nats.Conn
		err error
	)

	cfg := p.gatewayConfig.Get().Spec.Storage.JetStream
	natsURL := os.Getenv("NATS_SERVER_URL")
	natsSeedPath := os.Getenv("NKEY_SEED_FILENAME")
	if cfg == nil {
		cfg = &v1beta1.JetStreamStorageSpec{}
	}
	if cfg.Endpoint == "" {
		cfg.Endpoint = natsURL
	}
	if cfg.NkeySeedPath == "" {
		cfg.NkeySeedPath = natsSeedPath
	}
	nc, err = natsutil.AcquireNATSConnection(
		p.ctx,
		cfg,
		natsutil.WithLogger(p.logger),
		natsutil.WithNatsOptions([]nats.Option{
			nats.ErrorHandler(func(nc *nats.Conn, s *nats.Subscription, err error) {
				if s != nil {
					p.logger.Error("nats : async error in %q/%q: %v", s.Subject, s.Queue, err)
				} else {
					p.logger.Warn("nats : async error outside subscription")
				}
			}),
		}),
	)
	if err != nil {
		p.logger.With(logger.Err(err)).Error("fatal error connecting to NATs")
	}
	p.natsConn.Set(nc)
	mgr, err := p.natsConn.Get().JetStream()
	if err != nil {
		panic(err)
	}
	p.js.Set(mgr)
	b := alertingStorage.NewDefaultAlertingBroker(mgr)
	p.storageClientSet.Set(b.NewClientSet())
	// spawn a reindexing task
	go func() {
		err := p.storageClientSet.Get().ForceSync(p.ctx)
		if err != nil {
			panic(err)
		}
		clStatus, err := p.GetClusterStatus(p.ctx, &emptypb.Empty{})
		if err != nil {
			p.logger.With(logger.Err(err)).Error("failed to get cluster status")
			return
		}
		if clStatus.State == alertops.InstallState_Installed || clStatus.State == alertops.InstallState_InstallUpdating {
			syncInfo, err := p.getSyncInfo(p.ctx)
			if err != nil {
				p.logger.With(logger.Err(err)).Error("failed to get sync info")
			} else {
				for _, comp := range p.Components() {
					comp.Sync(p.ctx, syncInfo)
				}
			}
			conf, err := p.GetClusterConfiguration(p.ctx, &emptypb.Empty{})
			if err != nil {
				p.logger.With(logger.Err(err)).Error("failed to get cluster configuration")
				return
			}
			peers := listPeers(int(conf.GetNumReplicas()))
			p.logger.Info(fmt.Sprintf("reindexing known alerting peers to : %v", peers))
			ctxca, ca := context.WithTimeout(context.Background(), 5*time.Second)
			defer ca()
			alertingClient, err := p.alertingClient.GetContext(ctxca)
			if err != nil {
				p.logger.Error(err.Error())
				return
			}

			alertingClient.MemberlistClient().SetKnownPeers(peers)
			for _, comp := range p.Components() {
				comp.SetConfig(server.Config{
					Client: alertingClient,
				})
			}
		}
	}()
	<-p.ctx.Done()
}

func UseCachingProvider(c caching.CachingProvider[proto.Message]) {
	c.SetCache(caching.NewInMemoryGrpcTtlCache(50*1024*1024, time.Minute))
}

func (p *Plugin) UseAPIExtensions(intf system.ExtensionClientInterface) {
	services := []string{"CortexAdmin", "CortexOps"}
	cc, err := intf.GetClientConn(p.ctx, services...)
	if err != nil {
		if p.ctx.Err() != nil {
			// Plugin is shutting down, don't exit
			return
		}
	}
	p.adminClient.Set(cortexadmin.NewCortexAdminClient(cc))
	p.cortexOpsClient.Set(cortexops.NewCortexOpsClient(cc))
	<-p.ctx.Done()
}

func (p *Plugin) handleDriverNotifications() {
	for {
		select {
		case <-p.ctx.Done():
			p.logger.Info("shutting down cluster driver update handler")
			return
		case client := <-p.clusterNotifier:
			p.logger.Info("updating alerting client based on cluster status")
			serverCfg := server.Config{
				Client: client.Clone(),
			}
			for _, comp := range p.Components() {
				comp.SetConfig(serverCfg)
			}
		}
	}
}

func (p *Plugin) useWatchers(client managementv1.ManagementClient) {
	cw := p.newClusterWatcherHooks(p.ctx, alarms.NewAgentStream())
	clusterCrud, clusterHealthStatus, cortexBackendStatus := func() { p.watchGlobalCluster(client, cw) },
		func() { p.watchGlobalClusterHealthStatus(client, alarms.NewAgentStream()) },
		func() { p.watchCortexClusterStatus() }

	p.globalWatchers = management.NewConditionWatcher(
		clusterCrud,
		clusterHealthStatus,
		cortexBackendStatus,
	)
	p.globalWatchers.WatchEvents()
}

func listPeers(replicas int) []alertingClient.AlertingPeer {
	peers := []alertingClient.AlertingPeer{}
	for i := 0; i < replicas; i++ {
		peers = append(peers, alertingClient.AlertingPeer{
			ApiAddress:      fmt.Sprintf("%s-%d.%s:9093", shared.AlertmanagerService, i, shared.AlertmanagerService),
			EmbeddedAddress: fmt.Sprintf("%s-%d.%s:3000", shared.AlertmanagerService, i, shared.AlertmanagerService),
		})
	}
	return peers
}
