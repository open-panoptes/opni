package adapt

import (
	"fmt"
	"strings"

	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/config/v1beta1"
	"github.com/samber/lo"
)

type v1beta1types interface {
	*v1beta1.GatewayConfigSpec | // -> *configv1.GatewayConfigSpec
		*v1beta1.StorageSpec | // -> *configv1.StorageSpec
		*v1beta1.EtcdStorageSpec | // -> *configv1.EtcdSpec
		*v1beta1.JetStreamStorageSpec | // -> *configv1.JetStreamSpec
		*v1beta1.MTLSSpec | // -> *configv1.MTLSSpec
		*v1beta1.CertsSpec | // -> *configv1.CertsSpec
		*v1beta1.PluginsSpec | // -> *configv1.PluginsSpec
		*v1beta1.KeyringSpec | // -> *configv1.KeyringSpec
		*v1beta1.AgentUpgradesSpec | // -> *configv1.AgentUpgradesSpec
		*v1beta1.BinaryPluginsSpec | // -> *configv1.PluginUpgradesSpec
		*v1beta1.RateLimitSpec | // -> *configv1.RateLimitingSpec
		// Enums
		v1beta1.CacheBackend | // -> configv1.CacheBackend
		v1beta1.PatchEngine | // -> configv1.PatchEngine
		v1beta1.ImageResolverType | // -> configv1.KubernetesAgentUpgradeSpec_ImageResolver
		v1beta1.StorageType // -> configv1.StorageBackend
}

type v1beta1outputtypes interface {
	*v1beta1.GatewayConfigSpec |
		*v1beta1.StorageSpec |
		*v1beta1.EtcdStorageSpec |
		*v1beta1.JetStreamStorageSpec |
		*v1beta1.MTLSSpec |
		*v1beta1.CertsSpec |
		*v1beta1.KeyringSpec |
		*v1beta1.AgentUpgradesSpec |
		// *v1beta1.BinaryPluginsSpec can't be converted back
		*v1beta1.RateLimitSpec |
		// Enums
		v1beta1.CacheBackend |
		v1beta1.PatchEngine |
		v1beta1.ImageResolverType |
		v1beta1.StorageType
}

type v1configtypes interface {
	*configv1.GatewayConfigSpec | // -> *v1beta1.GatewayConfigSpec
		*configv1.StorageSpec | // -> *v1beta1.StorageSpec
		*configv1.EtcdSpec | // -> *v1beta1.EtcdStorageSpec
		*configv1.JetStreamSpec | // -> *v1beta1.JetStreamStorageSpec
		*configv1.MTLSSpec | // -> *v1beta1.MTLSSpec
		*configv1.CertsSpec | // -> *v1beta1.CertsSpec
		*configv1.KeyringSpec | // -> *v1beta1.KeyringSpec
		*configv1.AgentUpgradesSpec | // -> *v1beta1.AgentUpgradesSpec
		*configv1.RateLimitingSpec | // -> *v1beta1.RateLimitSpec
		// Enums
		configv1.CacheBackend | // -> v1beta1.CacheBackend
		configv1.PatchEngine | // -> v1beta1.PatchEngine
		configv1.KubernetesAgentUpgradeSpec_ImageResolver | // -> v1beta1.ImageResolverType
		configv1.StorageBackend // -> v1beta1.StorageType
}

type v1configoutputtypes interface {
	v1configtypes |
		*configv1.PluginUpgradesSpec |
		*configv1.BinaryPluginUpgradeSpec |
		*configv1.PluginsSpec
}

func V1ConfigOf[O v1configoutputtypes, T v1beta1types](in T) O {
	return v1ConfigOfUnchecked(in).(O)
}

func v1ConfigOfUnchecked(in any) any {
	switch in := in.(type) {
	case *v1beta1.GatewayConfigSpec:
		if in == nil {
			return (*configv1.GatewayConfigSpec)(nil)
		}
		return &configv1.GatewayConfigSpec{
			Server: &configv1.ServerSpec{
				HttpListenAddress: &in.HTTPListenAddress,
				GrpcListenAddress: lo.ToPtr(strings.TrimPrefix(in.GRPCListenAddress, "tcp://")),
				AdvertiseAddress:  &in.GRPCAdvertiseAddress,
			},
			Management: &configv1.ManagementServerSpec{
				HttpListenAddress: lo.ToPtr(in.Management.GetHTTPListenAddress()),
				GrpcListenAddress: lo.ToPtr(strings.TrimPrefix(in.Management.GetGRPCListenAddress(), "tcp://")),
				AdvertiseAddress:  &in.Management.GRPCAdvertiseAddress,
			},
			Relay: &configv1.RelayServerSpec{
				GrpcListenAddress: lo.ToPtr(strings.TrimPrefix(in.Management.RelayListenAddress, "tcp://")),
				AdvertiseAddress:  &in.Management.RelayAdvertiseAddress,
			},
			Health: &configv1.HealthServerSpec{
				HttpListenAddress: &in.MetricsListenAddress,
			},
			Dashboard: &configv1.DashboardServerSpec{
				HttpListenAddress: lo.ToPtr(in.Management.GetWebListenAddress()),
				Hostname:          &in.Hostname,
				AdvertiseAddress:  &in.Management.WebAdvertiseAddress,
				TrustedProxies:    in.TrustedProxies,
			},
			Storage: V1ConfigOf[*configv1.StorageSpec](&in.Storage),
			Certs:   V1ConfigOf[*configv1.CertsSpec](&in.Certs),
			Plugins: V1ConfigOf[*configv1.PluginsSpec](&in.Plugins),
			Keyring: V1ConfigOf[*configv1.KeyringSpec](&in.Keyring),
			Upgrades: &configv1.UpgradesSpec{
				Agents: V1ConfigOf[*configv1.AgentUpgradesSpec](&in.AgentUpgrades),
				Plugins: &configv1.PluginUpgradesSpec{
					Binary: V1ConfigOf[*configv1.BinaryPluginUpgradeSpec](&in.Plugins.Binary),
				},
			},
			RateLimiting: V1ConfigOf[*configv1.RateLimitingSpec](in.RateLimit),
			Auth:         &configv1.AuthSpec{},
		}
	case *v1beta1.StorageSpec:
		if in == nil {
			return (*configv1.StorageSpec)(nil)
		}
		return &configv1.StorageSpec{
			Backend:   V1ConfigOf[configv1.StorageBackend](in.Type).Enum(),
			Etcd:      V1ConfigOf[*configv1.EtcdSpec](in.Etcd),
			JetStream: V1ConfigOf[*configv1.JetStreamSpec](in.JetStream),
		}
	case *v1beta1.EtcdStorageSpec:
		if in == nil {
			return (*configv1.EtcdSpec)(nil)
		}
		return &configv1.EtcdSpec{
			Endpoints: in.Endpoints,
			Certs:     V1ConfigOf[*configv1.MTLSSpec](in.Certs),
		}
	case *v1beta1.JetStreamStorageSpec:
		if in == nil {
			return (*configv1.JetStreamSpec)(nil)
		}
		return &configv1.JetStreamSpec{
			Endpoint:     &in.Endpoint,
			NkeySeedPath: &in.NkeySeedPath,
		}
	case *v1beta1.MTLSSpec:
		if in == nil {
			return (*configv1.MTLSSpec)(nil)
		}
		return &configv1.MTLSSpec{
			ServerCA:   &in.ServerCA,
			ClientCA:   &in.ClientCA,
			ClientCert: &in.ClientCert,
			ClientKey:  &in.ClientKey,
		}
	case *v1beta1.CertsSpec:
		if in == nil {
			return (*configv1.CertsSpec)(nil)
		}
		return &configv1.CertsSpec{
			CaCert:          in.CACert,
			CaCertData:      lo.Ternary(len(in.CACertData) > 0, lo.ToPtr(string(in.CACertData)), nil),
			ServingCert:     in.ServingCert,
			ServingCertData: lo.Ternary(len(in.ServingCertData) > 0, lo.ToPtr(string(in.ServingCertData)), nil),
			ServingKey:      in.ServingKey,
			ServingKeyData:  lo.Ternary(len(in.ServingKeyData) > 0, lo.ToPtr(string(in.ServingKeyData)), nil),
		}
	case *v1beta1.PluginsSpec:
		if in == nil {
			return (*configv1.PluginsSpec)(nil)
		}
		return &configv1.PluginsSpec{
			Dir: &in.Dir,
			Cache: &configv1.CacheSpec{
				Backend: V1ConfigOf[configv1.CacheBackend](in.Binary.Cache.Backend).Enum(),
				Filesystem: &configv1.FilesystemCacheSpec{
					Dir: &in.Binary.Cache.Filesystem.Dir,
				},
			},
		}
	case *v1beta1.KeyringSpec:
		if in == nil {
			return (*configv1.KeyringSpec)(nil)
		}
		return &configv1.KeyringSpec{
			RuntimeKeyDirs: in.EphemeralKeyDirs,
		}
	case *v1beta1.AgentUpgradesSpec:
		if in == nil {
			return (*configv1.AgentUpgradesSpec)(nil)
		}
		return &configv1.AgentUpgradesSpec{
			Kubernetes: &configv1.KubernetesAgentUpgradeSpec{
				ImageResolver: V1ConfigOf[configv1.KubernetesAgentUpgradeSpec_ImageResolver](in.Kubernetes.ImageResolver).Enum(),
			},
		}
	case *v1beta1.BinaryPluginsSpec:
		if in == nil {
			return (*configv1.PluginUpgradesSpec)(nil)
		}
		return &configv1.BinaryPluginUpgradeSpec{
			PatchEngine: V1ConfigOf[configv1.PatchEngine](in.Cache.PatchEngine).Enum(),
		}
	case *v1beta1.RateLimitSpec:
		if in == nil {
			return (*configv1.RateLimitingSpec)(nil)
		}
		return &configv1.RateLimitingSpec{
			Rate: &in.Rate,
			Burst: func() *int32 {
				burst := int32(in.Burst)
				return &burst
			}(),
		}
	case v1beta1.CacheBackend:
		return configv1.CacheBackend_Filesystem
	case v1beta1.PatchEngine:
		switch in {
		case v1beta1.PatchEngineBsdiff:
			return configv1.PatchEngine_Bsdiff
		case v1beta1.PatchEngineZstd:
			fallthrough
		default:
			return configv1.PatchEngine_Zstd
		}
	case v1beta1.ImageResolverType:
		switch in {
		case v1beta1.ImageResolverNoop:
			fallthrough
		default:
			return configv1.KubernetesAgentUpgradeSpec_Noop
		case v1beta1.ImageResolverKubernetes:
			return configv1.KubernetesAgentUpgradeSpec_Kubernetes
		}
	case v1beta1.StorageType:
		switch in {
		case v1beta1.StorageTypeEtcd:
			fallthrough
		default:
			return configv1.StorageBackend_Etcd
		case v1beta1.StorageTypeJetStream:
			return configv1.StorageBackend_JetStream
		}
	default:
		panic(fmt.Sprintf("unsupported conversion from type: %T", in))
	}
}

func V1BetaConfigOf[O v1beta1outputtypes, T v1configtypes](in T) O {
	return v1betaConfigOfUnchecked(in).(O)
}

func v1betaConfigOfUnchecked(in any) any {
	switch in := in.(type) {
	case *configv1.GatewayConfigSpec:
		if in == nil {
			return (*v1beta1.GatewayConfigSpec)(nil)
		}
		return &v1beta1.GatewayConfigSpec{
			HTTPListenAddress:    in.GetServer().GetHttpListenAddress(),
			GRPCListenAddress:    in.GetServer().GetGrpcListenAddress(),
			GRPCAdvertiseAddress: in.GetServer().GetAdvertiseAddress(),
			Management: v1beta1.ManagementSpec{
				GRPCAdvertiseAddress:  in.GetManagement().GetAdvertiseAddress(),
				HTTPListenAddress:     in.GetManagement().GetHttpListenAddress(),
				GRPCListenAddress:     in.GetManagement().GetGrpcListenAddress(),
				RelayListenAddress:    in.GetRelay().GetGrpcListenAddress(),
				RelayAdvertiseAddress: in.GetRelay().GetAdvertiseAddress(),
				WebListenAddress:      in.GetDashboard().GetHttpListenAddress(),
				WebAdvertiseAddress:   in.GetDashboard().GetAdvertiseAddress(),
			},
			MetricsListenAddress: in.GetHealth().GetHttpListenAddress(),
			Hostname:             in.GetDashboard().GetHostname(),
			Storage:              *V1BetaConfigOf[*v1beta1.StorageSpec](in.Storage),
			Certs:                *V1BetaConfigOf[*v1beta1.CertsSpec](in.Certs),
			Plugins: v1beta1.PluginsSpec{
				Dir: in.GetPlugins().GetDir(),
				Binary: v1beta1.BinaryPluginsSpec{
					Cache: v1beta1.CacheSpec{
						PatchEngine: V1BetaConfigOf[v1beta1.PatchEngine](in.GetUpgrades().GetPlugins().GetBinary().GetPatchEngine()),
						Backend:     V1BetaConfigOf[v1beta1.CacheBackend](in.GetPlugins().GetCache().GetBackend()),
						Filesystem: v1beta1.FilesystemCacheSpec{
							Dir: in.GetPlugins().GetCache().GetFilesystem().GetDir(),
						},
					},
				},
			},
			Keyring: *V1BetaConfigOf[*v1beta1.KeyringSpec](in.Keyring),
			AgentUpgrades: v1beta1.AgentUpgradesSpec{
				Kubernetes: v1beta1.KubernetesAgentUpgradeSpec{
					ImageResolver: V1BetaConfigOf[v1beta1.ImageResolverType](in.GetUpgrades().GetAgents().GetKubernetes().GetImageResolver()),
				},
			},
			RateLimit: V1BetaConfigOf[*v1beta1.RateLimitSpec](in.RateLimiting),
			Metrics: v1beta1.MetricsSpec{
				Path: "/metrics",
			},
			TrustedProxies: in.GetDashboard().GetTrustedProxies(),
		}
	case *configv1.StorageSpec:
		if in == nil {
			return (*v1beta1.StorageSpec)(nil)
		}
		return &v1beta1.StorageSpec{
			Type:      V1BetaConfigOf[v1beta1.StorageType](in.GetBackend()),
			Etcd:      V1BetaConfigOf[*v1beta1.EtcdStorageSpec](in.Etcd),
			JetStream: V1BetaConfigOf[*v1beta1.JetStreamStorageSpec](in.JetStream),
		}
	case *configv1.EtcdSpec:
		if in == nil {
			return (*v1beta1.EtcdStorageSpec)(nil)
		}
		return &v1beta1.EtcdStorageSpec{
			Endpoints: in.GetEndpoints(),
			Certs:     V1BetaConfigOf[*v1beta1.MTLSSpec](in.Certs),
		}
	case *configv1.JetStreamSpec:
		if in == nil {
			return (*v1beta1.JetStreamStorageSpec)(nil)
		}
		return &v1beta1.JetStreamStorageSpec{
			Endpoint:     in.GetEndpoint(),
			NkeySeedPath: in.GetNkeySeedPath(),
		}
	case *configv1.MTLSSpec:
		if in == nil {
			return (*v1beta1.MTLSSpec)(nil)
		}
		return &v1beta1.MTLSSpec{
			ServerCA:   in.GetServerCA(),
			ClientCA:   in.GetClientCA(),
			ClientCert: in.GetClientCert(),
			ClientKey:  in.GetClientKey(),
		}
	case *configv1.CertsSpec:
		if in == nil {
			return (*v1beta1.CertsSpec)(nil)
		}
		return &v1beta1.CertsSpec{
			CACert:          in.CaCert,
			CACertData:      []byte(in.GetCaCertData()),
			ServingCert:     in.ServingCert,
			ServingCertData: []byte(in.GetServingCertData()),
			ServingKey:      in.ServingKey,
			ServingKeyData:  []byte(in.GetServingKeyData()),
		}
	case *configv1.KeyringSpec:
		if in == nil {
			return (*v1beta1.KeyringSpec)(nil)
		}
		return &v1beta1.KeyringSpec{
			EphemeralKeyDirs: in.GetRuntimeKeyDirs(),
		}
	case *configv1.RateLimitingSpec:
		if in == nil {
			return (*v1beta1.RateLimitSpec)(nil)
		}
		return &v1beta1.RateLimitSpec{
			Rate:  in.GetRate(),
			Burst: int(in.GetBurst()),
		}
	case configv1.CacheBackend:
		return v1beta1.CacheBackendFilesystem
	case configv1.PatchEngine:
		switch in.Number() {
		case configv1.PatchEngine_Bsdiff.Number():
			return v1beta1.PatchEngineBsdiff
		case configv1.PatchEngine_Zstd.Number():
			fallthrough
		default:
			return v1beta1.PatchEngineZstd
		}
	case configv1.KubernetesAgentUpgradeSpec_ImageResolver:
		switch in.Number() {
		case configv1.KubernetesAgentUpgradeSpec_Noop.Number():
			return v1beta1.ImageResolverNoop
		case configv1.KubernetesAgentUpgradeSpec_Kubernetes.Number():
			fallthrough
		default:
			return v1beta1.ImageResolverKubernetes
		}
	case configv1.StorageBackend:
		switch in.Number() {
		case configv1.StorageBackend_Etcd.Number():
			return v1beta1.StorageTypeEtcd
		case configv1.StorageBackend_JetStream.Number():
			fallthrough
		default:
			return v1beta1.StorageTypeJetStream
		}
	default:
		panic(fmt.Sprintf("unsupported conversion from type: %T", in))
	}
}
