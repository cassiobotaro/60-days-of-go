package main

import (
	"io"
	"log"
	"os"

	"github.com/fsouza/go-dockerclient"
)

// Run creates and start a container ataching writers to stdout and stderr
func Run(client *docker.Client, stdout, stderr io.Writer) string {
	// create a container
	container, err := client.CreateContainer(docker.CreateContainerOptions{
		Name: "teste-hello",
		Config: &docker.Config{
			Cmd:       []string{"/bin/bash", "-c", "for i in {1..10}; do echo $i; sleep 2; done"},
			Image:     "debian:8",
			StdinOnce: true,
			OpenStdin: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	// start the created container
	err = client.StartContainer(container.ID, nil)
	if err != nil {
		log.Fatal("StartContainer:", err)
	}
	// attach to container binding writers to stdout and stderr
	// Note: logs are true because we can lost some computation done before attach
	_, err = client.AttachToContainerNonBlocking(docker.AttachToContainerOptions{
		Container:    container.ID,
		RawTerminal:  true,
		Stream:       true,
		Stdin:        true,
		Stdout:       true,
		Stderr:       true,
		Logs:         true,
		OutputStream: stdout,
		ErrorStream:  stderr,
	})
	// return container id
	return container.ID
}

func main() {
	// connect to local socket
	endpoint := "unix:///var/run/docker.sock"
	client, err := docker.NewClient(endpoint)
	if err != nil {
		panic(err)
	}
	id := Run(client, os.Stdout, os.Stderr)
	// wait computation inside container
	_, err = client.WaitContainer(id)
	if err != nil {
		return
	}
}
