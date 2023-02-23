package main
import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/philolo1/book/pkg/routes"
  "github.com/joho/godotenv"
  "github.com/philolo1/book/pkg/models"
  "os"
)

func main() {
  if err := godotenv.Load(); err != nil {
    panic("Error loading .env file")
  }
  fmt.Printf("Values: %v", os.Getenv("DATABASE_CONNECTION_STRING"))
  models.InitializeDB()
  r := mux.NewRouter()
  routes.RegisterBookStoreRoutes(r)
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe("localhost:9010", r))
}
