package alerting

import (
	"github.com/open-panoptes/opni/pkg/agent"
	"github.com/open-panoptes/opni/pkg/capabilities/wellknown"
	streamext "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/stream"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/apis/node"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/apis/rules"
	"google.golang.org/grpc"
)

func (p *Plugin) StreamServers() []streamext.Server {
	return []streamext.Server{
		{
			Desc:              &node.NodeAlertingCapability_ServiceDesc,
			Impl:              &p.node,
			RequireCapability: wellknown.CapabilityAlerting,
		},
		{
			Desc:              &rules.RuleSync_ServiceDesc,
			Impl:              p.AlarmServerComponent,
			RequireCapability: wellknown.CapabilityAlerting,
		},
	}
}

func (p *Plugin) UseStreamClient(cc grpc.ClientConnInterface) {
	p.delegate.Set(streamext.NewDelegate(cc, agent.NewClientSet))
}
