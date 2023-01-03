package main

import (
  "fmt"
  "sync"
  "log"
  "net/http"
)

var urls = []string{
  "https://google.com",
  "https://twitter.com",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Call me", r.URL)
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }
  var wg sync.WaitGroup
  for _, url:= range urls {
    wg.Add(1)
    go func(url string) {
      resp, err := http.Get(url)
      if err != nil {
        fmt.Fprintf(w, "%+v\n", resp)
      }
      fmt.Fprintf(w, "%+v\n", resp.Status)
      wg.Done()
    }(url)
  }
  wg.Wait()
}


func addNum(wg *sync.WaitGroup,c *chan int) {
  for l:=0; l < 100; l++ {
    *c <- l
  }
   wg.Done()
}

func main() {
    fmt.Println("Hello")


    var wg sync.WaitGroup
    wg.Add(2)
    c := make(chan int)  // Create a new channel
    go addNum(&wg, &c)
    go addNum(&wg, &c)

    for l:=0; l < 200; l++ {
      fmt.Println(<-c)  // Output: 3
    }

    http.HandleFunc("/", fetchStatus)

    fmt.Println("Serving :8000")
    log.Fatal(http.ListenAndServe(":8000", nil))
}
