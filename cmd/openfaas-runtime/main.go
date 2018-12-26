package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/sks/mqttfaas/internal/docker"
)

func main() {
	log.Println("Starting the process")
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("could not create docker client :%s", err)
	}
	err = docker.WaitForDaemonToBeRunning(dockerClient)
	if err != nil {
		log.Fatalf("docker did not come up after so much time: %s", err)
	}
	err = initializeSwarm(dockerClient)
	if err != nil {
		if !strings.Contains(err.Error(), "This node is already part of a swarm") {
			log.Fatalf("Failed to initialize swarm mode: %s", err)
		}
	}

	err = startFAASSwarm(dockerClient)
	if err != nil {
		log.Fatalf("Error loading the faas swarm :%s\n", err)
	}
	select {}
}

func initializeSwarm(dockerClient *client.Client) error {
	_, err := dockerClient.SwarmInit(context.Background(), swarm.InitRequest{
		ListenAddr:    "0.0.0.0",
		AdvertiseAddr: "",
		Spec: swarm.Spec{
			Orchestration: swarm.OrchestrationConfig{
				TaskHistoryRetentionLimit: func(i int64) *int64 { return &i }(1),
			},
		},
	})
	return err
}

func startFAASSwarm(dockerClient *client.Client) error {
	faasImageTargz, err := os.Open("/var/data/faasswarm.tar.gz")
	if err != nil {
		return err
	}
	_, err = dockerClient.ImageLoad(context.Background(), faasImageTargz, false)
	if err != nil {
		return err
	}
	return command("docker", "stack", "deploy", "-c", "/var/data/docker-compose.yml", "faas-swarm").Run()
}

//command execute the command
func command(name string, args ...string) *exec.Cmd {
	log.Printf(`Command "%s %s"\n`, name, strings.Join(args, " "))
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd
}
