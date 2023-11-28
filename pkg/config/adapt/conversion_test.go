package adapt_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samber/lo"

	"github.com/rancher/opni/pkg/config/adapt"
	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/config/v1beta1"
	"github.com/rancher/opni/pkg/test/testutil"
	"google.golang.org/protobuf/proto"
)

var _ = Describe("Conversion", func() {
	DescribeTable("V1ConfigOf",
		func(in any, out proto.Message) {
			v1 := adapt.V1ConfigOfUnchecked(in)
			Expect(v1).To(testutil.ProtoEqual(out))
			Expect(adapt.V1BetaConfigOfUnchecked(v1)).To(BeEquivalentTo(in))
		},
		Entry(nil, &v1beta1.GatewayConfigSpec{
			HTTPListenAddress:    "HTTPListenAddress",
			GRPCListenAddress:    "GRPCListenAddress",
			GRPCAdvertiseAddress: "GRPCAdvertiseAddress",
			MetricsListenAddress: "MetricsListenAddress",
			Hostname:             "Hostname",
			Management: v1beta1.ManagementSpec{
				GRPCListenAddress:     "ManagementGRPCListenAddress",
				HTTPListenAddress:     "ManagementHTTPListenAddress",
				WebListenAddress:      "ManagementWebListenAddress",
				RelayListenAddress:    "ManagementRelayListenAddress",
				RelayAdvertiseAddress: "ManagementRelayAdvertiseAddress",
				GRPCAdvertiseAddress:  "ManagementGRPCAdvertiseAddress",
				WebAdvertiseAddress:   "ManagementWebAdvertiseAddress",
			},
			Metrics: v1beta1.MetricsSpec{
				Path: "/metrics",
			},
			TrustedProxies: []string{
				"TrustedProxies",
			},
			Storage: v1beta1.StorageSpec{
				Type: v1beta1.StorageTypeEtcd,
				Etcd: &v1beta1.EtcdStorageSpec{
					Endpoints: []string{
						"Endpoints",
					},
					Certs: &v1beta1.MTLSSpec{
						ServerCA:   "ServerCA",
						ClientCA:   "ClientCA",
						ClientCert: "ClientCert",
						ClientKey:  "ClientKey",
					},
				},
				JetStream: &v1beta1.JetStreamStorageSpec{
					Endpoint:     "Endpoint",
					NkeySeedPath: "NkeySeedPath",
				},
			},
			Certs: v1beta1.CertsSpec{
				CACert:          lo.ToPtr("CACert"),
				CACertData:      []byte("CACertData"),
				ServingCert:     lo.ToPtr("ServingCert"),
				ServingCertData: []byte("ServingCertData"),
				ServingKey:      lo.ToPtr("ServingKey"),
				ServingKeyData:  []byte("ServingKeyData"),
			},
			Plugins: v1beta1.PluginsSpec{
				Dir: "Dir",
				Binary: v1beta1.BinaryPluginsSpec{
					Cache: v1beta1.CacheSpec{
						PatchEngine: v1beta1.PatchEngineBsdiff,
						Backend:     v1beta1.CacheBackendFilesystem,
						Filesystem: v1beta1.FilesystemCacheSpec{
							Dir: "Dir",
						},
					},
				},
			},
			Keyring: v1beta1.KeyringSpec{
				EphemeralKeyDirs: []string{
					"EphemeralKeyDirs",
				},
			},
			AgentUpgrades: v1beta1.AgentUpgradesSpec{
				Kubernetes: v1beta1.KubernetesAgentUpgradeSpec{
					ImageResolver: v1beta1.ImageResolverKubernetes,
				},
			},
			RateLimit: &v1beta1.RateLimitSpec{
				Rate:  1,
				Burst: 1,
			},
		}, &configv1.GatewayConfigSpec{
			Server: &configv1.ServerSpec{
				HttpListenAddress: lo.ToPtr("HTTPListenAddress"),
				GrpcListenAddress: lo.ToPtr("GRPCListenAddress"),
				AdvertiseAddress:  lo.ToPtr("GRPCAdvertiseAddress"),
			},
			Management: &configv1.ManagementServerSpec{
				HttpListenAddress: lo.ToPtr("ManagementHTTPListenAddress"),
				GrpcListenAddress: lo.ToPtr("ManagementGRPCListenAddress"),
				AdvertiseAddress:  lo.ToPtr("ManagementGRPCAdvertiseAddress"),
			},
			Relay: &configv1.RelayServerSpec{
				GrpcListenAddress: lo.ToPtr("ManagementRelayListenAddress"),
				AdvertiseAddress:  lo.ToPtr("ManagementRelayAdvertiseAddress"),
			},
			Health: &configv1.HealthServerSpec{
				HttpListenAddress: lo.ToPtr("MetricsListenAddress"),
			},
			Dashboard: &configv1.DashboardServerSpec{
				HttpListenAddress: lo.ToPtr("ManagementWebListenAddress"),
				Hostname:          lo.ToPtr("Hostname"),
				AdvertiseAddress:  lo.ToPtr("ManagementWebAdvertiseAddress"),
				TrustedProxies:    []string{"TrustedProxies"},
			},
			Storage: &configv1.StorageSpec{
				Backend: configv1.StorageBackend_Etcd.Enum(),
				Etcd: &configv1.EtcdSpec{
					Endpoints: []string{"Endpoints"},
					Certs: &configv1.MTLSSpec{
						ServerCA:   lo.ToPtr("ServerCA"),
						ClientCA:   lo.ToPtr("ClientCA"),
						ClientCert: lo.ToPtr("ClientCert"),
						ClientKey:  lo.ToPtr("ClientKey"),
					},
				},
				JetStream: &configv1.JetStreamSpec{
					Endpoint:     lo.ToPtr("Endpoint"),
					NkeySeedPath: lo.ToPtr("NkeySeedPath"),
				},
			},
			Certs: &configv1.CertsSpec{
				CaCert:          lo.ToPtr("CACert"),
				CaCertData:      lo.ToPtr("CACertData"),
				ServingCert:     lo.ToPtr("ServingCert"),
				ServingCertData: lo.ToPtr("ServingCertData"),
				ServingKey:      lo.ToPtr("ServingKey"),
				ServingKeyData:  lo.ToPtr("ServingKeyData"),
			},
			Plugins: &configv1.PluginsSpec{
				Dir: lo.ToPtr("Dir"),
				Cache: &configv1.CacheSpec{
					Backend: configv1.CacheBackend_Filesystem.Enum(),
					Filesystem: &configv1.FilesystemCacheSpec{
						Dir: lo.ToPtr("Dir"),
					},
				},
			},
			Keyring: &configv1.KeyringSpec{
				RuntimeKeyDirs: []string{"EphemeralKeyDirs"},
			},
			Upgrades: &configv1.UpgradesSpec{
				Agents: &configv1.AgentUpgradesSpec{
					Driver: configv1.AgentUpgradesSpec_Kubernetes.Enum(),
					Kubernetes: &configv1.KubernetesAgentUpgradeSpec{
						ImageResolver: configv1.KubernetesAgentUpgradeSpec_Kubernetes.Enum(),
					},
				},
				Plugins: &configv1.PluginUpgradesSpec{
					Driver: configv1.PluginUpgradesSpec_Binary.Enum(),
					Binary: &configv1.BinaryPluginUpgradeSpec{
						PatchEngine: configv1.PatchEngine_Bsdiff.Enum(),
					},
				},
			},
			RateLimiting: &configv1.RateLimitingSpec{
				Rate:  lo.ToPtr[float64](1),
				Burst: lo.ToPtr[int32](1),
			},
			Auth: &configv1.AuthSpec{},
		}),
	)
	DescribeTable("V1ConfigOf",
		func(in any, out proto.Message) {
			v1 := adapt.V1ConfigOfUnchecked(in)
			Expect(v1).To(Equal(out))
			switch in.(type) {
			case *v1beta1.AgentUpgradesSpec, *v1beta1.BinaryPluginsSpec, *v1beta1.PluginsSpec:
				// cannot be converted back
			default:
				Expect(adapt.V1BetaConfigOfUnchecked(v1)).To(BeEquivalentTo(in))
			}
		},
		Entry(nil, (*v1beta1.GatewayConfigSpec)(nil), (*configv1.GatewayConfigSpec)(nil)),
		Entry(nil, (*v1beta1.StorageSpec)(nil), (*configv1.StorageSpec)(nil)),
		Entry(nil, (*v1beta1.EtcdStorageSpec)(nil), (*configv1.EtcdSpec)(nil)),
		Entry(nil, (*v1beta1.JetStreamStorageSpec)(nil), (*configv1.JetStreamSpec)(nil)),
		Entry(nil, (*v1beta1.MTLSSpec)(nil), (*configv1.MTLSSpec)(nil)),
		Entry(nil, (*v1beta1.CertsSpec)(nil), (*configv1.CertsSpec)(nil)),
		Entry(nil, (*v1beta1.PluginsSpec)(nil), (*configv1.PluginsSpec)(nil)),
		Entry(nil, (*v1beta1.KeyringSpec)(nil), (*configv1.KeyringSpec)(nil)),
		Entry(nil, (*v1beta1.AgentUpgradesSpec)(nil), (*configv1.AgentUpgradesSpec)(nil)),
		Entry(nil, (*v1beta1.BinaryPluginsSpec)(nil), (*configv1.PluginUpgradesSpec)(nil)),
		Entry(nil, (*v1beta1.RateLimitSpec)(nil), (*configv1.RateLimitingSpec)(nil)),
	)
	DescribeTable("V1ConfigOf",
		func(in any, out any) {
			v1 := adapt.V1ConfigOfUnchecked(in)
			Expect(v1).To(Equal(out))
			if in == v1beta1.StorageTypeCRDs {
				Expect(adapt.V1BetaConfigOfUnchecked(v1)).To(Equal(v1beta1.StorageTypeEtcd))
			} else {
				Expect(adapt.V1BetaConfigOfUnchecked(v1)).To(Equal(in))
			}
		},
		Entry(nil, v1beta1.CacheBackendFilesystem, configv1.CacheBackend_Filesystem),
		Entry(nil, v1beta1.PatchEngineBsdiff, configv1.PatchEngine_Bsdiff),
		Entry(nil, v1beta1.PatchEngineZstd, configv1.PatchEngine_Zstd),
		Entry(nil, v1beta1.ImageResolverKubernetes, configv1.KubernetesAgentUpgradeSpec_Kubernetes),
		Entry(nil, v1beta1.ImageResolverNoop, configv1.KubernetesAgentUpgradeSpec_Noop),
		Entry(nil, v1beta1.StorageTypeEtcd, configv1.StorageBackend_Etcd),
		Entry(nil, v1beta1.StorageTypeCRDs, configv1.StorageBackend_Etcd),
		Entry(nil, v1beta1.StorageTypeJetStream, configv1.StorageBackend_JetStream),
	)
})
