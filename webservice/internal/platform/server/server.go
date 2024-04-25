package server

import (
	"fmt"
	"github.com/cdelgado23/go-learning-projects/webservice/internal/platform/server/handler/hello"
	"github.com/cdelgado23/go-learning-projects/webservice/internal/platform/server/handler/paths"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Print("Starting web service on ", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s Server) registerRoutes() {
	s.engine.GET("/hello", hello.HelloHandler())
	s.engine.GET("/paths", paths.GetPathsHandler())
}
