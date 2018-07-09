package types

import "github.com/docker/docker/api/types"

//FunctionMetadata ...
type FunctionMetadata struct {
	Name                   string
	Image                  string
	DeleteAfterUse         bool
	NotInterestedInFiredBy bool
}

//NewMetadata ...
func NewMetadata(img types.ImageSummary) FunctionMetadata {
	image := img.ID
	if len(img.RepoTags) > 0 {
		image = img.RepoTags[0]
	}
	labels := img.Labels
	return FunctionMetadata{
		Name:                   labels["mqtt_faas"],
		DeleteAfterUse:         len(labels["mqtt_faas_single_use_only"]) != 0,
		NotInterestedInFiredBy: len(labels["mqtt_faas_no_fired_by"]) != 0,
		Image: image,
	}
}
