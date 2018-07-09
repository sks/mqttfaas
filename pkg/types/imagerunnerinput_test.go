package types_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sks/mqttfaas/pkg/types"
)

var _ = Describe("ImageRunnerInput", func() {
	var imageRunnerInput *types.ImageRunnerInput
	BeforeEach(func() {
		imageRunnerInput = &types.ImageRunnerInput{
			Topic:   "some_topic",
			Message: []byte{},
			FunctionMetadata: types.FunctionMetadata{
				Image:          "some_function_image:latest",
				DeleteAfterUse: false,
			},
		}
	})
	Describe("Name", func() {
		It("Returns name as a combination of topic and image", func() {
			Expect(imageRunnerInput.Name()).To(Equal("mqttfaas_some_topic-some_func"))
		})
		It("truncates the name to 30 chars", func() {
			imageRunnerInput.Topic = "/some/very/very/very/long/topic/name/that/some/one/crazy/has/given"
			Expect(imageRunnerInput.Name()).To(Equal("mqttfaas_some_very_very_very_"))
		})
		It("does not bother putting topic for functions that dont need topic name", func() {
			imageRunnerInput.FunctionMetadata.NotInterestedInFiredBy = true
			Expect(imageRunnerInput.Name()).To(Equal("mqttfaas_-some_function_image"))
		})
		It("dont bother trying to name a container that is single use", func() {
			imageRunnerInput.FunctionMetadata.DeleteAfterUse = true
			Expect(imageRunnerInput.Name()).To(Equal(""))
		})
	})
})
