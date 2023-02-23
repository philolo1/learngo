package models

import (
  "github.com/jinzhu/gorm"
  "github.com/philolo1/book/pkg/config"
)

var db *gorm.DB

type Book struct {
  gorm.Model
  Name string `json:"name"`
  Author string `json:"author"`
  Publication string `json:"publication"`
}

func InitializeDB() {
  config.Connect()
  db = config.GetDB()
  db.AutoMigrate(&Book{})
}

func CreateBook(b *Book) *Book {
  db.NewRecord(b)
  db.Create(&b)
  return b
}

func GetAllBooks() []Book {
  var books []Book
  db.Find(&books)
  return books
}

func GetBookById(Id int64) *Book {
  var book Book
  db.Find(&book, Id)
  return &book
}

func DeleteBookById(Id int64) *gorm.DB {
  var book Book
  return db.Delete(book, Id)
}

func DeleteBooks() {
  db.Delete(&Book{})
}
