package types

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//ConfigFileContent config file content
type ConfigFileContent struct {
	Databus      Databus       `yaml:"databus"`
	HTTPRequests []HTTPRequest `yaml:"httprequests"`
}

//ParseConfigFile parse the config
func ParseConfigFile(confFile string) (*ConfigFileContent, error) {
	log.Printf("Reading the conf file %s\n", confFile)
	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		return nil, err
	}
	var c ConfigFileContent
	err = yaml.Unmarshal(yamlFile, &c)
	return &c, err
}
