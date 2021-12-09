package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/IAmSurajBobade/go_microservices/data"
)

type Products struct {
	log *log.Logger
}

func NewProduct(logger *log.Logger) *Products {
	return &Products{log: logger}
}

func (product *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		product.getProducts(rw, r)
		return
	case http.MethodPost:
		product.addProducts(rw, r)
		return
	case http.MethodPut:
		reg := regexp.MustCompile(`/([0-9]+)`)
		matches := reg.FindAllStringSubmatch(r.URL.Path, -1)
		//fmt.Println(matches, len(matches), len(matches[0]))
		if len(matches) > 1 || len(matches[0]) != 2 {
			http.Error(rw, "Invalid input", http.StatusBadRequest)
			return
		}
		id, _ := strconv.Atoi(matches[0][1])
		//fmt.Printf("Id: %#v", id)
		product.updateProducts(id, rw, r)
		return
	}
	rw.WriteHeader(http.StatusNotImplemented)
}

func (product *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	product.log.Println("Handling GET request")
	prods := data.GetProducts()

	err := prods.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}

func (product *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	product.log.Println("Handling POST request")
	prods := &data.Product{}
	err := prods.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusBadRequest)
		return
	}
	data.AddProducts(*prods)
}

func (product *Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	product.log.Println("Handling PUT request")
	prods := &data.Product{}
	if err := prods.FromJSON(r.Body); err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusBadRequest)
		return
	}
	err := data.UpdateProducts(id, *prods)
	if err != data.ErrProdNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Error occurred", http.StatusInternalServerError)
		return
	}
}
