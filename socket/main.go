
package main

import (
	"fmt"
  "time"
	"log"
  "strings"
	"net/http"
	"github.com/gorilla/websocket"
  "sync"
)

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var mutex sync.Mutex

var socketConnections []*websocket.Conn = make([]*websocket.Conn, 0, 10)

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
    messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// print out that message for clarity
		log.Println(string(p))

    if err := conn.WriteMessage(messageType, p); err != nil {
      log.Println(err)
    	return
    }
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func heartbeat() {
  ticker := time.NewTicker(1 * time.Second)
  defer ticker.Stop()

  for {
    select {
    case <-ticker.C:
      mutex.Lock()
      for _, value := range socketConnections {
        value.WriteMessage(websocket.TextMessage, []byte("hearbeat"))
      }
      // do something every second
      fmt.Printf("Tick! %d \n", len(socketConnections))
      mutex.Unlock()
    }
  }
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
  defer ws.Close()

  mutex.Lock()
  socketConnections = append(socketConnections, ws)
  mutex.Unlock()
  defer func() {
    mutex.Lock()
    defer mutex.Unlock()
    for index, value := range socketConnections {
      if (value == ws) {
        fmt.Printf("Remove item")
        socketConnections = append(socketConnections[:index],socketConnections[index +1:]...)
        return
      }
    }
    fmt.Println("End fun")
  }()

  for {
    messageType, p, err := ws.ReadMessage()
    if err != nil {
      log.Println("Connection Closed")
			break
		}

    word := strings.Trim(string(p), "\n")

    fmt.Printf("compare: '%s':'close'\n", word)

    if word == "close" {
      ws.WriteMessage(messageType, []byte(p))
      log.Println("break")
      break
    }

    log.Println(string(p))
  }

	// log.Println("Client Connected")
	// err = ws.WriteMessage(1, []byte("Hi Client!"))
	// if err != nil {
	// 	log.Println(err)
	// }
	// // listen indefinitely for new messages coming
	// // through on our WebSocket connection
	// reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World, connect with ws://127.0.0.1:8080/ws")
  go heartbeat()
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

