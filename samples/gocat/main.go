package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const topicToPostTo = "cat/output"
const outputFormat = `{"topic":%q,"data": %q}`

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input %s\n", err)
		return
	}
	inputFromTopic := os.Getenv("FIRED_BY")
	if inputFromTopic == topicToPostTo {
		return
	}
	// Any message to stderr is ignored
	fmt.Fprintf(os.Stderr, "Topic = %q\n", inputFromTopic)
	inputMessage := strings.TrimSpace(string(data))

	fmt.Printf(outputFormat, topicToPostTo, inputMessage)
}
