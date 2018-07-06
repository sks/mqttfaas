package databus

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type messageHandler interface {
	Handle(topic string, message []byte)
}

func (databus Databus) handleMessage(_ mqtt.Client, msg mqtt.Message) {
	databus.messageHandler.Handle(msg.Topic(), msg.Payload())
}

func (databus Databus) newConnectionHandler(client mqtt.Client) {
	log.Printf("Subscribing to Topic %q \n", databus.topic)
	token := client.Subscribe(databus.topic, 2, databus.handleMessage)
	token.Wait()
	err := token.Error()
	if err != nil {
		log.Printf("Error Subsribing to topic %q: %s\n", databus.topic, err)
	}
}

func (databus Databus) newOnConnectionLostHandler(client mqtt.Client, err error) {
	log.Printf("Disconnected from MQTT Server: %s \n", err)
}

//Databus one heck of a databus
type Databus struct {
	client         mqtt.Client
	topic          string
	messageHandler messageHandler
}

//Connect connect to the databus
func (databus *Databus) Connect() error {
	token := databus.client.Connect()
	token.Wait()
	return token.Error()
}

//Publish the message to the topic
func (databus *Databus) Publish(topic string, message interface{}) error {
	token := databus.client.Publish(topic, 1, false, message)
	token.Wait()
	return token.Error()
}

//New create a new databus
func New(mqttServer string, topic string, messageHandler messageHandler) *Databus {
	databus := &Databus{}
	databus.topic = topic
	databus.messageHandler = messageHandler

	clientOptions := mqtt.NewClientOptions()
	clientOptions.AddBroker(mqttServer)
	clientOptions.SetOnConnectHandler(databus.newConnectionHandler)
	clientOptions.SetConnectionLostHandler(databus.newOnConnectionLostHandler)

	databus.client = mqtt.NewClient(clientOptions)

	return databus
}
