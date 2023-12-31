// Code generated by internal/codegen/pathbuilder/generator.go. DO NOT EDIT.
// source: github.com/rancher/opni/pkg/config/v1/gateway_config.proto

package configv1

import (
	v1 "github.com/rancher/opni/pkg/apis/core/v1"
	protopath "google.golang.org/protobuf/reflect/protopath"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type (
	gatewayConfigSpecPathBuilder          protopath.Path
	revisionPathBuilder                   protopath.Path
	timestampPathBuilder                  protopath.Path
	serverSpecPathBuilder                 protopath.Path
	managementServerSpecPathBuilder       protopath.Path
	relayServerSpecPathBuilder            protopath.Path
	healthServerSpecPathBuilder           protopath.Path
	dashboardServerSpecPathBuilder        protopath.Path
	certsSpecPathBuilder                  protopath.Path
	storageSpecPathBuilder                protopath.Path
	etcdSpecPathBuilder                   protopath.Path
	mTLSSpecPathBuilder                   protopath.Path
	jetStreamSpecPathBuilder              protopath.Path
	pluginsSpecPathBuilder                protopath.Path
	pluginFiltersPathBuilder              protopath.Path
	cacheSpecPathBuilder                  protopath.Path
	filesystemCacheSpecPathBuilder        protopath.Path
	keyringSpecPathBuilder                protopath.Path
	upgradesSpecPathBuilder               protopath.Path
	agentUpgradesSpecPathBuilder          protopath.Path
	kubernetesAgentUpgradeSpecPathBuilder protopath.Path
	pluginUpgradesSpecPathBuilder         protopath.Path
	binaryPluginUpgradeSpecPathBuilder    protopath.Path
	rateLimitingSpecPathBuilder           protopath.Path
	authSpecPathBuilder                   protopath.Path
	basicAuthSpecPathBuilder              protopath.Path
	openIDAuthSpecPathBuilder             protopath.Path
)

func (*GatewayConfigSpec) ProtoPath() gatewayConfigSpecPathBuilder {
	return gatewayConfigSpecPathBuilder{protopath.Root(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor())}
}

func (p gatewayConfigSpecPathBuilder) Revision() revisionPathBuilder {
	return revisionPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p gatewayConfigSpecPathBuilder) Server() serverSpecPathBuilder {
	return serverSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p gatewayConfigSpecPathBuilder) Management() managementServerSpecPathBuilder {
	return managementServerSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p gatewayConfigSpecPathBuilder) Relay() relayServerSpecPathBuilder {
	return relayServerSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p gatewayConfigSpecPathBuilder) Health() healthServerSpecPathBuilder {
	return healthServerSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(5))))
}
func (p gatewayConfigSpecPathBuilder) Dashboard() dashboardServerSpecPathBuilder {
	return dashboardServerSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(6))))
}
func (p gatewayConfigSpecPathBuilder) Storage() storageSpecPathBuilder {
	return storageSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(7))))
}
func (p gatewayConfigSpecPathBuilder) Certs() certsSpecPathBuilder {
	return certsSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(8))))
}
func (p gatewayConfigSpecPathBuilder) Plugins() pluginsSpecPathBuilder {
	return pluginsSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(9))))
}
func (p gatewayConfigSpecPathBuilder) Keyring() keyringSpecPathBuilder {
	return keyringSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(10))))
}
func (p gatewayConfigSpecPathBuilder) Upgrades() upgradesSpecPathBuilder {
	return upgradesSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(11))))
}
func (p gatewayConfigSpecPathBuilder) RateLimiting() rateLimitingSpecPathBuilder {
	return rateLimitingSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(12))))
}
func (p gatewayConfigSpecPathBuilder) Auth() authSpecPathBuilder {
	return authSpecPathBuilder(append(p, protopath.FieldAccess(((*GatewayConfigSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(13))))
}
func (p revisionPathBuilder) Timestamp() timestampPathBuilder {
	return timestampPathBuilder(append(p, protopath.FieldAccess(((*v1.Revision)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p dashboardServerSpecPathBuilder) WebCerts() certsSpecPathBuilder {
	return certsSpecPathBuilder(append(p, protopath.FieldAccess(((*DashboardServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(5))))
}
func (p storageSpecPathBuilder) Etcd() etcdSpecPathBuilder {
	return etcdSpecPathBuilder(append(p, protopath.FieldAccess(((*StorageSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p storageSpecPathBuilder) JetStream() jetStreamSpecPathBuilder {
	return jetStreamSpecPathBuilder(append(p, protopath.FieldAccess(((*StorageSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p etcdSpecPathBuilder) Certs() mTLSSpecPathBuilder {
	return mTLSSpecPathBuilder(append(p, protopath.FieldAccess(((*EtcdSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p pluginsSpecPathBuilder) Filters() pluginFiltersPathBuilder {
	return pluginFiltersPathBuilder(append(p, protopath.FieldAccess(((*PluginsSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p pluginsSpecPathBuilder) Cache() cacheSpecPathBuilder {
	return cacheSpecPathBuilder(append(p, protopath.FieldAccess(((*PluginsSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p cacheSpecPathBuilder) Filesystem() filesystemCacheSpecPathBuilder {
	return filesystemCacheSpecPathBuilder(append(p, protopath.FieldAccess(((*CacheSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p upgradesSpecPathBuilder) Agents() agentUpgradesSpecPathBuilder {
	return agentUpgradesSpecPathBuilder(append(p, protopath.FieldAccess(((*UpgradesSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p upgradesSpecPathBuilder) Plugins() pluginUpgradesSpecPathBuilder {
	return pluginUpgradesSpecPathBuilder(append(p, protopath.FieldAccess(((*UpgradesSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p agentUpgradesSpecPathBuilder) Kubernetes() kubernetesAgentUpgradeSpecPathBuilder {
	return kubernetesAgentUpgradeSpecPathBuilder(append(p, protopath.FieldAccess(((*AgentUpgradesSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p pluginUpgradesSpecPathBuilder) Binary() binaryPluginUpgradeSpecPathBuilder {
	return binaryPluginUpgradeSpecPathBuilder(append(p, protopath.FieldAccess(((*PluginUpgradesSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p authSpecPathBuilder) Basic() basicAuthSpecPathBuilder {
	return basicAuthSpecPathBuilder(append(p, protopath.FieldAccess(((*AuthSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p authSpecPathBuilder) Openid() openIDAuthSpecPathBuilder {
	return openIDAuthSpecPathBuilder(append(p, protopath.FieldAccess(((*AuthSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}

func (p revisionPathBuilder) Revision() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*v1.Revision)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p timestampPathBuilder) Seconds() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*timestamppb.Timestamp)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p timestampPathBuilder) Nanos() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*timestamppb.Timestamp)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p serverSpecPathBuilder) HttpListenAddress() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*ServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p serverSpecPathBuilder) GrpcListenAddress() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*ServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p serverSpecPathBuilder) AdvertiseAddress() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*ServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p managementServerSpecPathBuilder) HttpListenAddress() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*ManagementServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p managementServerSpecPathBuilder) GrpcListenAddress() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*ManagementServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p managementServerSpecPathBuilder) AdvertiseAddress() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*ManagementServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p relayServerSpecPathBuilder) GrpcListenAddress() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*RelayServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p relayServerSpecPathBuilder) AdvertiseAddress() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*RelayServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(5))))
}
func (p healthServerSpecPathBuilder) HttpListenAddress() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*HealthServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p dashboardServerSpecPathBuilder) HttpListenAddress() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*DashboardServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p dashboardServerSpecPathBuilder) AdvertiseAddress() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*DashboardServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p dashboardServerSpecPathBuilder) Hostname() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*DashboardServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p dashboardServerSpecPathBuilder) TrustedProxies() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*DashboardServerSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p certsSpecPathBuilder) CaCert() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*CertsSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p certsSpecPathBuilder) CaCertData() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*CertsSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p certsSpecPathBuilder) ServingCert() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*CertsSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p certsSpecPathBuilder) ServingCertData() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*CertsSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p certsSpecPathBuilder) ServingKey() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*CertsSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(5))))
}
func (p certsSpecPathBuilder) ServingKeyData() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*CertsSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(6))))
}
func (p storageSpecPathBuilder) Backend() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*StorageSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p etcdSpecPathBuilder) Endpoints() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*EtcdSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p mTLSSpecPathBuilder) ServerCA() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*MTLSSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p mTLSSpecPathBuilder) ServerCAData() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*MTLSSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p mTLSSpecPathBuilder) ClientCA() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*MTLSSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p mTLSSpecPathBuilder) ClientCAData() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*MTLSSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p mTLSSpecPathBuilder) ClientCert() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*MTLSSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(5))))
}
func (p mTLSSpecPathBuilder) ClientCertData() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*MTLSSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(6))))
}
func (p mTLSSpecPathBuilder) ClientKey() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*MTLSSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(7))))
}
func (p mTLSSpecPathBuilder) ClientKeyData() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*MTLSSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(8))))
}
func (p jetStreamSpecPathBuilder) Endpoint() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*JetStreamSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p jetStreamSpecPathBuilder) NkeySeedPath() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*JetStreamSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p pluginsSpecPathBuilder) Dir() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*PluginsSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p pluginFiltersPathBuilder) Exclude() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*PluginFilters)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p cacheSpecPathBuilder) Backend() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*CacheSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p filesystemCacheSpecPathBuilder) Dir() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*FilesystemCacheSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p keyringSpecPathBuilder) RuntimeKeyDirs() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*KeyringSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p kubernetesAgentUpgradeSpecPathBuilder) ImageResolver() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*KubernetesAgentUpgradeSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p binaryPluginUpgradeSpecPathBuilder) PatchEngine() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*BinaryPluginUpgradeSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p rateLimitingSpecPathBuilder) Rate() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*RateLimitingSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p rateLimitingSpecPathBuilder) Burst() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*RateLimitingSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p authSpecPathBuilder) Backend() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*AuthSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p openIDAuthSpecPathBuilder) Issuer() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*OpenIDAuthSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(1))))
}
func (p openIDAuthSpecPathBuilder) CaCertData() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*OpenIDAuthSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(2))))
}
func (p openIDAuthSpecPathBuilder) ClientId() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*OpenIDAuthSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(3))))
}
func (p openIDAuthSpecPathBuilder) ClientSecret() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*OpenIDAuthSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(4))))
}
func (p openIDAuthSpecPathBuilder) IdentifyingClaim() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*OpenIDAuthSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(5))))
}
func (p openIDAuthSpecPathBuilder) Scopes() protopath.Path {
	return protopath.Path(append(p, protopath.FieldAccess(((*OpenIDAuthSpec)(nil)).ProtoReflect().Descriptor().Fields().ByNumber(6))))
}
