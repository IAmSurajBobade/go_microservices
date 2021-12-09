package handlers

import (
	"log"
	"net/http"

	"github.com/IAmSurajBobade/go_microservices/data"
)

type Products struct {
	log *log.Logger
}

func NewProduct(logger *log.Logger) *Products {
	return &Products{log: logger}
}

func (product *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		product.getProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusNotImplemented)
}

func (product *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	prods := data.GetProducts()

	err := prods.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
