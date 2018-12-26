package retrier_test

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sks/mqttfaas/internal/retrier"
)

var _ = Describe("Retry", func() {
	var (
		sampleFunctionCallCount    int
		sampleFunctionReturnValues []string
		sampleFunction             = func() error {
			sampleFunctionCallCount++
			if sampleFunctionCallCount > len(sampleFunctionReturnValues) {
				return nil
			}
			return errors.New(sampleFunctionReturnValues[sampleFunctionCallCount-1])
		}
	)
	BeforeEach(func() {
		sampleFunctionCallCount = 0
		sampleFunctionReturnValues = []string{}

	})
	It("calls the given function for the given attempts", func() {
		err := retrier.Call(2, time.Nanosecond, sampleFunction)
		Expect(err).NotTo(HaveOccurred())
		Expect(sampleFunctionCallCount).To(Equal(1))
	})
	It("Calls the function multiple times if it errored out before", func() {
		sampleFunctionReturnValues = append(sampleFunctionReturnValues, "something bad happened")
		err := retrier.Call(2, time.Nanosecond, sampleFunction)
		Expect(err).NotTo(HaveOccurred())
		Expect(sampleFunctionCallCount).To(Equal(2))
	})
	It("returns the error if there are too many errors", func() {
		sampleFunctionReturnValues = append(sampleFunctionReturnValues, "1.something bad happened", "2. something bad happened")
		err := retrier.Call(2, time.Nanosecond, sampleFunction)
		Expect(err).To(MatchError(`2. something bad happened`))
		Expect(sampleFunctionCallCount).To(Equal(2))
	})
})
