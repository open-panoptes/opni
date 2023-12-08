package interfaces

import (
	alertingv1 "github.com/open-panoptes/opni/pkg/apis/alerting/v1"
)

type Routable interface {
	GetRoutingLabels() map[string]string
	GetRoutingAnnotations() map[string]string
	GetRoutingGoldenSignal() string
	// build-time identifier to flag optimizations where possible
	// notifications have the default namespace
	Namespace() string
}

var _ Routable = (*alertingv1.AlertCondition)(nil)
var _ Routable = (*alertingv1.Notification)(nil)
