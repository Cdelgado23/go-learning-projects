package bootstrap

import (
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/path"
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/platform/server"
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/platform/storage/inmemory"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {

	pathFinder := path.NewPathFinderService(inmemory.NewInmemoryNodeRepository(10, 50))

	srv := server.New(host, port, pathFinder)
	return srv.Run()
}
