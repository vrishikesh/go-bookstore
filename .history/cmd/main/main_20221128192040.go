package main

import (
	"github.com/gorilla/mux"

	"go-bookstore/pkg/routes"
)

func main()  {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
}
