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

  if uint(id) != bookDetails.ID {
    // Return a 404 Not Found status code if the item is not found
    http.NotFound(w, r)
    return
  }



  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  res, _ := json.Marshal(bookDetails)
  w.Write(res)
}

type Status struct {
  Message string `json:"message"`
}


func DeleteBooks(w http.ResponseWriter, r *http.Request) {
  models.DeleteBooks()

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)

  jsonData, _ := json.Marshal(Status{Message: "ok"})
  w.Write(jsonData)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  bookId := vars["bookId"]
  id, err:= strconv.ParseInt(bookId, 0, 0)
  if err != nil {
    fmt.Println("error while parsing")
  }

  res := models.DeleteBookById(id)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  jsonData, _ := json.Marshal(Status{Message: fmt.Sprintf("Rows effected %d", res.RowsAffected)})
  w.Write(jsonData)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
  // TODO add variables later
  // vars := mux.Vars(r)
  var book models.Book
  json.NewDecoder(r.Body).Decode(&book)
  fmt.Printf("vars %v\n", book)
  newBook := book
  models.CreateBook(&newBook)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)

  jsonData, _ := json.Marshal(newBook)
  w.Write(jsonData)
}

