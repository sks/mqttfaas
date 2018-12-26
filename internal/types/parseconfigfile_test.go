package types_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sks/mqttfaas/internal/types"
)

var _ = Describe("Parse Config file", func() {
	Context("ParseConfigFile", func() {
		It("Returns the config file content in a struct", func() {
			content, err := types.ParseConfigFile("testdata/config.yml")
			Expect(err).NotTo(HaveOccurred())
			Expect(*content).To(Equal(types.ConfigFileContent{
				Databus: types.Databus{
					Type: "mqtt",
					Config: map[string]string{
						"key1": "value1",
					},
				},
				HTTPRequests: []types.HTTPRequest{
					{Topics: []string{"/wordcount/input"},
						HTTPPath:    "wc",
						OutputTopic: "/dev/stdout",
						ErrorTopic:  "/dev/stderr",
						Headers: map[string]string{
							"x_custom_header_1": "x_custom_header_1_value",
						},
						Query: map[string]string{
							"x_custom_query_1": "x_custom_query_1_value",
						},
						ContentType: "application/octet-stream",
					},
				},
			}))
		})
	})
})
