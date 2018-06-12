package dbtest

import "github.com/docker/distribution"
import "github.com/docker/go-units"
import "github.com/docker/go-connections/sockets"
import "github.com/pkg/errors"

//privateDummyFunc "activates" the [[constraint]]s needed to override
//the dependencies for the docker library.
//[[override]] clauses work to control transient dependencies
//for this project, but if someone imports dbtest,
//only [[constraints]] are percolated up.
func privateDummyFunc() {
	_ = distribution.Descriptor{}
	_ = units.GB
	_ = sockets.ErrProtocolNotAvailable
	_ = errors.New("A")
}
