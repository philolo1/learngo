package controllers

import (
  "encoding/json"
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "strconv"
  "github.com/philolo1/book/pkg/models"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
  newBooks := models.GetAllBooks()
  res, _ := json.Marshal(newBooks)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  bookId := vars["bookId"]
  id, err:= strconv.ParseInt(bookId, 0, 0)
  if err != nil {
    fmt.Println("error while parsing")
  }
  bookDetails := models.GetBookById(id)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  res, _ := json.Marshal(bookDetails)
  w.Write(res)
}

type Status struct {
  Message string `json:"message"`
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
  // TODO add variables later
  // vars := mux.Vars(r)
  // bookId := vars["bookId"]
  newBook := models.Book{
    Name: "Test",
    Author: "author",
    Publication: "Something",
  }
  models.CreateBook(&newBook)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)


  status := Status{Message: "OK"}
  jsonData, _ := json.Marshal(status)
  w.Write(jsonData)
}

