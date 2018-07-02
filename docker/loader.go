package docker

import (
	"context"
	"io"
	"io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

//Loader wraps a docker Client and exposes a simpler interface
//for creating ephemeral Services
type Loader struct {
	client *client.Client
}

//NewLoaderFromClient creates a Loader from a docker Client
func NewLoaderFromClient(dockerClient *client.Client) Loader {
	return Loader{
		client: dockerClient,
	}
}

//NewLoader creates a Docker Client and returns a new Loader
func NewLoader() (Loader, error) {
	loader := Loader{}
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		return loader, err
	}
	loader.client = dockerClient
	return loader, nil
}

//StartService creates a new Service from a given service specification.
//The process is cancellable with a context.
func (l Loader) StartService(ctx context.Context, spec IServiceSpecification) (Service, error) {
	outService := Service{}
	reader, err := l.client.ImagePull(ctx, spec.GetFullyQualifiedImageName(), types.ImagePullOptions{})
	if err != nil {
		return outService, err
	}

	//I don't understand why this works:
	io.Copy(ioutil.Discard, reader)
	//but this doesn't:
	//reader.Close()

	hostConfigPtr := spec.GetHostConfig()
	if hostConfigPtr == nil {
		hostConfig := container.HostConfig{
			AutoRemove: true,
		}
		hostConfigPtr = &hostConfig
	} else {
		hostConfigPtr.AutoRemove = true
	}
	resp, err := l.client.ContainerCreate(ctx,
		spec.GetContainerConfig(),
		hostConfigPtr,
		nil,
		"",
	)
	if err != nil {
		return outService, err
	}
	err = l.client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return outService, err
	}

	outService.container, err = l.client.ContainerInspect(ctx, resp.ID)
	if err != nil {
		return outService, err
	}

	return outService, nil
}

//StopService stops and removes a Service. Blocks until container is removed.
func (l Loader) StopService(ctx context.Context, service IService) error {
	err := l.client.ContainerStop(ctx, service.GetID(), nil)
	if err != nil {
		return err
	}
	statusCh, errCh := l.client.ContainerWait(ctx, service.GetID(), container.WaitConditionRemoved)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err = <-errCh:
		return err
	case <-statusCh:
	}
	return nil
}

//Close closes the socket used to contact the Docker API
func (l Loader) Close() error {
	return l.client.Close()
}
