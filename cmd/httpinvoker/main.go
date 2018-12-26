package main

import (
	"log"

	"github.com/sks/mqttfaas/internal/config"
	"github.com/sks/mqttfaas/internal/databus"
	"github.com/sks/mqttfaas/internal/httpinvoker"
	"github.com/sks/mqttfaas/internal/httpregistry"
	"github.com/sks/mqttfaas/internal/types"
)

func main() {
	log.Println("Starting the process")
	cfg := config.New()
	configFileContent, err := types.ParseConfigFile(cfg.ConfFile)
	if err != nil {
		log.Printf("Error reading config file %s:   #%v", cfg.ConfFile, err)
		panic(err)
	}
	httpRegistry := httpregistry.New(configFileContent.HTTPRequests)
	httpInvoker := httpinvoker.New(cfg.HTTPPrefix)
	var dataBus databus.Databus
	dataBus, err = databus.New(configFileContent.Databus, func(topic string, msg []byte) {
		go func(topic string, msg []byte) {
			httpRequests := httpRegistry.GetAllRelavantRequests(topic)
			for i := range httpRequests {
				outputTopic, dataToPublish := httpInvoker.Call(&httpRequests[i], topic, msg)
				if len(topic) != 0 {
					dataBus.Publish(outputTopic, dataToPublish)
				}
			}
		}(topic, msg)
	})
	if err != nil {
		log.Fatalf("Error connecting to MQTT: %s ", err)
	}
	err = dataBus.Connect()
	if err != nil {
		log.Fatalf("Error connecting to MQTT: %s ", err)
	}
	select {}
}
