package containerrunner

import (
	"context"
	"fmt"
	"path/filepath"

	dockertypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/sks/mqttfaas/pkg/types"
)

func (c *ContainerRunner) getContainerWithName(ctx context.Context, containerName string) (string, error) {
	filterArgs := filters.NewArgs()
	filterArgs.Add("name", containerName)
	containers, err := c.dockerCLI.ContainerList(ctx, dockertypes.ContainerListOptions{
		Filters: filterArgs,
		All:     true,
	})
	if err != nil {
		return "", err
	}
	if len(containers) == 0 {
		return "", nil
	}
	return containers[0].ID, nil
}

func (c *ContainerRunner) createContainer(ctx context.Context, input *types.ImageRunnerInput) (string, error) {
	containerName := input.Name()
	containerID, err := c.getContainerWithName(ctx, containerName)
	if err != nil {
		return "", err
	}
	if containerID != "" {
		return containerID, nil
	}
	createResponse, err := c.dockerCLI.ContainerCreate(ctx, input.ContainerConfig(c.defaultEnv), &container.HostConfig{
		Binds: []string{
			fmt.Sprintf("%s:%s", filepath.Join(c.configuration.DataDir, containerName), "/data"),
		},
	}, &network.NetworkingConfig{}, containerName)
	if err != nil {
		return "", err
	}
	return createResponse.ID, nil
}
