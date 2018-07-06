package version

import (
	"os"
)

var appVersion = "DEV"

func init() {
	version := os.Getenv("MQTTFAAS_VERSION")
	if version != "" {
		appVersion = version
	}
}

// Version version of the application
func Version() string {
	return appVersion
}
