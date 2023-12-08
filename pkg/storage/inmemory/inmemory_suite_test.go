package inmemory_test

import (
	"bytes"
	"testing"

	"github.com/open-panoptes/opni/pkg/storage"
	. "github.com/open-panoptes/opni/pkg/test/conformance/storage"
	_ "github.com/open-panoptes/opni/pkg/test/setup"
	"github.com/open-panoptes/opni/pkg/util/future"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/open-panoptes/opni/pkg/storage/inmemory"
)

func TestInmemory(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Inmemory Suite")
}

type testBroker struct{}

func (t testBroker) KeyValueStore(string) storage.KeyValueStore {
	return inmemory.NewKeyValueStore(bytes.Clone)
}

var _ = Describe("In-memory KV Store", Ordered, Label("integration"), KeyValueStoreTestSuite(future.Instant(testBroker{}), NewBytes, Equal))
