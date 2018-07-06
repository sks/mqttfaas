package processor

import (
	"context"

	"github.com/sks/mqttfaas/pkg/types"
)

//ImageFinder get the list of images to run based on topic
type ImageFinder interface {
	GetImages(context.Context, string) ([]string, error)
}

//ContainerRunner ...
type ContainerRunner interface {
	Run(context.Context, *types.ImageRunnerInput) ([]byte, error)
}
