package types

import (
	"fmt"
	"regexp"
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
	Topic     string
	Message   []byte
	ImageName string
	name      string
}

//Name ...
func (i *ImageRunnerInput) Name() string {
	if i.name == "" {
		i.name = fmt.Sprintf("%s-%s", cleanText(i.Topic), cleanText(i.ImageName))
	}
	return i.name
}
