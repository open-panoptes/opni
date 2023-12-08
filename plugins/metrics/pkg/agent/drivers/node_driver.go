package drivers

import (
	"context"
	"fmt"

	"github.com/open-panoptes/opni/pkg/config/v1beta1"
	"github.com/open-panoptes/opni/pkg/plugins/driverutil"
	"github.com/open-panoptes/opni/pkg/rules"
	"github.com/open-panoptes/opni/pkg/util/notifier"
	"github.com/open-panoptes/opni/plugins/metrics/apis/node"
	"github.com/open-panoptes/opni/plugins/metrics/apis/remoteread"
)

type MetricsNodeConfigurator interface {
	ConfigureNode(nodeId string, conf *node.MetricsCapabilityConfig) error
}

type MetricsNodeDriver interface {
	MetricsNodeConfigurator
	DiscoverPrometheuses(context.Context, string) ([]*remoteread.DiscoveryEntry, error)
	ConfigureRuleGroupFinder(config *v1beta1.RulesSpec) notifier.Finder[rules.RuleGroup]
}

var NodeDrivers = driverutil.NewDriverCache[MetricsNodeDriver]()

type ConfigurationError struct {
	NodeId string
	Cfg    *node.MetricsCapabilityConfig
	Err    error
}

func (e *ConfigurationError) Error() string {
	return fmt.Errorf("error configuring node %s: %w", e.NodeId, e.Err).Error()
}
