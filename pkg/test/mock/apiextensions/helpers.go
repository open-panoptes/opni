package mock_apiextensions

import apiextensions "github.com/open-panoptes/opni/pkg/plugins/apis/apiextensions"

type MockManagementAPIExtensionServerImpl struct {
	apiextensions.UnsafeManagementAPIExtensionServer
	*MockManagementAPIExtensionServer
}
