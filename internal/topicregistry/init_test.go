package topicregistry_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMQTTMatcher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "mqtt matcher")
}
