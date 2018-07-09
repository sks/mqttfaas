package config_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sks/mqttfaas/pkg/config"
)

var _ = Describe("Config", func() {
	var configuration *config.Config
	BeforeEach(func() {
		configuration = config.New()
	})
	Describe("new", func() {
		It("returns default values", func() {
			Expect(*configuration).To(Equal(config.Config{
				MQTTConnectionString: "tcp://localhost:1883",
				TopicsToListenTo:     "#",
				DontUseHotContainers: false,
				CleanupTime:          300000000000,
				FunctionTimeout:      5000000000,
				DataDir:              "/data",
			}))
		})
		Context("Envs are set", func() {
			BeforeEach(func() {
				os.Setenv("DELETE_CONTAINER_ONCE_DONE", "true")
				os.Setenv("TOPICS_OF_INTEREST", "my_fancy_topic")
				os.Setenv("MQTT_CONNECTION_STRING", "tcp://mqtt:1883")
				os.Setenv("DATA_DIRECTORY", "/my-data")
			})
			AfterEach(func() {
				os.Unsetenv("DELETE_CONTAINER_ONCE_DONE")
				os.Unsetenv("TOPICS_OF_INTEREST")
				os.Unsetenv("MQTT_CONNECTION_STRING")
			})
			It("honors the env values", func() {
				configuration = config.New()
				Expect(*configuration).To(Equal(config.Config{
					MQTTConnectionString: "tcp://mqtt:1883",
					TopicsToListenTo:     "my_fancy_topic",
					DontUseHotContainers: true,
					CleanupTime:          300000000000,
					FunctionTimeout:      5000000000,
					DataDir:              "/my-data",
				}))
			})
		})

	})
})
