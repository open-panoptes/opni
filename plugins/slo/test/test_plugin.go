package test

import (
	"time"

	"github.com/open-panoptes/opni/pkg/plugins/meta"
	"github.com/open-panoptes/opni/pkg/test"
	sloapi "github.com/open-panoptes/opni/plugins/slo/apis/slo"
	"github.com/open-panoptes/opni/plugins/slo/pkg/slo"
)

func init() {
	sloapi.MinEvaluateInterval = time.Second
	test.EnablePlugin(meta.ModeGateway, slo.Scheme)
}
