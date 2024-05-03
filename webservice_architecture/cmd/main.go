package main

import (
	"github.com/cdelgado23/go-learning-projects/webservice-architecture/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
