package models

import (
	"Book_api/pkg/config"

	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	config.ConnectDB()
	Db = config.GetDB()
	Db.AutoMigrate(&Book{})
}

type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// func (b *Book) CreateBook() *Book {
// 	db.Create(&b)
// 	return b
// }

// func GetAllBook() []Book {
// 	var books []Book
// 	db.Find(&books)
// 	return books
// }

// func GetBookById(Id int64) Book {
// 	var book Book
// 	db.First(Id).Find(&book)
// 	return book
// }

// func DeleteBook(Id int64) Book {
// 	var book Book
// 	db.First(Id).Delete(&book)
// 	return book
// }
