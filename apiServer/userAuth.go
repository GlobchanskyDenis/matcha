package apiServer

import (
	"MatchaServer/handlers"
	"MatchaServer/errDef"
	"encoding/json"
	"net/http"
)

func (server *Server) deviceHandler(w http.ResponseWriter, r *http.Request, uid int) error {
	var (
		knownDevice bool
		err         error
	)

	devices, err := server.Db.GetDevicesByUid(uid)
	if err != nil {
		consoleLogError(r, "/user/auth/", "GetDevicesByUid returned error - " + err.Error())
		return errDef.DatabaseError
	}
	for _, device := range devices {
		if device.Device == r.UserAgent() {
			knownDevice = true
		}
	}
	if !knownDevice {
		err = server.Db.SetNewDevice(uid, r.UserAgent())
		if err != nil {
			consoleLogError(r, "/user/auth/", "SetNewDevice returned error - " + err.Error())
			return errDef.DatabaseError
		}
		err = server.session.SendNotifToLoggedUser(uid, 0, `device from `+r.Host+" found:"+r.UserAgent())
		if err != nil {
			consoleLogError(r, "/user/auth/", "SendNotifToLoggedUser returned error - " + err.Error())
			return errDef.WebSocketError
		}
	}
	return nil
}

// USER AUTHORISATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) userAuth(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, pass, token, tokenWS, response string
		err                                           error
		request                                       map[string]interface{}
		isExist, ok                                   bool
	)

	defer func() {
		if err := recover(); err != nil {
			println(RED_BG + "PANIC!!!!! " + err.(error).Error() + NO_COLOR)
		}
	}()

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/auth/", "request decode error - " + err.Error())
		server.error(w, errDef.InvalidRequestBody)
		return
	}

	arg, isExist := request["mail"]
	if !isExist {
		consoleLogWarning(r, "/user/auth/", "mail not exist")
		server.error(w, errDef.NoArgument.WithArguments("Поле mail отсутствует", "mail field expected"))
		return
	}

	mail, ok = arg.(string)
	if !ok {
		consoleLogWarning(r, "/user/auth/", "mail has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле mail имеет неверный тип", "mail field has wrong type"))
		return
	}

	arg, isExist = request["pass"]
	if !isExist {
		consoleLogWarning(r, "/user/auth/", "password not exist")
		server.error(w, errDef.NoArgument.WithArguments("Поле pass отсутствует", "pass field expected"))
		return
	}

	pass, ok = arg.(string)
	if !ok {
		consoleLogWarning(r, "/user/auth/", "password has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле pass имеет неверный тип", "pass field has wrong type"))
		return
	}

	message = "request was recieved, mail: " + BLUE + mail + NO_COLOR + " password: hidden "
	consoleLog(r, "/user/auth/", message)

	// Simple validation
	if mail == "" || pass == "" {
		consoleLogWarning(r, "/user/auth/", "mail or password is empty")
		server.error(w, errDef.AuthFail)
		return
	}

	user, err := server.Db.GetUserForAuth(mail, handlers.PassHash(pass))
	if errDef.AuthFail.IsOverlapWithError(err) {
		consoleLogWarning(r, "/user/auth/", "Authorization for user "+BLUE+mail+NO_COLOR+" failed")
		server.error(w, errDef.AuthFail)
		return
	} else if err != nil {
		consoleLogError(r, "/user/auth/", "GetUserForAuth returned error "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	if user.Status == "not confirmed" {
		consoleLogWarning(r, "/user/auth/", "user "+BLUE+user.Mail+NO_COLOR+" should confirm its email")
		server.error(w, errDef.NotConfirmedMail)
		return
	}

	// Check if this device is unknown yet - then make notification that new device if found
	err = server.deviceHandler(w, r, user.Uid)
	if err != nil {
		server.error(w, err.(errDef.ApiError))
		return
	}

	token, err = server.session.AddUserToSession(user.Uid)
	if err != nil {
		consoleLogError(r, "/user/auth/", "Cannot add user to session - "+err.Error())
		server.error(w, errDef.UnknownInternalError)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		consoleLogError(r, "/user/auth/", "Marshal returned error "+err.Error())
		server.error(w, errDef.MarshalError)
		return
	}

	tokenWS, err = server.session.CreateTokenWS(user.Uid) //handlers.TokenWebSocketAuth(mail)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		consoleLogError(r, "/user/auth/", "cannot create web socket token - "+err.Error())
		server.error(w, errDef.WebSocketError)
		return
	}

	// This is my valid case. Response status will be set automaticly to 200.
	w.WriteHeader(http.StatusOK) // 200
	response = `{"x-auth-token":"` + token + `","ws-auth-token":"` + tokenWS + `",` + string(jsonUser[1:])
	w.Write([]byte(response))
	consoleLogSuccess(r, "/user/auth/", "User "+BLUE+mail+NO_COLOR+" was authenticated successfully")
}

// HTTP HANDLER FOR DOMAIN /auth/ . IT HANDLES:
// AUTHENTICATE USER BY POST METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) HandlerUserAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "POST" {
		server.userAuth(w, r)
	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)
		consoleLog(r, "/user/auth/", "client wants to know what methods are allowed")
	} else {
		// ALL OTHERS METHODS
		consoleLogWarning(r, "/user/auth/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
