package topicregistry_test

import (
	"github.com/sks/mqttfaas/internal/topicregistry"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("IsSubTopic", func() {
	DescribeTable("subtopic matcher", func(topic, wildcard string, expectation bool) {
		Expect(topicregistry.IsSubTopic(topic, wildcard)).To(Equal(expectation))
	},
		Entry("same topic", "a/b/c", "a/b/c", true),
		Entry("same topic with slashes", "/", "/", true),
		Entry("same topic with slashes", "//", "//", true),
		Entry("topic with no child selection", "a/b/c", "a/b", false),
		Entry("wild card is #", "a/b/c", "#", true),
		Entry("subtopic with +", "a/b/c", "a/+/c", true),
		Entry("subtopic with #", "a/b/c", "a/#", true),
		Entry("subtopic with + and #", "a/b/c/d", "a/+/#", true),
		Entry("subtopic with 2 +", "a/b/c/d", "a/+/+/d", true),
		Entry("different subtopic", "a/b/c/dog", "a/+/+/e", false),
		Entry("wild card childs", "a/b/c", "a/#", true),
		Entry("+ sign only at one level", "a/b/c", "a/+", false),
		Entry("different topics", "a/b", "12/b", false),
		Entry("different topics", "a/1/b", "a/2/b", false),
		Entry("substring but different", "house/388/room/living/temperature", "output/house/388/room/living/temperature", false),
	)
})
