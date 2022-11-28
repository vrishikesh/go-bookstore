package routes

import (
	"github.com/gorilla/mux"

	"go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	// router.HandleFunc("/")
}