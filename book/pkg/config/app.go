package config

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "fmt"
  "os"
)

var (
  db * gorm.DB
)

func Connect() {
  d, err := gorm.Open("mysql", os.Getenv("DATABASE_CONNECTION_STRING"))
  if err != nil {
    panic(err)
  }
  fmt.Println("Successfully Connected to database")
  db = d
}

func GetDB() *gorm.DB {
  return db
}
