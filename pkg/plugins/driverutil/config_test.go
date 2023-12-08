package driverutil_test

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/open-panoptes/opni/pkg/storage"
	"github.com/open-panoptes/opni/pkg/storage/inmemory"
	conformance_driverutil "github.com/open-panoptes/opni/pkg/test/conformance/driverutil"
	_ "github.com/open-panoptes/opni/pkg/test/setup"
	"github.com/open-panoptes/opni/pkg/test/testdata/plugins/ext"
	"github.com/open-panoptes/opni/pkg/util"
)

func newValueStore() storage.ValueStoreT[*ext.SampleConfiguration] {
	return inmemory.NewValueStore[*ext.SampleConfiguration](util.ProtoClone)
}

var _ = Describe("Defaulting Config Tracker", Label("unit"), conformance_driverutil.DefaultingConfigTrackerTestSuite(newValueStore, newValueStore))
