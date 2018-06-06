package dbtest

import (
	"context"

	"github.com/activecm/dbtest/docker"
	"github.com/docker/docker/api/types/container"
	mgo "gopkg.in/mgo.v2"
)

//MongoDBContainer extends docker.Service to provide
//on demand MongoDB connections
type MongoDBContainer struct {
	docker.Service
}

//NewMongoDBContainer starts a new instance of MongoDB
//with Docker
func NewMongoDBContainer(ctx context.Context, loader docker.Loader) (MongoDBContainer, error) {
	mongoOut := MongoDBContainer{}
	mongoContainerSpec := container.Config{Image: "mongo:latest"}
	mongoSpec := docker.ServiceSpecification{
		FullyQualifiedImageName: "library/mongo",
		ContainerConfig:         &mongoContainerSpec,
	}
	var err error
	mongoOut.Service, err = loader.StartService(ctx, &mongoSpec)
	if err != nil {
		return mongoOut, err
	}

	err = WaitForTCP(ctx, mongoOut.getSocketAddress())
	return mongoOut, err
}

//NewSession returns a new mgo.Session object dialed into to this server
func (m MongoDBContainer) NewSession() (*mgo.Session, error) {
	return mgo.Dial(m.GetMongoDBURI())
}

//GetMongoDBURI returns the MongoDB URI for connecting to the container
func (m MongoDBContainer) GetMongoDBURI() string {
	return "mongodb://" + m.getSocketAddress()
}

func (m MongoDBContainer) getSocketAddress() string {
	return m.GetIPAddress() + ":27017"
}
