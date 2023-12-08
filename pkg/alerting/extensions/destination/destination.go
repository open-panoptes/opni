package destination

import (
	"context"

	"github.com/open-panoptes/opni/pkg/alerting/drivers/config"
	alertingv1 "github.com/open-panoptes/opni/pkg/apis/alerting/v1"
)

var (
	defaultSeverity = alertingv1.OpniSeverity_Info.String()
)

const (
	missingTitle = "missing alert title"
	missingBody  = "missing alert body"
)

type Destination interface {
	Push(context.Context, config.WebhookMessage) error
	Name() string
}
