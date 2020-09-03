package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errDef"
	"MatchaServer/handlers"
	"encoding/json"
	"net/http"
	"strconv"
)

// USER DATA RETURN BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) userGet(w http.ResponseWriter, r *http.Request) {
	var (
		message, token string
		user							  User
		uid								  int
		err                                           error
		request                                       map[string]interface{}
		isExist, ok                                   bool
	)

	message = "request for own user data"
	server.Log(r, message)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		server.LogError(r, "request decode error - "+err.Error())
		server.error(w, errDef.InvalidRequestBody)
		return
	}

	arg, isExist := request["x-auth-token"]
	if !isExist {
		server.LogWarning(r, "x-auth-token not exists")
		server.error(w, errDef.NoArgument.WithArguments("Поле x-auth-token отсутствует", "x-auth-token field expected"))
		return
	}

	token, ok = arg.(string)
	if !ok {
		server.LogWarning(r, "token have wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле x-auth-token имеет неверный тип", "x-auth-token field has wrong type"))
		return
	}

	if token == "" {
		server.LogWarning(r, "token is empty")
		server.error(w, errDef.UserNotLogged)
		return
	}

	uid, err = handlers.TokenUidDecode(token)
	if err != nil {
		server.LogWarning(r, "TokenUidDecode returned error - "+err.Error())
		server.error(w, errDef.UserNotLogged)
		return
	}

	if !server.session.IsUserLoggedByUid(uid) {
		server.LogWarning(r, "user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+" is not logged")
		server.error(w, errDef.UserNotLogged)
		return
	}

	user, err = server.Db.GetUserByUid(uid)
	if errDef.RecordNotFound.IsOverlapWithError(err) {
		server.LogWarning(r, "GetUserByUid - record not found")
		server.error(w, errDef.UserNotExist)
		return
	} else if err != nil {
		server.LogError(r, "GetUser returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	if user.Status == "not confirmed" {
		server.LogWarning(r, "user "+BLUE+user.Mail+NO_COLOR+" should confirm its email")
		server.error(w, errDef.NotConfirmedMail)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		server.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errDef.MarshalError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonUser)
	server.LogSuccess(r, "User "+BLUE+mail+NO_COLOR+" was authenticated successfully")
}

// HTTP HANDLER FOR DOMAIN /user/get/ . IT HANDLES:
// IT RETURNS OWN USER DATA IN RESPONSE BY POST METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) HandlerUserGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "POST" {
		server.userGet(w, r)
	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)
		server.Log(r, "client wants to know what methods are allowed")
	} else {
		// ALL OTHERS METHODS
		server.LogWarning(r, "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
