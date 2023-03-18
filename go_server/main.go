package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
	"encoding/json"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
  if r.URL.Path != "/hello" {
    http.Error(w, "404 not found", http.StatusNotFound)
    return
  }

  if r.Method != "GET" {
    http.Error(w, "method not foun", http.StatusNotFound)
    return
  }
	resp := make(map[string]string)
	resp["message"] = "hello"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)

}

func formHandler(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/form" {
    http.Error(w, "404 not found", http.StatusNotFound)
  }

  if r.Method != "POST" {
    http.Error(w, "method not found", http.StatusNotFound)
  }

  name := r.FormValue("name")
  address := r.FormValue("address")

  fmt.Fprintf(w, "name: %s, address %s", name, address)

}

func main() {
  fileServer := http.FileServer(http.Dir("./static"))
  http.Handle("/", fileServer)
  http.HandleFunc("/hello", helloHandler)
  http.HandleFunc("/form", formHandler)

  port := os.Getenv("PORT")
	if port == "" {
    log.Panicf("PORT Variable not set")
	}

	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
