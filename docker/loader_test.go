package docker

import (
	"context"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/require"
)

func TestNewLoader(t *testing.T) {
	cli, err := client.NewEnvClient()
	require.Nil(t, err)
	require.NotNil(t, cli)
	l := NewLoaderFromClient(cli)
	require.NotNil(t, l.client)
}

func TestStartHelloWorld(t *testing.T) {
	loader, err := NewLoader()
	require.Nil(t, err)
	helloWorldContainerSpec := container.Config{
		Image: "hello-world",
	}
	helloWorldSpec := ServiceSpecification{
		FullyQualifiedImageName: "library/hello-world",
		ContainerConfig:         &helloWorldContainerSpec,
	}
	service, err := loader.StartService(context.Background(), &helloWorldSpec)
	require.Nil(t, err)
	require.NotNil(t, service.container)
	require.True(t, len(service.GetID()) > 0)
}

func TestStartStopService(t *testing.T) {
	loader, err := NewLoader()
	require.Nil(t, err)
	mongoContainerSpec := container.Config{Image: "mongo:latest"}
	mongoSpec := ServiceSpecification{
		FullyQualifiedImageName: "library/mongo",
		ContainerConfig:         &mongoContainerSpec,
	}
	service, err := loader.StartService(context.Background(), &mongoSpec)
	require.Nil(t, err)
	require.NotNil(t, service.container)
	err = loader.StopService(context.Background(), service)
	require.Nil(t, err)

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Quiet: true,
		All:   true,
	})

	require.Nil(t, err)
	for i := range containers {
		if containers[i].ID == service.GetID() {
			require.FailNow(t, "Service not removed after StopService")
		}
	}
}
