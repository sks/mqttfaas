package outputprocessor_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sks/mqttfaas/internal/outputprocessor"
	"github.com/sks/mqttfaas/internal/outputprocessor/outputprocessorfakes"
	"github.com/sks/mqttfaas/pkg/faas"
)

var _ = Describe("Processor", func() {
	var (
		outputProcessor *outputprocessor.OutputProcessor
		outputs         []*faas.Output
		channel         <-chan *faas.Output
		publisher       *outputprocessorfakes.FakePublisher
		faasOutput      *faas.Output
	)
	BeforeEach(func() {
		outputs = []*faas.Output{}
		channel = make(chan *faas.Output)
		go func() {
			for {
				outputs = append(outputs, <-channel)
			}
		}()

		publisher = new(outputprocessorfakes.FakePublisher)
		outputProcessor = outputprocessor.New(channel, publisher)
		faasOutput = faas.NewOutput([]byte(`{"topic":"my_topic", "data":"my_fancy_output"}`), nil)
	})
	Describe("Process", func() {
		It("process the output", func() {
			err := outputProcessor.Process(faasOutput)
			Expect(err).NotTo(HaveOccurred())
			Expect(publisher.PublishCallCount()).To(Equal(1))
			topic, data := publisher.PublishArgsForCall(0)
			Expect(topic).To(Equal("my_topic"))
			Expect(data).To(Equal("my_fancy_output"))
		})
		Context("negative cases", func() {
			It("Returns error when the message itself has error", func() {
				faasOutput = faas.NewOutput([]byte(``), errors.New(`something bad happened with output`))
				err := outputProcessor.Process(faasOutput)
				Expect(err).To(MatchError(`something bad happened with output`))
			})

			It("returns error when msg cannot be json unmarshalled", func() {
				faasOutput = faas.NewOutput([]byte(`asdasd`), nil)
				err := outputProcessor.Process(faasOutput)
				Expect(err).To(MatchError(`Could not convert "asdasd" to a mqtt message: invalid character 'a' looking for beginning of value`))
			})
			It("does not throw error when the data is empty", func() {
				faasOutput = faas.NewOutput(nil, nil)
				err := outputProcessor.Process(faasOutput)
				Expect(err).NotTo(HaveOccurred())
				Expect(publisher.PublishCallCount()).To(Equal(0))
			})
		})
	})
})
