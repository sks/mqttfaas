package config

import "os"

//Config ...
type Config struct {
	MQTTConnectionString string
	HTTPPrefix           string
	ConfFile             string
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
	return &Config{
		HTTPPrefix: getEnvOrDefault("HTTP_PREFIX", "http://localhost:8080/function"),
		ConfFile:   getEnvOrDefault("CONF_FILE", "/data/config.yml"),
	}
}
