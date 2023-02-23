package config

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/philolo1/book/pkg/env"
  "fmt"
)

var (
  db * gorm.DB
)

func Connect() {
  d, err := gorm.Open("mysql", env.GetEnv().DatabaseConnectionString)
  if err != nil {
    panic(err)
  }
  fmt.Println("Successfully Connected to database")
  db = d
}

func GetDB() *gorm.DB {
  return db
}
