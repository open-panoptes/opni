package drivers

import (
	"context"

	"github.com/open-panoptes/opni/pkg/plugins/driverutil"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/apis/node"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/apis/rules"
)

type ConfigPropagator interface {
	ConfigureNode(nodeId string, conf *node.AlertingCapabilityConfig) error
}

type NodeDriver interface {
	ConfigPropagator
	DiscoverRules(ctx context.Context) (*rules.RuleManifest, error)
}

var NodeDrivers = driverutil.NewDriverCache[NodeDriver]()
