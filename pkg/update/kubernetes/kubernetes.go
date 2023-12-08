package kubernetes

import (
	"github.com/open-panoptes/opni/pkg/oci"
)

const (
	UpdateStrategy = "kubernetes"
)

var (
	ComponentImageMap = map[ComponentType]oci.ImageType{
		AgentComponent:      oci.ImageTypeMinimal,
		ControllerComponent: oci.ImageTypeOpni,
	}
)

type ComponentType string

const (
	AgentComponent      ComponentType = "agent"
	ControllerComponent ComponentType = "controller"
)
