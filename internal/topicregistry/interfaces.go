package topicregistry

import (
	"context"

	"github.com/docker/docker/api/types"
)

//DockerCLI ...
type DockerCLI interface {
	ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
}
