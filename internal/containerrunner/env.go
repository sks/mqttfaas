package containerrunner

import (
	"fmt"
	"os"
	"strings"
)

const mqttFaas = "MQTT_FAAS"

var propsToCopy = []string{
	"http_proxy", "https_proxy", "no_proxy",
}

//GetDefaultEnvs ...
func GetDefaultEnvs() []string {
	defaultEnv := []string{}

	for _, e := range os.Environ() {
		if strings.Index(strings.Split(e, "=")[0], mqttFaas) == 0 {
			defaultEnv = append(defaultEnv, e)
		}
	}
	//Copy over proxy infos
	for _, prop := range propsToCopy {
		val := os.Getenv(prop)
		if val != "" {
			defaultEnv = append(defaultEnv, fmt.Sprintf("%s=%s", prop, val))
		}
	}
	return defaultEnv
}
