package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IAmSurajBobade/go_microservices/handlers"
)

func main() {
	logger := log.New(os.Stdout, "go-micro-api - ", log.LstdFlags)

	prods := handlers.NewProduct(logger)

	mux := http.NewServeMux()

	mux.Handle("/", prods)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal, 4)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, syscall.SIGTERM)

	slg := <-signalChannel
	logger.Println("Received terminate, graceful shutdown", slg)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
