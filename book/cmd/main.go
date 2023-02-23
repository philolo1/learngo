package main
import (
  "fmt"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/philolo1/book/pkg/routes"
  "github.com/philolo1/book/pkg/models"
  "github.com/philolo1/book/pkg/env"
)

func main() {
  fmt.Printf("Env: %s\n", env.GetEnv())
  models.InitializeDB()
  r := mux.NewRouter()
  routes.RegisterBookStoreRoutes(r)
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe("localhost:9010", r))
}
