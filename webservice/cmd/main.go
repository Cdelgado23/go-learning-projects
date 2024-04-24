package main

import (
	"github.com/cdelgado23/go-learning-projects/webservice/internal/platform/server"
	"log"
)

const (
	host = "localhost"
	port = 8080
)

func main() {
	srv := server.New(host, port)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
