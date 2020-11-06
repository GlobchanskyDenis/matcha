package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"MatchaServer/handlers"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type userMessage struct {
	Type string `json:"type"`
	UidReceiver int `json:"uidReceiver"`
	Body string `json:"body"`
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

func requestMessageValidator(message userMessage) error {
	switch message.Type {
	case "logout":
		return nil
	case "message":
		if message.UidReceiver <= 0 || message.Body == "" {
			return errors.NewArg("Невалидное сообщение", "Invalid message")
		}
		return nil
	default:
		return errors.NewArg("Незнакомый тип сообщения", "Unknown message type")
	}
	return errors.NewArg("Незнакомый тип сообщения", "Unknown message type")
}

func (server *Server) sendMessage(r *http.Request, ws *websocket.Conn, uid int, message userMessage) {
	isExists, err := server.Db.IsUserExistsByUid(message.UidReceiver)
	if !isExists {
		server.Logger.LogWarning(r, `user #`+BLUE+strconv.Itoa(message.UidReceiver)+NO_COLOR+
			` not exists in database. Skip request`)
		err = server.wsWriteErrorMessage(r, ws, `user #`+strconv.Itoa(message.UidReceiver)+` not exists in database`)
		if err != nil {
			server.Logger.LogError(r, "wsWriteErrorMessage returned error - "+err.Error())
		}
		return
	}
	_, err = server.Db.SetNewMessage(uid, message.UidReceiver, message.Body)
	if err != nil {
		server.Logger.LogWarning(r, `SetNewMessage returned error - `+err.Error())
		err = server.wsWriteErrorMessage(r, ws, `user #`+strconv.Itoa(message.UidReceiver)+` not exists in database`)
		if err != nil {
			server.Logger.LogError(r, "wsWriteErrorMessage returned error - "+err.Error())
		}
		return
	}
	err = server.Session.SendMessageToLoggedUser(message.UidReceiver, uid, message.Body)
	if err != nil {
		server.Logger.LogWarning(r, `SendMessageToLoggedUser returned error - `+err.Error())
		err = server.wsWriteErrorMessage(r, ws, `user #`+strconv.Itoa(message.UidReceiver)+` not exists in database`)
		if err != nil {
			server.Logger.LogError(r, "wsWriteErrorMessage returned error - "+err.Error())
		}
		return
	}
	server.Logger.LogSuccess(r, "message from user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+
			" ("+BLUE+r.Host+NO_COLOR+") to user #"+BLUE+strconv.Itoa(message.UidReceiver)+NO_COLOR+" transmitted")
}

// INFINITE LOOP THAT HANDLES MESSAGES FROM CURRENT USER
func (server *Server) wsReader(r *http.Request, ws *websocket.Conn, uid int) (logout bool) {
	var message userMessage

	for {
		//  Получим сообщение от пользователя
		_, jsonMessage, err := ws.ReadMessage()
		if err != nil {
			closeErr, ok := err.(*websocket.CloseError)
			if ok && (closeErr.Code == 1000 || closeErr.Code == 1001) {
				server.Logger.LogWarning(r, closeErr.Error())
			} else {
				server.Logger.LogError(r, "ws.ReadMessage returned error - "+err.Error())
			}
			return false
		}

		// Распакуем сообщение из формата json
		err = json.Unmarshal(jsonMessage, &message)
		if err != nil {
			server.Logger.LogError(r, "request json decode failed - "+err.Error()+`. Skip request`)
			err = server.wsWriteErrorMessage(r, ws, "request json decode failed")
			if err != nil {
				server.Logger.LogError(r, "wsWriteErrorMessage returned error - "+err.Error())
				return false
			}
			continue
		}

		//  Провалидируем сообщение
		err = requestMessageValidator(message)
		if err != nil {
			server.Logger.LogError(r, `user message validation error. `+BLUE+err.Error()+NO_COLOR+` Skip request`)
			err = server.wsWriteErrorMessage(r, ws, err.Error())
			if err != nil {
				server.Logger.LogError(r, "wsWriteErrorMessage returned error - "+err.Error())
				return true
			}
			continue
		}

		//  В зависимости от типа сообщения - обработаем его
		switch message.Type {
		case "message":
			server.sendMessage(r, ws, uid, message)
		case "logout":
			return true
		}
	}
	return false
}

// WEB SOCKET HANDLER FOR DOMAIN /ws/auth/
// GET PARAMS login AND ws-auth-token SHOULD BE IN REQUEST
func (server *Server) WebSocketAuth(w http.ResponseWriter, r *http.Request) {
	var xAuthToken = r.URL.Query().Get("x-auth-token")
	var logMessage string
	var uid int
	var isLogged bool

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		server.Logger.LogError(r, "upgrader returned error - "+err.Error())
		return
	}

	uid, err = handlers.TokenUidDecode(xAuthToken)
	if err != nil {
		server.Logger.LogWarning(r, "TokenUidDecode returned error - "+err.Error())
		ws.Close()
		return
	}
	isLogged = server.Session.IsUserLoggedByUid(uid)
	if !isLogged {
		server.Logger.LogWarning(r, "User #"+strconv.Itoa(uid)+" is not logged")
		ws.Close()
		return
	}

	userAgent := "default"
	if len(r.Header["User-Agent"]) >= 1 {
		userAgent = r.Header["User-Agent"][0]
	}

	server.Logger.LogSuccess(r, "web socket was created. Uid="+BLUE+strconv.Itoa(uid)+NO_COLOR)

	server.Session.AddWSConnection(uid, ws, userAgent)

	isLogout := server.wsReader(r, ws, uid)

	userSessionWasClosed, err := server.Session.RemoveWSConnection(uid, userAgent, isLogout)
	if err != nil {
		server.Logger.LogWarning(r, "RemoveWSConnection returned error: "+err.Error())
	} else {
		if userSessionWasClosed {
			logMessage = "remove ws connection from session. " +
				"User session was closed"
		} else {
			logMessage = "remove ws connection from session. " +
				"User session wasn't close - other devices still logged"
		}
		server.Logger.Log(r, logMessage)
	}
	ws.Close()
}
