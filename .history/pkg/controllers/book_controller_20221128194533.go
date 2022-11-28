package controllers

import (
	"go-bookstore/pkg/models"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	models.CreateBook(&models.Book{
		Name:        "Hello",
		Author:      "Rishi",
		Publication: "Postric",
	})
}

func GetBook(w http.ResponseWriter, r *http.Request) {}

func UpdateBook(w http.ResponseWriter, r *http.Request) {}

func DeleteBook(w http.ResponseWriter, r *http.Request) {}
