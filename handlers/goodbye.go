package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	log *log.Logger
}

func NewGoodbye(log *log.Logger) *Hello {
	return &Hello{log: log}
}

func (Goodbye *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	Goodbye.log.Println("Good bye!")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", data)
}
