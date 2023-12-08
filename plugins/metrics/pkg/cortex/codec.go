package cortex

import "github.com/open-panoptes/opni/pkg/util"

var (
	orgIDCodec = util.NewDelimiterCodec("X-Scope-OrgID", "|")
)
