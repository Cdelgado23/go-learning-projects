package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const httpPort = ":8080"

func main() {
	fmt.Println("Starting web service on port", httpPort)

	srv := gin.New()
	srv.GET("/hello", helloHandler)

	log.Fatal(srv.Run(httpPort))
}

func helloHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello, World!")
}
