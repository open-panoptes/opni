package slo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	_ "github.com/open-panoptes/opni/pkg/test/setup"
	_ "github.com/open-panoptes/opni/plugins/alerting/test"
	_ "github.com/open-panoptes/opni/plugins/metrics/test"
)

func TestSlo(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	RegisterFailHandler(Fail)
	RunSpecs(t, "Slo Suite")
}
