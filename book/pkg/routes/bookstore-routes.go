package routes

import (
  "github.com/gorilla/mux"
  "github.com/philolo1/book/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
  router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
  router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
  router.HandleFunc("/books", controllers.DeleteBooks).Methods("DELETE")
  router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
  router.HandleFunc("/book/{bookId}", controllers.DeleteBookById).Methods("DELETE")
  // router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
}
