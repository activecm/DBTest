package docker

import (
	"github.com/docker/docker/api/types/container"
)

//IServiceSpecification defines the necessary data for creating a Service
type IServiceSpecification interface {
	GetFullyQualifiedImageName() string
	GetContainerConfig() *container.Config
	GetHostConfig() *container.HostConfig
}

//ServiceSpecification provides a simple implementation of IServiceSpecification
type ServiceSpecification struct {
	FullyQualifiedImageName string
	ContainerConfig         *container.Config
	HostConfig              *container.HostConfig
}

//GetFullyQualifiedImageName returns an image's fully qualified name
//i.e. "library/hello-world"
func (s *ServiceSpecification) GetFullyQualifiedImageName() string {
	return s.FullyQualifiedImageName
}

//GetContainerConfig returns the specification's container configuration.
//This contains host independant configuration details.
func (s *ServiceSpecification) GetContainerConfig() *container.Config {
	return s.ContainerConfig
}

//GetHostConfig returns the specification's host dependant configuration
//such as bind mounts.
func (s *ServiceSpecification) GetHostConfig() *container.HostConfig {
	return s.HostConfig
}
