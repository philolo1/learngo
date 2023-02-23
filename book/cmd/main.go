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
  url := "localhost:" + env.GetEnv().Port
  fmt.Printf("Serving at %s\n", url)
  log.Fatal(http.ListenAndServe(url, r))
}
