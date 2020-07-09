package myDatabase

import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

// consoleLogError(r, "/user/", "Error: request decode error - " + fmt.Sprintf("%s", err))

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func (conn *ConnDB) wsReader(r *http.Request, ws *websocket.Conn) {
	for {

		messageType, message, err := ws.ReadMessage()
		if err != nil {
			consoleLogError(r, "/ws/", "Error: ws.ReadMessage returned error - " + fmt.Sprintf("%s", err))
			return
		}
		consoleLog(r, "/ws/", "client said: " + string(message))
		message = []byte("hi from server")
		err = ws.WriteMessage(messageType, message)
		if err != nil {
			consoleLogError(r, "/ws/", "Error: ws.WriteMessage returned error - " + fmt.Sprintf("%s", err))
			return
		}
	}
}

// WEB SOCKET HANDLER FOR DOMAIN /ws/
// GET PARAMS login AND ws-auth-token SHOULD BE IN REQUEST
func (conn *ConnDB) WebSocketHandlerAuth(w http.ResponseWriter, r *http.Request) {
	var login = r.URL.Query().Get("login")
	var token = r.URL.Query().Get("ws-auth-token")

	consoleLog(r, "/ws/", "Request was recieved, login=\033[34m" + login + "\033[32m ws-auth-token=\033[34m" + token)
	
	upgrader.CheckOrigin = func(r *http.Request) bool { return true	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		consoleLogError(r, "/ws/", "Error: upgrader returned error - " + fmt.Sprintf("%s", err))
		return
	}
	if login == "" || token == "" {
		consoleLogWarning(r, "/ws/", "Warning: login or ws-auth-token is empty")
		ws.Close()
		return
	}
	tokenWS, err := conn.session.GetTokenWS(login)
	if err != nil {
		consoleLogError(r, "/ws/", "Error: GetTokenWS returned error - " + fmt.Sprintf("%s", err))
		ws.Close()
		return
	}
	if tokenWS != token {
		consoleLogWarning(r, "/ws/", "Warning: ws-auth-token is wrong! Close Web Socket for user \033[34m" + login)
		ws.Close()
		return
	}
	consoleLog(r, "/ws/", "WebSocket succesfully created")
	conn.wsReader(r, ws)
}
