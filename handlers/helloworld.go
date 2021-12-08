package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	log *log.Logger
}

func NewHello(log *log.Logger) *Hello {
	return &Hello{log: log}
}

func (hello *Hello) ServeHTTP(rw http.ResponseWriter, r http.Request) {
	hello.log.Println("Hello world!")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", data)
}
