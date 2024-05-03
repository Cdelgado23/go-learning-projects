package server

import (
	"fmt"
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/path"
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/internal/platform/server/handler"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr   string
	engine     *gin.Engine
	pathFinder path.PathFinderService
}

func New(host string, port uint, pathFinder path.PathFinderService) Server {
	srv := Server{
		httpAddr:   fmt.Sprintf("%s:%d", host, port),
		engine:     gin.New(),
		pathFinder: pathFinder,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Print("Starting web service on ", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s Server) registerRoutes() {
	s.engine.GET("/paths", handler.GetPathsHandler(s.pathFinder))
}