package k8sutil_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	_ "github.com/open-panoptes/opni/pkg/test/setup"
)

func TestK8sutil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "K8sutil Suite")
}
