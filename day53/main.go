package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/nuveo/gofn"
	"github.com/nuveo/gofn/provision"
)

const parallels = 10

func run(imageName string, remoteBuildURI string, wait *sync.WaitGroup) {
	buildOpts := &provision.BuildOptions{
		ImageName: imageName,
		RemoteURI: remoteBuildURI,
	}
	go func() {
		defer wait.Done()
		stdout, stderr, err := gofn.Run(buildOpts, nil)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Stderr: ", stderr)
		fmt.Println("Stdout: ", stdout)
	}()
}

func main() {
	wait := &sync.WaitGroup{}
	wait.Add(parallels)
	for i := 0; i < parallels; i++ {
		run("gofn/sample", "https://github.com/gofn/dockerfile-python-example.git", wait)
	}
	wait.Wait()
}
