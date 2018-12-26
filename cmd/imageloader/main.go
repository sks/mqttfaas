package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/docker/docker/client"
	"github.com/sks/mqttfaas/internal/docker"
)

func main() {
	log.Println("Starting the process")
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		log.Fatalf("Error with Docker client: %s.", err.Error())
	}

	err = docker.WaitForDaemonToBeRunning(dockerClient)
	if err != nil {
		log.Fatalf("docker did not come up after so much time: %s", err)
	}

	http.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		file, header, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("Loading the file %s\n", header.Filename)
		defer file.Close()
		resp, err := dockerClient.ImageLoad(r.Context(), file, false)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		d, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(d)
	})
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	log.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
