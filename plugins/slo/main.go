package main

import (
	"github.com/open-panoptes/opni/pkg/plugins"
	"github.com/open-panoptes/opni/pkg/plugins/meta"
	"github.com/open-panoptes/opni/plugins/slo/pkg/slo"
)

func main() {
	m := plugins.Main{
		Modes: meta.ModeSet{
			meta.ModeGateway: slo.Scheme,
		},
	}
	m.Exec()
}
