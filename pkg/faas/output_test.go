package faas_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sks/mqttfaas/pkg/faas"
)

var _ = Describe("Output", func() {
	var output *faas.Output
	BeforeEach(func() {
		output = faas.NewOutput([]byte(`{"topic":"output_topic","data":"my_fancy_data"}`), nil)
	})
	Describe("AsMessage", func() {
		It("json data ", func() {
			msg, err := output.AsMessage()
			Expect(err).NotTo(HaveOccurred())
			Expect(msg.Topic).To(Equal(`output_topic`))
			Expect(msg.Data).To(Equal("my_fancy_data"))
		})
		Context("negative cases", func() {
			It("returns error when the data is empty", func() {
				output = faas.NewOutput(nil, nil)
				_, err := output.AsMessage()
				Expect(err).To(MatchError(`No data to process`))
			})
			It("returns error when json marshalling fails", func() {
				output = faas.NewOutput([]byte("$%^&"), nil)
				_, err := output.AsMessage()
				Expect(err).To(MatchError(`Could not convert "$%^&" to a mqtt message: invalid character '$' looking for beginning of value`))
			})
		})
	})
})
