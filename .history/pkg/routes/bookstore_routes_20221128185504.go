package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods(http.MethodPost)
	router.HandleFunc("/book", controllers.GetBooks).Methods(http.MethodGet)
}
