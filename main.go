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
	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "go-micro-api - ", log.LstdFlags)

	prods := handlers.NewProduct(logger)

	mux := mux.NewRouter()

	getRouter := mux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", prods.GetProducts)
	putRouter := mux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", prods.UpdateProducts)
	putRouter.Use(prods.MiddlewareHTTPHandler)
	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", prods.AddProducts)
	postRouter.Use(prods.MiddlewareHTTPHandler)

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
		logger.Println("Server Started listening on ", server.Addr)
	}()

	signalChannel := make(chan os.Signal, 4)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, syscall.SIGTERM)

	slg := <-signalChannel
	logger.Println("Received terminate, graceful shutdown", slg)
	tc, contextFunc := context.WithTimeout(context.Background(), 30*time.Second)
	defer contextFunc()
	server.Shutdown(tc)
}
