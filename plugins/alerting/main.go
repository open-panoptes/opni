package main

import (
	"github.com/rancher/opni/pkg/plugins"
	"github.com/rancher/opni/pkg/plugins/meta"
	"github.com/rancher/opni/plugins/alerting/pkg/agent"
	"github.com/rancher/opni/plugins/alerting/pkg/alerting"

	_ "github.com/rancher/opni/plugins/alerting/pkg/agent/drivers/default_driver"
	_ "github.com/rancher/opni/plugins/alerting/pkg/alerting/drivers/alerting_manager"
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
