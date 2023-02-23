package config

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/philolo1/book/pkg/env"
  "fmt"
  "time"
  "sync"
  "context"
  "github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var (
  db * gorm.DB
  rdb * redis.Client
)

func printTime() {
  currentTime := time.Now()
  fmt.Println("The current time is:", currentTime)
}

func ConnectMySql(wg *sync.WaitGroup) {
  d, err := gorm.Open("mysql", env.GetEnv().DatabaseConnectionString)
  if err != nil {
    panic(err)
  }
  fmt.Println("Successfully Connected to mysql")
  db = d
  wg.Done()
}

func ConnectRedis(wg *sync.WaitGroup) {
  rdb = redis.NewClient(&redis.Options{
        Addr:     env.GetEnv().RedisHost,
        Password: env.GetEnv().RedisPassword, // no password set
        DB:       0,  // use default DB
    })

  fmt.Println("Successfully Connected to redis")
  wg.Done()
}

func Connect() {
  var wg sync.WaitGroup
	wg.Add(2)
  printTime()
  go ConnectMySql(&wg)
  printTime()
  go ConnectRedis(&wg)
  wg.Wait()
  printTime()
}

func GetRedis() *redis.Client {
  return rdb
}

func GetDB() *gorm.DB {
  return db
}
