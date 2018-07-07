package topicregistry

import (
	"context"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	faastype "github.com/sks/mqttfaas/pkg/types"
)

//TopicImageMapper ...
type TopicImageMapper struct {
	dockerCLI DockerCLI
}

//NewTopicImageMapper ...
func NewTopicImageMapper(dockerCLI DockerCLI) *TopicImageMapper {
	return &TopicImageMapper{
		dockerCLI,
	}
}

//GetImages for a given topic figure out the images to use
func (t *TopicImageMapper) GetImages(ctx context.Context, topic string) ([]faastype.FunctionMetadata, error) {
	filters := filters.NewArgs()
	filters.Add("label", "mqtt_faas")
	images, err := t.dockerCLI.ImageList(ctx, types.ImageListOptions{
		Filters: filters,
	})
	if err != nil {
		return nil, err
	}
	output := []faastype.FunctionMetadata{}
	for _, img := range images {
		if len(img.RepoTags) != 0 && t.imageShouldBeRun(img, topic) {
			output = append(output, faastype.NewMetadata(img))
		}
	}
	return output, nil
}

func (t *TopicImageMapper) imageShouldBeRun(img types.ImageSummary, topic string) bool {
	topicsToListenStr, ok := img.Labels["mqtt_faas_topic"]
	if !ok {
		return true
	}
	topicsToListen := strings.Split(topicsToListenStr, "|")
	for _, topicOfInterest := range topicsToListen {
		if isSubTopic(topic, topicOfInterest) {
			return true
		}
	}
	return false
}
