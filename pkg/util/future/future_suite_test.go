package future_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	_ "github.com/open-panoptes/opni/pkg/test/setup"
)

func TestFuture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Future Suite")
}
