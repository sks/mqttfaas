package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/docker/docker/client"
	"github.com/sks/mqttfaas/internal/containerrunner"
	"github.com/sks/mqttfaas/internal/databus"
	"github.com/sks/mqttfaas/internal/outputprocessor"
	"github.com/sks/mqttfaas/internal/processor"
	"github.com/sks/mqttfaas/internal/topicregistry"
	"github.com/sks/mqttfaas/internal/version"
	"github.com/sks/mqttfaas/pkg/config"
	"github.com/sks/mqttfaas/pkg/faas"
)

func main() {
	log.Printf("Starting MQTT FAAS: %s\n", version.Version())
	errChan := make(chan error)
	go func() {
		var err error
		for {
			err = <-errChan
			if err != nil {
				log.Panic(err)
			}
		}
	}()
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		signal := <-c
		errChan <- errors.New(signal.String())
	}()

	configuration := config.New()

	dockerCLI, err := client.NewEnvClient()
	errChan <- err
	funtionsOutputChan := make(chan *faas.Output)
	defer close(funtionsOutputChan)

	containerRunner := containerrunner.New(dockerCLI, configuration.DataDir, configuration.DontUseHotContainers)
	imageFinder := topicregistry.NewTopicImageMapper(dockerCLI)

	processor := processor.New(imageFinder, containerRunner, funtionsOutputChan, configuration.FunctionTimeout)

	dataBus := databus.New(configuration.MQTTConnectionString, configuration.TopicsToListenTo, processor)

	outputprocessor.New(funtionsOutputChan, dataBus)
	log.Printf("Connecting to MQTT server %s\n", configuration.MQTTConnectionString)

	errChan <- dataBus.Connect()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
