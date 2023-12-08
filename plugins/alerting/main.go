package main

import (
	"github.com/open-panoptes/opni/pkg/plugins"
	"github.com/open-panoptes/opni/pkg/plugins/meta"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/agent"
	"github.com/open-panoptes/opni/plugins/alerting/pkg/alerting"

	_ "github.com/open-panoptes/opni/plugins/alerting/pkg/agent/drivers/default_driver"
	_ "github.com/open-panoptes/opni/plugins/alerting/pkg/alerting/drivers/alerting_manager"
)

func main() {
	m := plugins.Main{
		Modes: meta.ModeSet{
			meta.ModeGateway: alerting.Scheme,
			meta.ModeAgent:   agent.Scheme,
		},
	}
	m.Exec()
}
