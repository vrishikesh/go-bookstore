package models

import (
	"fmt"

	"gorm.io/gorm"

	"go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	g := db.Create(b)
	if g.Error != nil {
		return nil, fmt.Errorf("unable to insert book: %w", g.Error)
	}
	return b, nil
}

func GetBooks() []Book {
	var books []Book
	db.Find(books)
	return books
}

func GetBookById(id int64) (*Book) {
	var books []Book
	db.Find(books)
	return books
}
 