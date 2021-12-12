// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/IAmSurajBobade/go_microservices/product-api/data"
	"github.com/gorilla/mux"
)

type Products struct {
	log *log.Logger
}

func NewProduct(logger *log.Logger) *Products {
	return &Products{log: logger}
}

func (product *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	product.log.Println("Handling GET request")
	prods := data.GetProducts()

	err := prods.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (product *Products) AddProducts(rw http.ResponseWriter, r *http.Request) {
	product.log.Println("Handling POST request")
	prods := r.Context().Value(KeyValue{}).(data.Product)
	data.AddProducts(prods)
}

func (product *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	val := mux.Vars(r)
	id, err := strconv.Atoi(val["id"])
	if err != nil {
		http.Error(rw, "Invalid ID provided", http.StatusBadRequest)
		return
	}
	product.log.Println("Handling PUT request for ID: ", id)
	prods := r.Context().Value(KeyValue{}).(data.Product)
	err = data.UpdateProducts(id, prods)
	if err == data.ErrProdNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Error occurred", http.StatusInternalServerError)
		return
	}
}

type KeyValue struct{}

func (prod Products) MiddlewareHTTPHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prods := data.Product{}
		if err := prods.FromJSON(r.Body); err != nil {
			prod.log.Println("[Error] deserializing product", err)
			http.Error(rw, "Unable to marshal JSON", http.StatusBadRequest)
			return
		}
		err := prods.Validate()
		if err != nil {
			prod.log.Println("[Error] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product %s", err),
				http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyValue{}, prods)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
