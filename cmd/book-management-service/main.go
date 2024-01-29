package main

import (
	"log"
	"net/http"

	"github.com/Max-Gabriel-Susman/book-management-service/internal/routes"
	"github.com/Max-Gabriel-Susman/book-management-service/internal/store"
	"github.com/gorilla/mux"
)

func main() {
	store.Init()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
