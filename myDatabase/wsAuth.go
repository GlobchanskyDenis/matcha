package myDatabase

import (
	. "MatchaServer/config"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// func (conn *ConnDB) wsWriteMessage(r *http.Request, ws *websocket.Conn, message string) error {
// 	err := ws.WriteMessage(1, []byte(message))
// 	if err != nil {
// 		consoleLogError(r, "/ws/", "Error: ws.WriteMessage returned error - " + fmt.Sprintf("%s", err))
// 		return nil
// 	}
// 	return nil
// }

func (conn *ConnDB) wsReader(r *http.Request, ws *websocket.Conn) {
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			consoleLogError(r, "/ws/", "ws.ReadMessage returned error - "+err.Error())
			return
		}
		consoleLog(r, "/ws/", "client said: "+string(message))
		message = []byte("hi from server")
		err = ws.WriteMessage(messageType, message)
		if err != nil {
			consoleLogError(r, "/ws/", "ws.WriteMessage returned error - "+err.Error())
			return
		}
	}
}

// WEB SOCKET HANDLER FOR DOMAIN /ws/
// GET PARAMS login AND ws-auth-token SHOULD BE IN REQUEST
func (conn *ConnDB) WebSocketHandlerAuth(w http.ResponseWriter, r *http.Request) {
	var token = r.URL.Query().Get("ws-auth-token")
	var uidStr = r.URL.Query().Get("uid")
	uid, err := strconv.Atoi(uidStr)
	if err != nil {
		consoleLogError(r, "/ws/", "Atoi returned error - "+err.Error())
		return
	}

	consoleLog(r, "/ws/", "Request was recieved, uid="+BLUE+uidStr+NO_COLOR+" ws-auth-token="+BLUE+token+NO_COLOR)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		consoleLogError(r, "/ws/", "upgrader returned error - "+err.Error())
		return
	}
	if uid < 1 || token == "" {
		consoleLogWarning(r, "/ws/", "wrong uid or ws-auth-token is empty")
		ws.Close()
		return
	}
	tokenWS, err := conn.session.GetTokenWS(uid)
	if err != nil {
		consoleLogError(r, "/ws/", "GetTokenWS returned error - "+err.Error())
		ws.Close()
		return
	}
	if tokenWS != token {
		consoleLogWarning(r, "/ws/", "ws-auth-token is wrong! Close Web Socket for user #"+BLUE+uidStr+NO_COLOR)
		ws.Close()
		return
	}
	consoleLogSuccess(r, "/ws/", "WebSocket was created")
	conn.wsReader(r, ws)
	ws.Close()
}
