package models

import (
	"crud/pkg/config"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var validate *validator.Validate

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name" validate:"required"`
	Author      string `json:"author" validate:"required"`
	Publication string `json:"publication" validate:"required"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// funsi rute
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks(q string) []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
