![gofn](docs/assets/logo.png)

[![Build Status](https://travis-ci.org/nuveo/gofn.svg?branch=master)](https://travis-ci.org/nuveo/gofn)
[![GoDoc](https://godoc.org/github.com/nuveo/gofn?status.png)](https://godoc.org/github.com/nuveo/gofn)
[![Go Report Card](https://goreportcard.com/badge/github.com/nuveo/gofn)](https://goreportcard.com/report/github.com/nuveo/gofn)
[![codecov](https://codecov.io/gh/nuveo/gofn/branch/master/graph/badge.svg)](https://codecov.io/gh/nuveo/gofn)
[![Join the chat at https://gitter.im/nuveo/gofn](https://badges.gitter.im/nuveo/gofn.svg)](https://gitter.im/nuveo/gofn?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Function process via container provider

## Premise

- Software makes a task
- After processing it dies
- Must put on stdout (print)

## Install

```bash
go get github.com/nuveo/gofn
```

## Example

```go
package main

import (
	"flag"
	"fmt"
	"log"

	"sync"

	"github.com/nuveo/gofn"
	"github.com/nuveo/gofn/iaas/digitalocean"
	"github.com/nuveo/gofn/provision"
)

const parallels = 3

func main() {
	wait := &sync.WaitGroup{}
	contextDir := flag.String("contextDir", "./", "a string")
	dockerfile := flag.String("dockerfile", "Dockerfile", "a string")
	imageName := flag.String("imageName", "", "a string")
	remoteBuildURI := flag.String("remoteBuildURI", "", "a string")
	volumeSource := flag.String("volumeSource", "", "a string")
	volumeDestination := flag.String("volumeDestination", "", "a string")
	remoteBuild := flag.Bool("remoteBuild", false, "true or false")
	input := flag.String("input", "", "a string")
	flag.Parse()
	wait.Add(parallels)
	for i := 0; i < parallels; i++ {
		run(*contextDir, *dockerfile, *imageName, *remoteBuildURI, *volumeSource, *volumeDestination, wait, *remoteBuild, *input)
	}
	wait.Wait()
}

func run(contextDir, dockerfile, imageName, remoteBuildURI, volumeSource, volumeDestination string, wait *sync.WaitGroup, remote bool, input string) {
	buildOpts := &provision.BuildOptions{
		ContextDir: contextDir,
		Dockerfile: dockerfile,
		ImageName:  imageName,
		RemoteURI:  remoteBuildURI,
		StdIN:      input,
	}
	if remote {
		buildOpts.Iaas = &digitalocean.Digitalocean{}
	}
	go func() {
		defer wait.Done()
		stdout, stderr, err := gofn.Run(buildOpts,
			&provision.VolumeOptions{
				Source:      volumeSource,
				Destination: volumeDestination,
			})
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Stderr: ", stderr)
		fmt.Println("Stdout: ", stdout)
	}()
}
```

### Run Example

```bash
	cd examples
	go run main.go -contextDir=testDocker -imageName=python -dockerfile=Dockerfile
	# or using volume
	go run main.go -contextDir=testDocker -imageName=python -dockerfile=Dockerfile -volumeSource=/tmp -volumeDestination=/tmp
	# or using remote Dockerfile
	go run main.go -remoteBuildURI=https://github.com/gofn/dockerfile-python-example.git -imageName="pythonexample"
	# you can also send a string that will be written to the stdin of the container
	go run main.go -contextDir=testDocker -imageName=python -dockerfile=Dockerfile -input "input string"
    # or run in digital ocean
    export DIGITALOCEAN_API_KEY="paste your key here"
    go run main.go -contextDir=testDocker -imageName=python -dockerfile=Dockerfile -remoteBuild=true
```

You can also compile with _go build_ or build and install with _go install_ command then run it as a native executable.

### Example Parameters

- -contextDir is the root directory where the Dockerfile, scripts, and other container dependencies are, by default current directory "./".

- -imageName is the name of the image you want to start, if it does not exist it will be automatically generated and if it exists the system will just start the container.

- -dockerFile is the name of the file containing the container settings, by default Dockerfile

- -volumeSource is the directory that will be mounted as a data volume. By default is empty string indicating his not used.

- -volumeDestination is the path mounted inside the container. By default is empty string indicating his  not used but if only omitted, volumeSource is used.

- -remoteBuildURI is remote URI containing the Dockerfile to build.By default is empty.
More details on [docker api docs](https://docs.docker.com/engine/reference/commandline/build/#/git-repositories)

- remoteBuild is a boolean that indicates if have to run localally or in a machine in digital ocean
Don't forget to export your api key.

- -input is a string that will be written to the stdin of the container

- -h Shows the list of parameters

gofn generates the images with "gofn/" as a prefix.
