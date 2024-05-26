package bootstrap

import (
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/node/application"
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/platform/server"
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/platform/storage/inmemory"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {

	pathFinder := node.NewPathFinderService(inmemory.NewInMemoryNodeRepository(10, 50))

	srv := server.New(host, port, pathFinder)
	return srv.Run()
}
