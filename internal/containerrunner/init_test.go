package containerrunner_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestContainerRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Container Runner")
}
