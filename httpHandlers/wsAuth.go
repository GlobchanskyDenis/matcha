package httpHandlers

import (
	. "MatchaServer/config"
	"MatchaServer/handlers"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	// "fmt"
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

func (conn *ConnAll) wsReader(r *http.Request, ws *websocket.Conn) {
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			consoleLogError(r, "/ws/auth/", "ws.ReadMessage returned error - "+err.Error())
			return
		}
		consoleLog(r, "/ws/auth/", "client said: "+string(message))
		message = []byte("hi from server")
		err = ws.WriteMessage(messageType, message)
		if err != nil {
			consoleLogError(r, "/ws/auth/", "ws.WriteMessage returned error - "+err.Error())
			return
		}
	}
}

// WEB SOCKET HANDLER FOR DOMAIN /ws/
// GET PARAMS login AND ws-auth-token SHOULD BE IN REQUEST
func (conn *ConnAll) WebSocketHandlerAuth(w http.ResponseWriter, r *http.Request) {
	var wsAuthToken = r.URL.Query().Get("ws-auth-token")
	var uidToken = r.URL.Query().Get("x-auth-token")
	var message string
	uid, err := handlers.TokenUidDecode(uidToken)
	if err != nil {
		consoleLogError(r, "/ws/auth/", "TokenUidDecode returned error "+err.Error())
		// I should open and close ws connection for browser didn't write errors in console.logs
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		ws.Close()
		return
	}

	consoleLog(r, "/ws/auth/", "Request was recieved, uid="+BLUE+strconv.Itoa(uid)+NO_COLOR+" ws-auth-token="+BLUE+wsAuthToken+NO_COLOR)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		consoleLogError(r, "/ws/auth/", "upgrader returned error - "+err.Error())
		return
	}
	if uid < 1 || wsAuthToken == "" {
		consoleLogWarning(r, "/ws/auth/", "wrong uid or ws-auth-token is empty")
		ws.Close()
		return
	}
	tokenWS, err := conn.session.GetTokenWS(uid)
	if err != nil {
		consoleLogError(r, "/ws/auth/", "GetTokenWS returned error - "+err.Error())
		ws.Close()
		return
	}
	if tokenWS != wsAuthToken {
		consoleLogWarning(r, "/ws/auth/", "ws-auth-token is wrong! Close Web Socket for user #"+BLUE+strconv.Itoa(uid)+NO_COLOR)
		ws.Close()
		return
	}

	err = conn.session.AddWSConnection(uid, ws, r.Host + ": " + r.UserAgent())
	if err != nil {
		consoleLogWarning(r, "/ws/auth/", "AddWSConnection returned error: " + err.Error())
	}

	consoleLogSuccess(r, "/ws/auth/", "WebSocket was created")

	conn.wsReader(r, ws)

	userSessionWasClosed, err := conn.session.RemoveWSConnection(uid, ws)
	if err != nil {
		consoleLogWarning(r, "/ws/auth/", "RemoveWSConnection returned error: " + err.Error())
	} else {
		if userSessionWasClosed {
			message = "ws connection is going for close, remove it from session. " + "User session was closed"
		} else {
			message = "ws connection is going for close, remove it from session. " + "User session wasnt close - other devices still logged"
		}
		consoleLog(r, "/ws/auth/", message)
	}
	ws.Close()
}
