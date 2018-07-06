package config

import (
	"os"
	"time"
)

//Config ...
type Config struct {
	MQTTConnectionString string
	TopicsToListenTo     string
	DontUseHotContainers bool
	FunctionTimeout      time.Duration
}

func getEnvOrDefault(env, defaultValue string) string {
	val := os.Getenv(env)
	if val != "" {
		return val
	}
	return defaultValue
}

//New ...
func New() *Config {
	deleteContainerOnceDone := getEnvOrDefault("DELETE_CONTAINER_ONCE_DONE", "")

	return &Config{
		MQTTConnectionString: getEnvOrDefault("MQTT_CONNECTION_STRING", "tcp://localhost:1883"),
		TopicsToListenTo:     getEnvOrDefault("TOPICS_OF_INTEREST", "#"),
		DontUseHotContainers: deleteContainerOnceDone != "",
		FunctionTimeout:      time.Duration(5 * time.Second),
	}
}
