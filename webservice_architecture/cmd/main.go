package main

import (
	"github.com/cdelgado23/go-learning-projects/webservice/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
