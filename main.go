package main

import (
	"log"
	"net/http"
	"os"

	"github.com/IAmSurajBobade/go_microservices/handlers"
)

func main() {
	logger := log.New(os.Stdout, "go-micro-api", log.LstdFlags)
	hw := handlers.NewHello(logger)
	gb := handlers.NewGoodbye(logger)

	mux := http.NewServeMux()
	mux.Handle("/", hw)
	mux.Handle("/goodbye", gb)

	http.ListenAndServe(":9090", mux)
}
