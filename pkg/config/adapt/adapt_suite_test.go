package adapt_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAdapt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Adapt Suite")
}
