package gateway

import (
	managementv1 "github.com/rancher/opni/pkg/apis/management/v1"
	configv1 "github.com/rancher/opni/pkg/config/v1"
	"github.com/rancher/opni/pkg/plugins/apis/system"

	_ "github.com/rancher/opni/pkg/storage/etcd"
	_ "github.com/rancher/opni/pkg/storage/jetstream"
)

// UseManagementAPI implements system.SystemPluginServer.
func (p *Plugin) UseManagementAPI(client managementv1.ManagementClient) {
	p.managementClient.C() <- client
	<-p.ctx.Done()
}

// UseManagementAPI implements system.SystemPluginServer.
func (p *Plugin) UseKeyValueStore(client system.KeyValueStoreClient) {
	p.keyValueStoreClient.C() <- client
	<-p.ctx.Done()
}

// UseManagementAPI implements system.SystemPluginServer.
func (p *Plugin) UseAPIExtensions(intf system.ExtensionClientInterface) {
	p.extensionClient.C() <- intf
	<-p.ctx.Done()
}

func (p *Plugin) UseConfigAPI(client configv1.GatewayConfigClient) {
	p.gatewayConfigClient.C() <- client
	<-p.ctx.Done()
}
