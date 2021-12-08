package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/IAmSurajBobade/go_microservices/handlers"
)

func main() {
	logger := log.New(os.Stdout, "go-micro-api - ", log.LstdFlags)
	hw := handlers.NewHello(logger)
	gb := handlers.NewGoodbye(logger)

	mux := http.NewServeMux()
	mux.Handle("/goodbye", gb)
	mux.Handle("/hello", hw)

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

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	slg := <-signalChannel
	logger.Println("Received terminate, graceful shutdown", slg)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)
}
