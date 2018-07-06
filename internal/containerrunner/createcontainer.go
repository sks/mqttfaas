package containerrunner

import (
	"context"
	"fmt"

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

	createResponse, err := c.dockerCLI.ContainerCreate(ctx, &container.Config{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Image:        input.ImageID,
		OpenStdin:    true,
		Tty:          false,
		StdinOnce:    true,
		Env: []string{
			fmt.Sprintf("FIRED_BY=%s", input.Topic),
		},
	}, &container.HostConfig{}, &network.NetworkingConfig{}, containerName)
	if err != nil {
		return "", err
	}
	return createResponse.ID, nil
}
