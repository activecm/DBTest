package dbtest

import (
	"context"
	"testing"

	"github.com/activecm/dbtest/docker"
	"github.com/stretchr/testify/require"
)

func TestMongoDB(t *testing.T) {
	l, err := docker.NewLoader()
	require.Nil(t, err)
	mongo, err := NewMongoDBContainer(context.Background(), l)
	require.Nil(t, err)
	ssn, err := mongo.NewSession()
	require.Nil(t, err)
	require.NotNil(t, ssn)
	err = ssn.Ping()
	require.Nil(t, err)
	err = l.StopService(context.Background(), mongo)
	require.Nil(t, err)
}
