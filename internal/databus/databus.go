package databus

import (
	"fmt"

	"github.com/sks/mqttfaas/internal/databus/mqtt"
	"github.com/sks/mqttfaas/internal/types"
)

//MessageHandler denotes the message handler
type MessageHandler func(topic string, message []byte)

//Databus generic as much as possible
type Databus interface {
	Publish(string, interface{}) error
	Connect() error
}

// New create a new Databus
func New(databus types.Databus, messageHandler MessageHandler) (Databus, error) {
	switch databus.Type {
	case "mqtt":
		mqttServer := databus.GetConfig("connection_string", "tcp://localhost:1883")
		topic := databus.GetConfig("topic", "#")
		return mqtt.New(mqttServer, topic, messageHandler), nil
	default:
		return nil, fmt.Errorf("%s is not a valid databus type. valid values are 'mqtt'", databus.Type)
	}
}
