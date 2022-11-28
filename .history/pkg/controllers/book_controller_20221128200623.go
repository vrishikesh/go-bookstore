package controllers

import (
	"encoding/json"
	"go-bookstore/pkg/models"
	"log"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	b, err := models.CreateBook(&models.Book{
		Name:        "Hello",
		Author:      "Rishi",
		Publication: "Postric",
	})
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	book, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {

}

func GetBook(w http.ResponseWriter, r *http.Request) {}

func UpdateBook(w http.ResponseWriter, r *http.Request) {}

func DeleteBook(w http.ResponseWriter, r *http.Request) {}
