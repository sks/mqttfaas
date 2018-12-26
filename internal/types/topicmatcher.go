package types

import "strings"

func isSubTopic(topic, wildcard string) bool {
	return getMatchingTopics(topic, wildcard) != nil
}

func getMatchingTopics(topic, wildcard string) []string {
	output := []string{}
	if topic == wildcard {
		return output
	}
	if wildcard == "#" {
		output = append(output, topic)
		return output
	}
	topicsSlice := strings.Split(topic, "/")
	wildcardSlice := strings.Split(wildcard, "/")
	for i := range topicsSlice {
		if len(wildcardSlice) <= i {
			return nil
		} else if wildcardSlice[i] == "+" {
			output = append(output, topicsSlice[i])
		} else if wildcardSlice[i] == "#" {
			output = append(output, topicsSlice[i])
			return output
		} else if wildcardSlice[i] != topicsSlice[i] {
			return nil
		}
	}

	if len(topicsSlice) == len(wildcardSlice) {
		return output
	}
	return nil

}
