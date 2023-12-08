package types

import (
	capabilityv1 "github.com/open-panoptes/opni/pkg/apis/capability/v1"
	"github.com/open-panoptes/opni/pkg/metrics/collector"
	"github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions"
	"github.com/open-panoptes/opni/pkg/plugins/apis/system"
)

type (
	HTTPAPIExtensionPlugin       = apiextensions.HTTPAPIExtensionClient
	ManagementAPIExtensionPlugin = apiextensions.ManagementAPIExtensionClient
	StreamAPIExtensionPlugin     = apiextensions.StreamAPIExtensionClient
	UnaryAPIExtensionPlugin      = apiextensions.UnaryAPIExtensionClient
	CapabilityBackendPlugin      = capabilityv1.BackendClient
	CapabilityNodePlugin         = capabilityv1.NodeClient
	MetricsPlugin                = collector.RemoteProducer
	SystemPlugin                 = system.SystemPluginServer
)
