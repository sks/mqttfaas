package types_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sks/mqttfaas/pkg/types"

	dockertypes "github.com/docker/docker/api/types"
)

var _ = Describe("FunctionMetadata", func() {
	var img dockertypes.ImageSummary
	BeforeEach(func() {
		img = dockertypes.ImageSummary{
			ID: "abcdefg",
			RepoTags: []string{
				"user/docker_image:latest",
			},
			Labels: map[string]string{
				"mqtt_faas":                 "my_fancy_faas_name",
				"mqtt_faas_single_use_only": "some_key",
			},
		}
	})
	Describe("New", func() {
		It("parses the image labels to figure out the metadatas for a function", func() {
			metadata := types.NewMetadata(img)
			Expect(metadata).To(Equal(types.FunctionMetadata{
				DeleteAfterUse: true,
				Image:          "user/docker_image:latest",
				Name:           "my_fancy_faas_name",
			}))
		})
		It("marks the function to be deleted soon after use if the label is available", func() {
			delete(img.Labels, "mqtt_faas_single_use_only")
			metadata := types.NewMetadata(img)
			Expect(metadata.DeleteAfterUse).To(BeFalse())
		})
		It("uses the image id if the tags are not available", func() {
			img.RepoTags = []string{}
			metadata := types.NewMetadata(img)
			Expect(metadata.Image).To(Equal("abcdefg"))
		})
	})
})
