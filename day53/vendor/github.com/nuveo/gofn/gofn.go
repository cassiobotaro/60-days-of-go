package gofn

import (
	"bytes"
	"io"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/nuveo/gofn/iaas"
	"github.com/nuveo/gofn/provision"
)

const dockerPort = ":2375"

// ProvideMachine provisioning a machine in the cloud
func ProvideMachine(service iaas.Iaas) (client *docker.Client, machine *iaas.Machine, err error) {
	machine, err = service.CreateMachine()
	if err != nil {
		if machine != nil {
			err = service.DeleteMachine(machine)
		}
		return
	}
	client, err = provision.FnClient(machine.IP + dockerPort)
	if err != nil {
		return
	}
	return
}

// PrepareContainer build an image if necessary and run the container
func PrepareContainer(client *docker.Client, buildOpts *provision.BuildOptions, volumeOpts *provision.VolumeOptions) (container *docker.Container, err error) {
	img, err := provision.FnFindImage(client, buildOpts.GetImageName())
	if err != nil && err != provision.ErrImageNotFound {
		return
	}

	var image string
	if img.ID == "" {
		image, _, err = provision.FnImageBuild(client, buildOpts)
		if err != nil {
			return
		}
	} else {
		image = buildOpts.GetImageName()
	}

	volume := ""
	if volumeOpts != nil {
		volume = provision.FnConfigVolume(volumeOpts)
	}

	container, err = provision.FnContainer(client, image, volume)
	if err != nil {
		return
	}
	return
}

// RunWait runs the conainer returning channels to control your status
func RunWait(client *docker.Client, container *docker.Container) (err error, running chan bool, errors chan error) {
	err = provision.FnStart(client, container.ID)
	if err != nil {
		return
	}
	running, errors = provision.FnWaitContainer(client, container.ID)
	return
}

// Attach allow to connect into a running container and interact using stdout, stderr and stdin
func Attach(client *docker.Client, container *docker.Container, stdin io.Reader, stdout io.Writer, stderr io.Writer) (docker.CloseWaiter, error) {
	return provision.FnAttach(client, container.ID, stdin, stdout, stderr)
}

// Run runs the designed image
func Run(buildOpts *provision.BuildOptions, volumeOpts *provision.VolumeOptions) (stdout string, stderr string, err error) {
	var client *docker.Client
	client, err = provision.FnClient("")
	if err != nil {
		return
	}

	if buildOpts.Iaas != nil {
		var machine *iaas.Machine
		client, machine, err = ProvideMachine(buildOpts.Iaas)
		if err != nil {
			return
		}
		defer buildOpts.Iaas.DeleteMachine(machine)
	}

	var container *docker.Container
	container, err = PrepareContainer(client, buildOpts, volumeOpts)

	var buffout *bytes.Buffer
	var bufferr *bytes.Buffer

	buffout, bufferr, err = provision.FnRun(client, container.ID, buildOpts.StdIN)
	if err != nil {
		return
	}
	stdout = buffout.String()
	stderr = bufferr.String()

	err = DestroyContainer(client, container)
	return

}

func DestroyContainer(client *docker.Client, container *docker.Container) error {
	return provision.FnRemove(client, container.ID)
}
