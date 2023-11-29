package machinery

import (
	"errors"
	"log/slog"

	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/config/v1beta1"
	"github.com/rancher/opni/pkg/oci"
	"github.com/rancher/opni/pkg/update"
)

func ConfigurePluginUpgrader(cfg v1beta1.PluginUpgradeSpec, pluginDir string, lg *slog.Logger) (update.SyncHandler, error) {
	switch cfg.Type {
	case v1beta1.PluginUpgradeBinary:
		builder := update.GetPluginSyncHandlerBuilder(cfg.Type)
		if builder == nil {
			return nil, errors.New("plugin provider not found")
		}
		return builder(pluginDir, lg)
	default:
		builder := update.GetPluginSyncHandlerBuilder("noop")
		return builder()
	}
}

func ConfigureAgentUpgrader(cfg *v1beta1.AgentUpgradeSpec, lg *slog.Logger) (update.SyncHandler, error) {
	switch {
	case cfg.Type == v1beta1.AgentUpgradeKubernetes:
		builder := update.GetAgentSyncHandlerBuilder(cfg.Type)
		if cfg.Kubernetes != nil {
			return builder(lg, cfg.Kubernetes.Namespace, cfg.Kubernetes.RepoOverride)
		}
		return builder(lg)
	case cfg.Type == v1beta1.AgentUpgradeNoop:
		builder := update.GetAgentSyncHandlerBuilder(cfg.Type)
		return builder()
	default:
		builder := update.GetAgentSyncHandlerBuilder("noop")
		return builder()
	}
}

func ConfigureOCIFetcher(providerType configv1.KubernetesAgentUpgradeSpec_ImageResolver, args ...any) (oci.Fetcher, error) {
	builder := oci.GetFetcherBuilder(providerType.String())
	if builder == nil {
		return nil, errors.New("oci provider not found")
	}
	return builder(args...)
}
