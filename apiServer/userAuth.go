package apiServer

import (
	. "MatchaServer/config"
	"MatchaServer/handlers"
	"MatchaServer/errDef"
	"encoding/json"
	"errors"
	"net/http"
)

func (server *Server) deviceHandler(w http.ResponseWriter, r *http.Request, uid int) error {
	var (
		devices     []Device
		device      Device
		knownDevice bool
		err         error
	)

	devices, err = server.Db.GetDevicesByUid(uid)
	if err != nil {
		return errors.New("GetDevicesByUid returned error " + err.Error())
	}
	for _, device = range devices {
		if device.Device == r.UserAgent() {
			knownDevice = true
		}
	}
	if !knownDevice {
		err = server.Db.SetNewDevice(uid, r.UserAgent())
		if err != nil {
			return errors.New("SetNewDevice returned error " + err.Error())
		}
		err = server.session.SendNotifToLoggedUser(uid, 0, `device from `+r.Host+" found:"+r.UserAgent())
		if err != nil {
			return errors.New("SendNotifToLoggedUser returned error " + err.Error())
		}
	}
	return nil
}

// USER AUTHORISATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) userAuth(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, pass, token, tokenWS, response string
		user                                          User
		err                                           error
		request                                       map[string]interface{}
		isExist                                       bool
	)

	defer func() {
		if err := recover(); err != nil {
			println("PANIC!!!!! " + err.(error).Error())
		}
	}()

	

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/auth/", "request decode error")
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "decode error" + `"}`))
		return
	}

	arg, isExist := request["mail"]
	if !isExist {
		consoleLogWarning(r, "/user/auth/", "mail not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "mail not exist" + `"}`))
		return
	}
	mail = arg.(string)

	arg, isExist = request["pass"]
	if !isExist {
		consoleLogWarning(r, "/user/auth/", "password not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "password not exist" + `"}`))
		return
	}
	pass = arg.(string)

	message = "request was recieved, mail: " + BLUE + mail + NO_COLOR + " password: hidden "
	consoleLog(r, "/user/auth/", message)

	// Simple validation
	if mail == "" || pass == "" {
		consoleLogWarning(r, "/user/auth/", "mail or password is empty")
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		w.Write([]byte(`{"error":"` + "mail or password is empty" + `"}`))
		return
	}

	user, err = server.Db.GetUserForAuth(mail, handlers.PassHash(pass))


	if errDef.IsRecordNotFoundError(err) {
		consoleLogWarning(r, "/user/auth/", "GetUserForAuth - record not found")
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	} else if err != nil {
		consoleLogError(r, "/user/auth/", "GetUserForAuth returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "database request failed" + `"}`))
		return
	}

	// if user.Uid == 0 {
	// 	// it means that no such users in database
	// 	consoleLogWarning(r, "/user/auth/", "wrong mail or password")
	// 	w.WriteHeader(http.StatusUnprocessableEntity) // 422
	// 	w.Write([]byte(`{"error":"` + "wrong mail or password" + `"}`))
	// 	return
	// }

	if user.Status == "not confirmed" {
		consoleLogWarning(r, "/user/auth/", "user "+BLUE+user.Mail+NO_COLOR+" should confirm its email")
		w.WriteHeader(http.StatusAccepted) // 202
		w.Write([]byte(`{"error":"` + "confirm email first" + `"}`))
		return
	}

	// Check if this device is unknown yet - then make notification that new device if found
	err = server.deviceHandler(w, r, user.Uid)
	if err != nil {
		consoleLogError(r, "/user/auth/", err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "Database or websocket error" + `"}`))
		return
	}

	token, err = server.session.AddUserToSession(user.Uid)
	if err != nil {
		consoleLogError(r, "/user/auth/", "AddUserToSession returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "Web socket error" + `"}`))
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		consoleLogWarning(r, "/user/auth/", "Marshal returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "cannot convert to json" + `"}`))
		return
	}

	tokenWS, err = server.session.CreateTokenWS(user.Uid) //handlers.TokenWebSocketAuth(mail)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		consoleLogError(r, "/user/auth/", "cannot create web socket token - "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "cannot create web socket token" + `"}`))
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
