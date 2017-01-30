package main

import (
	"fmt"
	"log"

	docker "github.com/fsouza/go-dockerclient"
	uuid "github.com/satori/go.uuid"
)

// code extracted from https://github.com/gofn/docker-in-docker
func main() {
	// connect to docker client
	// this code will run inside the container, but uses host socket to
	// create new containers
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	// create containers with random names
	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Name: fmt.Sprintf("gofn-%s", uuid.NewV4().String()),
		Config: &docker.Config{
			Image:     "debian:8",
			StdinOnce: true,
			OpenStdin: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// if nothing wrong happens, print created container inside host
	fmt.Println(container)
	fmt.Println("Container created!")
}
