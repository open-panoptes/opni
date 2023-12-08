package main

import (
	"github.com/open-panoptes/opni/pkg/plugins"
	"github.com/open-panoptes/opni/pkg/plugins/meta"
	_ "github.com/open-panoptes/opni/pkg/storage/etcd"
	_ "github.com/open-panoptes/opni/pkg/storage/jetstream"
	"github.com/open-panoptes/opni/plugins/example/pkg/example"
)

func main() {
	m := plugins.Main{
		Modes: meta.ModeSet{
			meta.ModeGateway: example.Scheme,
			meta.ModeAgent:   example.Scheme,
		},
	}
	m.Exec()
}
