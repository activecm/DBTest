# DBTest: Go Integration Testing Via Docker

### DBTest Allows You to Test Against a Real Database

Using the Docker API, your Go program can manage its own
testing dependencies.

This is especially useful when writing integration tests
(or when it's simply impossible to mock out a dependency).

## MongoDB

```go
//Create a new Loader. This may error is Docker is not accessible
l, err := docker.NewLoader()
//Create a new MongoDB container using the loader
mongo, err := NewMongoDBContainer(context.Background(), l)
//Obtain an mgo.Session to contact MongoDB
ssn, err := mongo.NewSession()
//Do something with the Session
err = ssn.Ping()
//Stop and remove the MongoDB container
err = l.StopService(context.Background(), mongo)
_ = err
```
