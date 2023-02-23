package models

import (
  "github.com/jinzhu/gorm"
  "github.com/philolo1/book/pkg/config"
  "github.com/redis/go-redis/v9"
  "encoding/json"
  "context"
  "fmt"
)

var db *gorm.DB
var rdb *redis.Client
var ctx = context.Background()

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
  rdb = config.GetRedis()
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
  redisKey := fmt.Sprintf("book:%d", Id)
  res, _ := rdb.Get(ctx, redisKey).Result()
  if (len(res) == 0) {
    db.Find(&book, Id)
    go SaveBookCache(redisKey, book)
  } else {
    json.Unmarshal([]byte(res), &book)
    // fmt.Println("res", res)
  }
  // cache
  return &book
}

func SaveBookCache(redisKey string, book Book) {
  if (len(book.Name) > 0) {
    jsonString, _ := json.Marshal(book)
    rdb.Set(ctx, redisKey, jsonString, 0)
  }
}



func DeleteBookById(Id int64) *gorm.DB {
  var book Book
  return db.Delete(book, Id)
}

func DeleteBooks() {
  db.Delete(&Book{})
}
