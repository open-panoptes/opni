package gateway

import (
	"github.com/open-panoptes/opni/pkg/agent"
	"github.com/open-panoptes/opni/pkg/capabilities/wellknown"
	streamext "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/stream"
	"github.com/open-panoptes/opni/plugins/logging/apis/node"
	collogspb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	"google.golang.org/grpc"
)

func (p *Plugin) StreamServers() []streamext.Server {
	return []streamext.Server{
		{
			Desc:              &node.NodeLoggingCapability_ServiceDesc,
			Impl:              &p.logging,
			RequireCapability: wellknown.CapabilityLogs,
		},
		{
			Desc:              &collogspb.LogsService_ServiceDesc,
			Impl:              p.otelForwarder.LogsForwarder,
			RequireCapability: wellknown.CapabilityLogs,
		},
		{
			Desc:              &coltracepb.TraceService_ServiceDesc,
			Impl:              p.otelForwarder.TraceForwarder,
			RequireCapability: wellknown.CapabilityLogs,
		},
	}
}

func (p *Plugin) UseStreamClient(cc grpc.ClientConnInterface) {
	p.delegate.Set(streamext.NewDelegate(cc, agent.NewClientSet))
}
