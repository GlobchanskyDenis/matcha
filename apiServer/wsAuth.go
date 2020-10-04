package apiServer

import (
	. "MatchaServer/common"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (server *Server) wsWriteErrorMessage(r *http.Request, ws *websocket.Conn, messageBody string) error {
	message := `{"type":"error","uidSender":0,"body":"` + messageBody + `"}`
	err := ws.WriteMessage(1, []byte(message))
	if err != nil {
		server.Logger.LogError(r, "ws.WriteMessage returned error - "+err.Error())
		return nil
	}
	return nil
}

// INFINITE LOOP THAT HANDLES MESSAGES FROM CURRENT USER
func (server *Server) wsReader(r *http.Request, ws *websocket.Conn, uid int) {
	var decodedMessage map[string]interface{}

	for {
		_, RequestMessage, err := ws.ReadMessage()
		if err != nil {
			closeErr, ok := err.(*websocket.CloseError)
			if ok && (closeErr.Code == 1000 || closeErr.Code == 1001) {
				server.Logger.LogWarning(r, closeErr.Error())
			} else {
				server.Logger.LogError(r, "ws.ReadMessage returned error - "+err.Error())
			}
			return
		}
		err = json.Unmarshal(RequestMessage, &decodedMessage)
		if err != nil {
			server.Logger.LogError(r, "request json decode failed - "+err.Error()+`. Skip request`)
			err = server.wsWriteErrorMessage(r, ws, "request json decode failed")
			if err != nil {
				server.Logger.LogError(r, "wsWriteErrorMessage returned error - "+err.Error())
				return
			}
			continue
		}
		arg, isExists := decodedMessage["uidReceiver"]
		if !isExists {
			server.Logger.LogWarning(r, `"uidReceiver" not exist in received by ws message. Skip request`)
			err = server.wsWriteErrorMessage(r, ws, `"uidReceiver" not exist in received by ws message`)
			if err != nil {
				server.Logger.LogError(r, "wsWriteErrorMessage returned error - "+err.Error())
				return
			}
			continue
		}
		uidReceiverFloat, ok := arg.(float64)
		if !ok {
			server.Logger.LogWarning(r, `wrong type of "uidReceiver". Skip request`)
			err = server.wsWriteErrorMessage(r, ws, `wrong type of "uidReceiver"`)
			if err != nil {
				server.Logger.LogError(r, "wsWriteErrorMessage returned error - "+err.Error())
				return
			}
			continue
		}
		uidReceiver := int(uidReceiverFloat)
		server.Logger.Log(r, "user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+" ("+BLUE+r.Host+NO_COLOR+
			") wants to send message to user #"+BLUE+strconv.Itoa(uidReceiver)+NO_COLOR)
		arg, isExists = decodedMessage["body"]
		if !isExists {
			server.Logger.LogWarning(r, `"body" not exist in received by ws message. Skip request`)
			err = server.wsWriteErrorMessage(r, ws, `"body" not exist in received by ws message`)
			if err != nil {
				server.Logger.LogError(r, "wsWriteErrorMessage returned error - "+err.Error())
				return
			}
			continue
		}
		messageBody, ok := arg.(string)
		if !ok {
			server.Logger.LogWarning(r, `wrong type of "body". Skip request`)
			err = server.wsWriteErrorMessage(r, ws, `wrong type of "body"`)
			if err != nil {
				server.Logger.LogError(r, "wsWriteErrorMessage returned error - "+err.Error())
				return
			}
			continue
		}
		isExists, err = server.Db.IsUserExistsByUid(uidReceiver)
		if !isExists {
			server.Logger.LogWarning(r, `user #`+BLUE+strconv.Itoa(uidReceiver)+NO_COLOR+
				` not exists in database. Skip request`)
			err = server.wsWriteErrorMessage(r, ws, `user #`+strconv.Itoa(uidReceiver)+` not exists in database`)
			if err != nil {
				server.Logger.LogError(r, "wsWriteErrorMessage returned error - "+err.Error())
				return
			}
			continue
		}
		_, err = server.Db.SetNewMessage(uid, uidReceiver, messageBody)
		if err != nil {
			server.Logger.LogWarning(r, `SetNewMessage returned error - `+err.Error())
			return
		}
		err = server.Session.SendMessageToLoggedUser(uidReceiver, uid, messageBody)
		if err != nil {
			server.Logger.LogWarning(r, `SendMessageToLoggedUser returned error - `+err.Error())
			return
		}
		server.Logger.LogSuccess(r, "message from user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+
			" ("+BLUE+r.Host+NO_COLOR+") to user #"+BLUE+strconv.Itoa(uidReceiver)+NO_COLOR+" transmitted")
	}
}

// WEB SOCKET HANDLER FOR DOMAIN /ws/auth/
// GET PARAMS login AND ws-auth-token SHOULD BE IN REQUEST
func (server *Server) WebSocketAuth(w http.ResponseWriter, r *http.Request) {
	var wsAuthToken = r.URL.Query().Get("ws-auth-token")
	var message string
	var uidStr = r.URL.Query().Get("uid")
	var uid, err = strconv.Atoi(uidStr)
	if err != nil {
		server.Logger.LogError(r, "uid is invalid "+err.Error())
		// I should open and close ws connection for browser didn't write errors in server..logs
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		ws.Close()
		return
	}

	server.Logger.Log(r, "Request was recieved, uid="+BLUE+strconv.Itoa(uid)+NO_COLOR+
		" ws-auth-token="+BLUE+wsAuthToken+NO_COLOR)

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		server.Logger.LogError(r, "upgrader returned error - "+err.Error())
		return
	}
	if uid < 1 || wsAuthToken == "" {
		server.Logger.LogWarning(r, "wrong uid or ws-auth-token is empty")
		ws.Close()
		return
	}
	tokenWS, err := server.Session.GetTokenWS(uid)
	if err != nil {
		server.Logger.LogError(r, "GetTokenWS returned error - "+err.Error())
		ws.Close()
		return
	}
	if tokenWS != wsAuthToken {
		server.Logger.LogWarning(r, "ws-auth-token is wrong! Close Web Socket for user #"+
			BLUE+strconv.Itoa(uid)+NO_COLOR)
		ws.Close()
		return
	}

	server.Session.AddWSConnection(uid, ws)

	server.Logger.LogSuccess(r, "WebSocket was created")

	server.wsReader(r, ws, uid)

	userSessionWasClosed, err := server.Session.RemoveWSConnection(uid, ws)
	if err != nil {
		server.Logger.LogWarning(r, "RemoveWSConnection returned error: "+err.Error())
	} else {
		if userSessionWasClosed {
			message = "remove ws connection from session. " +
				"User session was closed"
		} else {
			message = "remove ws connection from session. " +
				"User session wasn't close - other devices still logged"
		}
		server.Logger.Log(r, message)
	}
	ws.Close()
}
