package processor

import (
	"context"
	"log"
	"time"

	"github.com/sks/mqttfaas/pkg/faas"
	"github.com/sks/mqttfaas/pkg/types"
)

//Processor ...
type Processor struct {
	imageFinder     ImageFinder
	containerCLI    ContainerRunner
	outputChan      chan (*faas.Output)
	functionTimeout time.Duration
}

//New ...
func New(imageFinder ImageFinder, containerCLI ContainerRunner,
	outputChan chan *faas.Output,
	functionTimeout time.Duration) *Processor {
	return &Processor{
		imageFinder,
		containerCLI,
		outputChan,
		functionTimeout,
	}
}

//Handle ...
func (p *Processor) Handle(topic string, message []byte) {
	//based on the topic. figure out which all containers to run.
	ctx, cancel := context.WithTimeout(context.Background(), p.functionTimeout)
	defer cancel()

	images, err := p.imageFinder.GetImages(ctx, topic)
	if err != nil {
		log.Printf("Error Getting list of images for the topic %q: %s\n", topic, err)
		return
	}
	for _, image := range images {
		go p.execute(topic, message, image)
	}
}

func (p *Processor) execute(topic string, message []byte, image string) {
	ctx, cancel := context.WithTimeout(context.Background(), p.functionTimeout)
	defer cancel()

	output, err := p.containerCLI.Run(ctx, &types.ImageRunnerInput{
		ImageName: image,
		Message:   message,
		Topic:     topic,
	})
	p.outputChan <- faas.NewOutput(output, err)
}
