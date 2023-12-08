package routing_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/open-panoptes/opni/pkg/test"
	_ "github.com/open-panoptes/opni/pkg/test/setup"
	"github.com/open-panoptes/opni/pkg/test/testruntime"
	_ "github.com/open-panoptes/opni/plugins/alerting/test"
	_ "github.com/open-panoptes/opni/plugins/metrics/test"
)

func TestRouting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Routing Suite")
}

var (
	env          *test.Environment
	tmpConfigDir string
)

var _ = BeforeSuite(func() {
	testruntime.IfIntegration(func() {
		env = &test.Environment{}
		Expect(env).NotTo(BeNil())
		Expect(env.Start(test.WithEnableNodeExporter(true))).To(Succeed())
		DeferCleanup(env.Stop, "Test Suite Finished")
		tmpConfigDir = env.GenerateNewTempDirectory("alertmanager-config")
		Expect(tmpConfigDir).NotTo(Equal(""))
	})
})
