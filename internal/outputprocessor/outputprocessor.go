package outputprocessor

import (
	"log"

	"github.com/sks/mqttfaas/pkg/faas"
)

//OutputProcessor ...
type OutputProcessor struct {
	publisher Publisher
}

//New ...
func New(channel <-chan *faas.Output, publisher Publisher) *OutputProcessor {
	outputProcessor := &OutputProcessor{
		publisher,
	}
	go func() {
		for {
			err := outputProcessor.Process(<-channel)
			if err != nil {
				log.Printf("Error publishing the message %s\n", err)
			}
		}
	}()

	return outputProcessor
}

//Process ...
func (o *OutputProcessor) Process(output *faas.Output) error {
	if output.Err != nil {
		return output.Err
	}
	msg, err := output.AsMessage()
	if err != nil {
		if err == faas.ErrNoDataToProcess {
			return nil
		}
		return err
	}
	log.Printf("Publishing output to %s\n", msg.Topic)

	return o.publisher.Publish(msg.Topic, msg.Data)
}
