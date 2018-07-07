package outputprocessor_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOutputProcessor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "output processor")
}
