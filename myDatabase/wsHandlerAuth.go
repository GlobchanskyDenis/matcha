package myDatabase

import (
	// "fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func (conn *ConnDB) wsReader(ws *websocket.Conn) {
	for {

		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Println("reader error happened -", err)
			return
		}

		log.Println("client said:", string(p))
		p = []byte("hi from server")

		if err := ws.WriteMessage(messageType, p); err != nil {
			log.Println("writeMessage error -", err)
			return
		}
	}
}

func (conn *ConnDB) WebSocketHandlerAuth(w http.ResponseWriter, r *http.Request) {
	var (
		login = r.URL.Query().Get("login")
		tempPasswd = r.URL.Query().Get("tempPasswd")
	)

	log.Printf("WEBSOCKET login=%s passwd=%s\n", login, tempPasswd)
	
	upgrader.CheckOrigin = func(r *http.Request) bool { return true	}
	ws, err := upgrader.Upgrade(w, r, nil)


	if err != nil {
		log.Println("wsEndPoint error happened -", err)
		return
	}
	tokenWS, err := conn.session.GetTokenWS(login)
	if err != nil {
		log.Printf("ERROR!!! %s", err)
		ws.Close()
		return
	}
	log.Println("tokenWS = ", tokenWS)
	if tokenWS != tempPasswd {
		log.Println("ERROR!!! ws-auth-token is wrong!!!!")
		ws.Close()
		return
	}
	log.Println("Client Successfully connected...")

	conn.wsReader(ws)
}
