package main

import (
	"log"
	"net/http"
	"os"

	"github.com/IAmSurajBobade/go_microservices/handlers"
)

func main() {
	logger := log.New(os.Stdout, "awesome-go", log.Flags())
	hw := handlers.NewHello(logger)

	http.ListenAndServe(":9090", nil)
}
