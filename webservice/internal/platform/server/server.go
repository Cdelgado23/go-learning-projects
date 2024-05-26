package server

import (
	"fmt"
	"github.com/cdelgado23/go-learning-projects/webservice/internal/path"
	handler "github.com/cdelgado23/go-learning-projects/webservice/internal/platform/server/handler"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr   string
	engine     *gin.Engine
	pathFinder path.PathFinderService
}

func New(host string, port uint) Server {
	srv := Server{
		httpAddr:   fmt.Sprintf("%s:%d", host, port),
		engine:     gin.New(),
		pathFinder: path.NewPathFinderService(10, 50),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Print("Starting web application on ", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s Server) registerRoutes() {
	s.engine.GET("/paths", handler.GetPathsHandler(s.pathFinder))
}
