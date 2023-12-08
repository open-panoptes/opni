//go:build !minimal

package apis

import (
	opniaiv1beta1 "github.com/open-panoptes/opni/apis/ai/v1beta1"
	opniloggingv1beta1 "github.com/open-panoptes/opni/apis/logging/v1beta1"
	opnimonitoringv1beta1 "github.com/open-panoptes/opni/apis/monitoring/v1beta1"
)

func init() {
	addSchemeBuilders(
		opniaiv1beta1.AddToScheme,
		opniloggingv1beta1.AddToScheme,
		opnimonitoringv1beta1.AddToScheme,
	)
}
