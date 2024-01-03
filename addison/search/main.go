package main

import (
	"log"
	"net/http"
	"search/resources"

	"github.com/gorilla/mux"
)

func main() {
	resources.Init() // call the Init function and set the token

	r := mux.NewRouter()
	r.Handle("/search", resources.Router())
	log.Fatal(http.ListenAndServe(":3001", r))
}
