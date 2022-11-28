package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/books", controllers.GetBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", controllers.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods(http.MethodDelete)
}
