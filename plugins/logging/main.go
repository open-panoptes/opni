package main

import (
	"github.com/open-panoptes/opni/pkg/plugins"
	"github.com/open-panoptes/opni/pkg/plugins/meta"
	"github.com/open-panoptes/opni/plugins/logging/pkg/agent"
	"github.com/open-panoptes/opni/plugins/logging/pkg/gateway"

	_ "github.com/open-panoptes/opni/plugins/logging/pkg/agent/drivers/kubernetes_manager"
	_ "github.com/open-panoptes/opni/plugins/logging/pkg/gateway/drivers/backend/kubernetes_manager"
	_ "github.com/open-panoptes/opni/plugins/logging/pkg/gateway/drivers/management/kubernetes_manager"
)

func main() {
	m := plugins.Main{
		Modes: meta.ModeSet{
			meta.ModeGateway: gateway.Scheme,
			meta.ModeAgent:   agent.Scheme,
		},
	}
	m.Exec()
}
