package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"MatchaServer/handlers"
	"context"
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
		server.Logger.LogError(r, "GetDevicesByUid returned error - "+err.Error())
		return errors.DatabaseError.WithArguments(err)
	}
	for _, device := range devices {
		if device.Device == r.UserAgent() {
			knownDevice = true
		}
	}
	if !knownDevice {
		err = server.Db.SetNewDevice(uid, r.UserAgent())
		if err != nil {
			server.Logger.LogError(r, "SetNewDevice returned error - "+err.Error())
			return errors.DatabaseError.WithArguments(err)
		}
		err = server.Session.SendNotifToLoggedUser(uid, 0, `device from `+r.Host+" found:"+r.UserAgent())
		if err != nil {
			server.Logger.LogError(r, "SendNotifToLoggedUser returned error - "+err.Error())
			return errors.WebSocketError.WithArguments(err)
		}
	}
	return nil
}

// HTTP HANDLER FOR DOMAIN /user/auth/ . IT HANDLES:
// AUTHENTICATE USER BY POST METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) UserAuth(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, pass, token, tokenWS, response string
		err                                           error
		requestParams                                 map[string]interface{}
		isExist, ok                                   bool
		ctx                                           context.Context
	)

	ctx = r.Context()
	requestParams = ctx.Value("requestParams").(map[string]interface{})

	arg, isExist := requestParams["mail"]
	if !isExist {
		server.Logger.LogWarning(r, "mail not exist")
		server.error(w, errors.NoArgument.WithArguments("Поле mail отсутствует", "mail field expected"))
		return
	}

	mail, ok = arg.(string)
	if !ok {
		server.Logger.LogWarning(r, "mail has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Поле mail имеет неверный тип", "mail field has wrong type"))
		return
	}

	arg, isExist = requestParams["pass"]
	if !isExist {
		server.Logger.LogWarning(r, "password not exist")
		server.error(w, errors.NoArgument.WithArguments("Поле pass отсутствует", "pass field expected"))
		return
	}

	pass, ok = arg.(string)
	if !ok {
		server.Logger.LogWarning(r, "password has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Поле pass имеет неверный тип", "pass field has wrong type"))
		return
	}

	message = "request was recieved, mail: " + BLUE + mail + NO_COLOR + " password: hidden "
	server.Logger.Log(r, message)

	// Simple validation
	if mail == "" || pass == "" {
		server.Logger.LogWarning(r, "mail or password is empty")
		server.error(w, errors.AuthFail)
		return
	}

	user, err := server.Db.GetUserForAuth(mail, handlers.PassHash(pass))
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "Authorization for user "+BLUE+mail+NO_COLOR+" failed")
		server.error(w, errors.AuthFail)
		return
	} else if err != nil {
		server.Logger.LogError(r, "GetUserForAuth returned error "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	if user.Status == "not confirmed" {
		server.Logger.LogWarning(r, "user "+BLUE+user.Mail+NO_COLOR+" should confirm its email")
		server.error(w, errors.NotConfirmedMail)
		return
	}

	// Check if this device is unknown yet - then make notification that new device if found
	err = server.deviceHandler(w, r, user.Uid)
	if err != nil {
		server.error(w, err.(errors.ApiError))
		return
	}

	token, err = server.Session.AddUserToSession(user.Uid)
	if err != nil {
		server.Logger.LogError(r, "Cannot add user to session - "+err.Error())
		server.error(w, errors.UnknownInternalError.WithArguments(err))
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		server.Logger.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errors.MarshalError)
		return
	}

	tokenWS, err = server.Session.CreateTokenWS(user.Uid) //handlers.TokenWebSocketAuth(mail)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		server.Logger.LogError(r, "cannot create web socket token - "+err.Error())
		server.error(w, errors.WebSocketError.WithArguments(err))
		return
	}

	// This is my valid case. Response status will be set automaticly to 200.
	w.WriteHeader(http.StatusOK) // 200
	response = `{"x-auth-token":"` + token + `","ws-auth-token":"` + tokenWS + `",` + string(jsonUser[1:])
	w.Write([]byte(response))
	server.Logger.LogSuccess(r, "User "+BLUE+mail+NO_COLOR+" was authenticated successfully")
}
