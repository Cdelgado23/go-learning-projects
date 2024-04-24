package main

import (
	"fmt"
	"log"
	"net/http"
)

const httpPort = ":8080"

func main() {
	fmt.Println("Starting web service on port", httpPort)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(httpPort, mux))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
}
