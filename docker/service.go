package docker

import "github.com/docker/docker/api/types"

//IService represents an ephemeral Docker container.
//It mainly serves to simplify the Docker container API.
type IService interface {
	GetIPAddress() string
	GetID() string
}

//Service provides a concrete implementation for IService.
//Service may be embedded into specific Container types
//to implement IService
type Service struct {
	container types.ContainerJSON
}

//GetIPAddress returns the Service's internal IP Address
//accessible from the Docker host.
func (s Service) GetIPAddress() string {
	return s.container.NetworkSettings.IPAddress
}

//GetID returns the Service's Docker ID
func (s Service) GetID() string {
	return s.container.ID
}
