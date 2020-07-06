package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	// "strconv"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

// var i int

func reader(conn *websocket.Conn) {
	for {

		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("reader error happened -", err)
			return
		}

		log.Println("reader:", string(p))

		// p = []uint8("hi from server " + strconv.Itoa(i))
		// p = []byte("hi from server")
		// i++

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println("writeMessage error -", err)
			return
		}

	}
}

func wsEndPoint(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "WebsocketHere")
	// log.Printf("someone call web socket")

	upgrader.CheckOrigin = func(r *http.Request) bool { return true	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("wsEndPoint error happened -", err)
	}
	log.Println("Client Successfully connected...")

	reader(ws)
}

func main() {
	fmt.Println("start server at :3000")
	http.HandleFunc("/ws/", wsEndPoint)
	log.Fatal(http.ListenAndServe(":3000", nil))
}