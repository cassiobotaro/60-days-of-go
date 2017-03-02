package provision

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/nuveo/gofn/iaas"
	uuid "github.com/satori/go.uuid"
)

var (
	// ErrImageNotFound is raised when image is not found
	ErrImageNotFound = errors.New("provision: image not found")

	// ErrContainerNotFound is raised when image is not found
	ErrContainerNotFound = errors.New("provision: container not found")

	// Input receives a string that will be written to the stdin of the container in function FnRun
	Input string
)

// VolumeOptions are options to mount a host directory as data volume
type VolumeOptions struct {
	Source      string
	Destination string
}

// BuildOptions are options used in the image build
type BuildOptions struct {
	ContextDir              string
	Dockerfile              string
	DoNotUsePrefixImageName bool
	ImageName               string
	RemoteURI               string
	StdIN                   string
	Iaas                    iaas.Iaas
}

func (opts BuildOptions) GetImageName() string {
	if opts.DoNotUsePrefixImageName {
		return opts.ImageName
	}
	return "gofn/" + opts.ImageName
}

// FnClient instantiate a docker client
func FnClient(endPoint string) (client *docker.Client, err error) {
	if endPoint == "" {
		endPoint = "unix:///var/run/docker.sock"
	}

	client, err = docker.NewClient(endPoint)
	return
}

// FnRemove remove container
func FnRemove(client *docker.Client, containerID string) (err error) {
	err = client.RemoveContainer(docker.RemoveContainerOptions{ID: containerID, Force: true})
	return
}

// FnContainer create container
func FnContainer(client *docker.Client, image, volume string) (container *docker.Container, err error) {
	binds := []string{}
	if volume != "" {
		binds = append(binds, volume)
	}
	container, err = client.CreateContainer(docker.CreateContainerOptions{
		Name:       fmt.Sprintf("gofn-%s", uuid.NewV4().String()),
		HostConfig: &docker.HostConfig{Binds: binds},
		Config: &docker.Config{
			Image:     image,
			StdinOnce: true,
			OpenStdin: true,
		},
	})
	return
}

// FnImageBuild builds an image
func FnImageBuild(client *docker.Client, opts *BuildOptions) (Name string, Stdout *bytes.Buffer, err error) {
	if opts.Dockerfile == "" {
		opts.Dockerfile = "Dockerfile"
	}
	stdout := new(bytes.Buffer)
	Name = opts.GetImageName()
	err = client.BuildImage(docker.BuildImageOptions{
		Name:           Name,
		Dockerfile:     opts.Dockerfile,
		SuppressOutput: true,
		OutputStream:   stdout,
		ContextDir:     opts.ContextDir,
		Remote:         opts.RemoteURI,
	})
	if err != nil {
		return
	}
	Stdout = stdout
	return
}

// FnFindImage returns image data by name
func FnFindImage(client *docker.Client, imageName string) (image docker.APIImages, err error) {
	var imgs []docker.APIImages
	imgs, err = client.ListImages(docker.ListImagesOptions{Filter: imageName})
	if err != nil {
		return
	}
	if len(imgs) == 0 {
		err = ErrImageNotFound
		return
	}
	image = imgs[0]
	return
}

// FnFindContainer return container by image name
func FnFindContainer(client *docker.Client, imageName string) (container docker.APIContainers, err error) {
	var containers []docker.APIContainers
	containers, err = client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		return
	}

	if !strings.HasPrefix(imageName, "gofn") {
		imageName = "gofn/" + imageName
	}

	for _, v := range containers {
		if v.Image == imageName {
			container = v
			return
		}
	}
	err = ErrContainerNotFound
	return
}

// FnKillContainer kill the container
func FnKillContainer(client *docker.Client, containerID string) (err error) {
	err = client.KillContainer(docker.KillContainerOptions{ID: containerID})
	return
}

//FnAttach attach into a running container
func FnAttach(client *docker.Client, containerID string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (w docker.CloseWaiter, err error) {
	return client.AttachToContainerNonBlocking(docker.AttachToContainerOptions{
		Container:    containerID,
		RawTerminal:  true,
		Stream:       true,
		Stdin:        true,
		Stderr:       true,
		Stdout:       true,
		Logs:         true,
		InputStream:  stdin,
		ErrorStream:  stderr,
		OutputStream: stdout,
	})
}

// FnStart start the container
func FnStart(client *docker.Client, containerID string) error {
	return client.StartContainer(containerID, nil)
}

// FnRun runs the container
func FnRun(client *docker.Client, containerID, input string) (Stdout *bytes.Buffer, Stderr *bytes.Buffer, err error) {
	err = FnStart(client, containerID)
	if err != nil {
		return
	}

	// attach to write input
	_, err = FnAttach(client, containerID, strings.NewReader(input), nil, nil)
	if err != nil {
		return
	}

	done, errors := FnWaitContainer(client, containerID)
	select {
	case err = <-errors:
		return
	case <-done:
	}

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	FnLogs(client, containerID, stdout, stderr)

	Stdout = stdout
	Stderr = stderr
	return
}

// FnLogs logs all container activity
func FnLogs(client *docker.Client, containerID string, stdout io.Writer, stderr io.Writer) error {
	return client.Logs(docker.LogsOptions{
		Container:    containerID,
		Stdout:       true,
		Stderr:       true,
		ErrorStream:  stderr,
		OutputStream: stdout,
	})
}

// FnWaitContainer wait until container finnish your processing
func FnWaitContainer(client *docker.Client, containerID string) (chan bool, chan error) {
	done := make(chan bool)
	errors := make(chan error)
	go func() {
		_, err := client.WaitContainer(containerID)
		if err != nil {
			errors <- err
		}
		done <- true
	}()
	return done, errors
}

// FnConfigVolume set volume options
func FnConfigVolume(opts *VolumeOptions) string {
	if opts.Source == "" && opts.Destination == "" {
		return ""
	}
	if opts.Destination == "" {
		opts.Destination = opts.Source
	}
	return opts.Source + ":" + opts.Destination
}
