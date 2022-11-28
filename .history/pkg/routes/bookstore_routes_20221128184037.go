package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {})
}