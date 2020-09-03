package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errDef"
	"MatchaServer/handlers"
	"encoding/json"
	"net/http"
	"strconv"
)

// USER MAIL CONFIRM BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) userUpdateStatus(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, token string
		err                  error
		request              map[string]interface{}
		item                 interface{}
		isExist, ok          bool
	)

	message = "request for MAIL CONFIRM was recieved"
	server.Log(r, message)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		server.LogError(r, "request json decode failed - "+err.Error())
		server.error(w, errDef.InvalidRequestBody)
		return
	}

	item, isExist = request["x-reg-token"]
	if !isExist {
		server.LogError(r, "x-reg-token not exist in request")
		server.error(w, errDef.NoArgument.WithArguments("Поле x-reg-token отсутствует", "x-reg-token field expected"))
		return
	}

	token, ok = item.(string)
	if !ok {
		server.LogError(r, "x-reg-token has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле x-reg-token имеет неверный тип", "x-reg-token field has wrong type"))
		return
	}

	if token == "" {
		server.LogError(r, "x-reg-token is empty")
		server.error(w, errDef.UserNotLogged)
		return
	}

	mail, err = handlers.TokenMailDecode(token)
	if err != nil {
		server.LogWarning(r, "TokenMailDecode returned error - "+err.Error())
		server.error(w, errDef.UserNotLogged)
		return
	}

	user, err := server.Db.GetUserByMail(mail)
	if errDef.RecordNotFound.IsOverlapWithError(err) {
		server.LogWarning(r, "GetUserByMail - record not found")
		server.error(w, errDef.UserNotExist)
		return
	} else if err != nil {
		server.LogWarning(r, "GetUserByMail returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	user.Status = "confirmed"

	err = server.Db.UpdateUser(user)
	if err != nil {
		server.LogWarning(r, "UpdateUser returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	server.LogSuccess(r, "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
		" was updated its status successfully. No response body")
}

// HTTP HANDLER FOR DOMAIN /user/update/status . IT HANDLES:
// UPDATE USER STATUS BY PATCH METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) HandlerUserUpdateStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "PATCH" {
		server.userUpdateStatus(w, r)
	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)
		server.Log(r, "client wants to know what methods are allowed")
	} else {
		// ALL OTHERS METHODS
		server.LogWarning(r, "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
