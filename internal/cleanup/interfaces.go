package cleanup

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

//DockerCLI ...
//go:generate counterfeiter . DockerCLI
type DockerCLI interface {
	ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error)
}
