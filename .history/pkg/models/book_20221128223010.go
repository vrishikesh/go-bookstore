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

func CreateBook(b *Book) (*Book, error) {
	g := db.Create(b)
	if g.Error != nil {
		return nil, fmt.Errorf("unable to insert book: %w", g.Error)
	}
	return b, nil
}

func GetBooks() ([]Book, error) {
	var books []Book
	g := db.Find(&books)
	if g.Error != nil {
		return nil, fmt.Errorf("unable to insert book: %w", g.Error)
	}
	return books, nil
}

func GetBookById(id int64) (*Book, error) {
	var book Book
	g := db.Find(&book, id)
	if g.Error != nil {
		return nil, fmt.Errorf("unable to insert book: %w", g.Error)
	}
	return &book, nil
}

func (b *Book) Update(bk *Book) error {
	g := db.Model(b).Updates(bk)
	if g.Error != nil {
		return fmt.Errorf("unable to update book: %w", g.Error)
	}
	return nil
}

func (b *Book) Delete() error {
	g := db.Delete(b.ID)
	if g.Error != nil {
		return fmt.Errorf("unable to insert book: %w", g.Error)
	}
	return nil
}
