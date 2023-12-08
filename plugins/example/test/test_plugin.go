package test

import (
	"github.com/open-panoptes/opni/pkg/plugins/meta"
	"github.com/open-panoptes/opni/pkg/test"
	"github.com/open-panoptes/opni/plugins/example/pkg/example"
)

func init() {
	test.EnablePlugin(meta.ModeGateway, example.Scheme)
}
