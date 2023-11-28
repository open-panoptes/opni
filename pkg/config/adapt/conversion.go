package adapt

import (
	"fmt"

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

type v1configtypes interface {
	*configv1.GatewayConfigSpec | // -> *v1beta1.GatewayConfigSpec
		*configv1.StorageSpec | // -> *v1beta1.StorageSpec
		*configv1.EtcdSpec | // -> *v1beta1.EtcdStorageSpec
		*configv1.JetStreamSpec | // -> *v1beta1.JetStreamStorageSpec
		*configv1.MTLSSpec | // -> *v1beta1.MTLSSpec
		*configv1.CertsSpec | // -> *v1beta1.CertsSpec
		*configv1.KeyringSpec | // -> *v1beta1.KeyringSpec
		*configv1.RateLimitingSpec | // -> *v1beta1.RateLimitSpec
		// Enums
		configv1.CacheBackend | // -> *v1beta1.CacheBackend
		configv1.PatchEngine | // -> *v1beta1.PatchEngine
		configv1.KubernetesAgentUpgradeSpec_ImageResolver | // -> *v1beta1.ImageResolverType
		configv1.StorageBackend // -> *v1beta1.StorageType
}

func V1ConfigOf[T v1beta1types](in T) any {
	return v1ConfigOfUnchecked(in)
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
				GrpcListenAddress: &in.GRPCListenAddress,
				AdvertiseAddress:  &in.GRPCAdvertiseAddress,
			},
			Management: &configv1.ManagementServerSpec{
				HttpListenAddress: lo.ToPtr(in.Management.GetHTTPListenAddress()),
				GrpcListenAddress: lo.ToPtr(in.Management.GetGRPCListenAddress()),
				AdvertiseAddress:  &in.Management.GRPCAdvertiseAddress,
			},
			Relay: &configv1.RelayServerSpec{
				GrpcListenAddress: &in.Management.RelayListenAddress,
				AdvertiseAddress:  &in.Management.RelayAdvertiseAddress,
			},
			Health: &configv1.HealthServerSpec{
				HttpListenAddress: &in.MetricsListenAddress,
			},
			Dashboard: &configv1.DashboardServerSpec{
				HttpListenAddress: &in.Management.WebListenAddress,
				Hostname:          &in.Hostname,
				AdvertiseAddress:  &in.Management.WebAdvertiseAddress,
				TrustedProxies:    in.TrustedProxies,
			},
			Storage: V1ConfigOf(&in.Storage).(*configv1.StorageSpec),
			Certs:   V1ConfigOf(&in.Certs).(*configv1.CertsSpec),
			Plugins: V1ConfigOf(&in.Plugins).(*configv1.PluginsSpec),
			Keyring: V1ConfigOf(&in.Keyring).(*configv1.KeyringSpec),
			Upgrades: &configv1.UpgradesSpec{
				Agents: V1ConfigOf(&in.AgentUpgrades).(*configv1.AgentUpgradesSpec),
				Plugins: &configv1.PluginUpgradesSpec{
					Driver: configv1.PluginUpgradesSpec_Binary.Enum(),
					Binary: V1ConfigOf(&in.Plugins.Binary).(*configv1.BinaryPluginUpgradeSpec),
				},
			},
			RateLimiting: V1ConfigOf(in.RateLimit).(*configv1.RateLimitingSpec),
			Auth:         &configv1.AuthSpec{},
		}
	case *v1beta1.StorageSpec:
		if in == nil {
			return (*configv1.StorageSpec)(nil)
		}
		return &configv1.StorageSpec{
			Backend:   V1ConfigOf(in.Type).(configv1.StorageBackend).Enum(),
			Etcd:      V1ConfigOf(in.Etcd).(*configv1.EtcdSpec),
			JetStream: V1ConfigOf(in.JetStream).(*configv1.JetStreamSpec),
		}
	case *v1beta1.EtcdStorageSpec:
		if in == nil {
			return (*configv1.EtcdSpec)(nil)
		}
		return &configv1.EtcdSpec{
			Endpoints: in.Endpoints,
			Certs:     V1ConfigOf(in.Certs).(*configv1.MTLSSpec),
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
				Backend: V1ConfigOf(in.Binary.Cache.Backend).(configv1.CacheBackend).Enum(),
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
			Driver: configv1.AgentUpgradesSpec_Kubernetes.Enum(),
			Kubernetes: &configv1.KubernetesAgentUpgradeSpec{
				ImageResolver: V1ConfigOf(in.Kubernetes.ImageResolver).(configv1.KubernetesAgentUpgradeSpec_ImageResolver).Enum(),
			},
		}
	case *v1beta1.BinaryPluginsSpec:
		if in == nil {
			return (*configv1.PluginUpgradesSpec)(nil)
		}
		return &configv1.BinaryPluginUpgradeSpec{
			PatchEngine: V1ConfigOf(in.Cache.PatchEngine).(configv1.PatchEngine).Enum(),
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

func V1BetaConfigOf[T v1configtypes](in T) any {
	return v1betaConfigOfUnchecked(in)
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
			Storage:              *V1BetaConfigOf(in.Storage).(*v1beta1.StorageSpec),
			Certs:                *V1BetaConfigOf(in.Certs).(*v1beta1.CertsSpec),
			Plugins: v1beta1.PluginsSpec{
				Dir: in.GetPlugins().GetDir(),
				Binary: v1beta1.BinaryPluginsSpec{
					Cache: v1beta1.CacheSpec{
						PatchEngine: V1BetaConfigOf(in.GetUpgrades().GetPlugins().GetBinary().GetPatchEngine()).(v1beta1.PatchEngine),
						Backend:     V1BetaConfigOf(in.GetPlugins().GetCache().GetBackend()).(v1beta1.CacheBackend),
						Filesystem: v1beta1.FilesystemCacheSpec{
							Dir: in.GetPlugins().GetCache().GetFilesystem().GetDir(),
						},
					},
				},
			},
			Keyring: *V1BetaConfigOf(in.Keyring).(*v1beta1.KeyringSpec),
			AgentUpgrades: v1beta1.AgentUpgradesSpec{
				Kubernetes: v1beta1.KubernetesAgentUpgradeSpec{
					ImageResolver: V1BetaConfigOf(in.GetUpgrades().GetAgents().GetKubernetes().GetImageResolver()).(v1beta1.ImageResolverType),
				},
			},
			RateLimit: V1BetaConfigOf(in.RateLimiting).(*v1beta1.RateLimitSpec),
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
			Type:      V1BetaConfigOf(in.GetBackend()).(v1beta1.StorageType),
			Etcd:      V1BetaConfigOf(in.Etcd).(*v1beta1.EtcdStorageSpec),
			JetStream: V1BetaConfigOf(in.JetStream).(*v1beta1.JetStreamStorageSpec),
		}
	case *configv1.EtcdSpec:
		if in == nil {
			return (*v1beta1.EtcdStorageSpec)(nil)
		}
		return &v1beta1.EtcdStorageSpec{
			Endpoints: in.GetEndpoints(),
			Certs:     V1BetaConfigOf(in.Certs).(*v1beta1.MTLSSpec),
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
