// Package apis can be imported to ensure all plugin APIs are added to client schemes.
package apis

import (
	_ "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions"
	_ "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/http"
	_ "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/management"
	_ "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions/stream"
	_ "github.com/open-panoptes/opni/pkg/plugins/apis/capability"
	_ "github.com/open-panoptes/opni/pkg/plugins/apis/health"
	_ "github.com/open-panoptes/opni/pkg/plugins/apis/metrics"
	_ "github.com/open-panoptes/opni/pkg/plugins/apis/system"
)
