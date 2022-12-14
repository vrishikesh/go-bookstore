package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-bookstore/pkg/models"
	"go-bookstore/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	b, err := models.CreateBook(&models.Book{
		Name:        "Hello",
		Author:      "Rishi",
		Publication: "Postric",
	})
	if err != nil {
		utils.ErrorJson(w, err)
		return
	}

	book, _ := json.Marshal(b)
	utils.Json(w, http.StatusCreated, book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	b, err := models.GetBooks()
	if err != nil {
		utils.ErrorJson(w, err)
		return
	}

	books, _ := json.Marshal(b)
	utils.Json(w, http.StatusOK, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	primaryId, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utils.ErrorJson(w, err)
		return
	}

	b, err := models.GetBookById(primaryId)
	if err != nil {
		utils.ErrorJson(w, err)
		return
	}

	book, _ := json.Marshal(b)
	utils.Json(w, http.StatusOK, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	primaryId, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utils.ErrorJson(w, err)
		return
	}

	b, err := models.GetBookById(primaryId)
	if err != nil {
		utils.ErrorJson(w, err)
		return
	}

	bk, err := b.Update(&models.Book{
		Author: "updated author",
	})
	if err != nil {
		utils.ErrorJson(w, err)
		return
	}

	book, _ := json.Marshal(bk)
	utils.Json(w, http.StatusAccepted, book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	primaryId, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		utils.ErrorJson(w, err)
		return
	}

	b, err := models.GetBookById(primaryId)
	if err != nil {
		utils.ErrorJson(w, err)
		return
	}

	err = b.Remove()
	if err != nil {
		utils.ErrorJson(w, err)
		return
	}

	utils.Json(w, http.StatusNoContent)
}
