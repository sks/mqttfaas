package docker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/client"
)

//WaitForDaemonToBeRunning waits for daemon to be up and accepting connection
func WaitForDaemonToBeRunning(dockerClient *client.Client) error {
	retryAttempt := 0

	for {
		dockerVersion, err := dockerClient.ServerVersion(context.Background())
		if err == nil {
			log.Printf("Docker API version: %s, %s\n", dockerVersion.APIVersion, dockerVersion.Version)
			return nil
		}
		log.Printf("Error connecting to docker daemon, Sleeping for 5 seconds: %s.", err)
		time.Sleep(5 * time.Second)
		retryAttempt++
		if retryAttempt > 10 {
			return fmt.Errorf("Error with Docker client: %s", err.Error())
		}
	}
}
