package types

import (
	"fmt"
	"regexp"

	"github.com/docker/docker/api/types/container"
)

var cleaningRegex *regexp.Regexp

func init() {
	var err error
	cleaningRegex, err = regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		panic(err)
	}
}

func cleanText(input string) string {
	if []rune(input)[0] == '/' {
		input = input[1:]
	}
	return cleaningRegex.ReplaceAllString(input, "_")
}

//ImageRunnerInput ...
type ImageRunnerInput struct {
	Topic            string
	Message          []byte
	FunctionMetadata FunctionMetadata
	name             string
}

//Name ...
func (i *ImageRunnerInput) Name() string {
	if i.FunctionMetadata.DeleteAfterUse {
		return ""
	}
	if i.name == "" {
		topic := ""
		if !i.FunctionMetadata.NotInterestedInFiredBy {
			topic = cleanText(i.Topic)
		}
		i.name = fmt.Sprintf("mqttfaas_%s-%s", topic, cleanText(i.FunctionMetadata.Image))
		if len(i.name) > 30 {
			i.name = i.name[0:29]
		}
	}
	return i.name
}

//ContainerConfig ...
func (i *ImageRunnerInput) ContainerConfig(defaultEnv []string) *container.Config {
	var env = defaultEnv
	if !i.FunctionMetadata.NotInterestedInFiredBy {
		env = append(env, fmt.Sprintf("FIRED_BY=%s", i.Topic))
	}

	return &container.Config{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Image:        i.FunctionMetadata.Image,
		OpenStdin:    true,
		Tty:          false,
		StdinOnce:    true,
		Labels: map[string]string{
			"mqttfaas_runtime": i.Name(),
			"mqtt_topic":       i.Topic,
		},
		Volumes: map[string]struct{}{
			"/data": *new(struct{}),
		},
		Env: env,
	}
}
