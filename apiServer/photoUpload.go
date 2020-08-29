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
func (server *Server) photoUpload(w http.ResponseWriter, r *http.Request) {
	var (
		message, body, token  string
		uid, pid              int
		err                   error
		request               map[string]interface{}
		item                  interface{}
		isExist, isLogged, ok bool
	)

	message = "request for PHOTO UPLOAD was recieved"
	consoleLog(r, "/photo/upload/", message)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/photo/upload/", "request json decode failed - "+err.Error())
		server.error(w, errDef.InvalidRequestBody)
		return
	}

	item, isExist = request["x-auth-token"]
	if !isExist {
		consoleLogWarning(r, "/photo/upload/", "x-auth-token not exist in request")
		server.error(w, errDef.NoArgument.WithArguments("Поле x-auth-token отсутствует", "x-auth-token field expected"))
		return
	}

	token, ok = item.(string)
	if !ok {
		consoleLogWarning(r, "/photo/upload/", "x-auth-token has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле x-auth-token имеет неверный тип", "x-auth-token field has wrong type"))
		return
	}

	if token == "" {
		consoleLogWarning(r, "/photo/upload/", "x-auth-token is empty")
		server.error(w, errDef.UserNotLogged)
		return
	}

	uid, err = handlers.TokenUidDecode(token)
	if err != nil {
		consoleLogWarning(r, "/photo/upload/", "TokenUidDecode returned error - "+err.Error())
		server.error(w, errDef.UserNotLogged)
		return
	}

	item, isExist = request["src"]
	if !isExist {
		consoleLogWarning(r, "/photo/upload/", "src not exist in request")
		server.error(w, errDef.NoArgument.WithArguments("Поле src отсутствует", "src field expected"))
		return
	}

	body, ok = item.(string)
	if !ok {
		consoleLogWarning(r, "/photo/upload/", "src has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле src имеет неверный тип", "src field has wrong type"))
		return
	}

	isLogged = server.session.IsUserLoggedByUid(uid)
	if !isLogged {
		consoleLogWarning(r, "/photo/upload/", "User #"+strconv.Itoa(uid)+" is not logged")
		server.error(w, errDef.UserNotLogged)
		return
	}

	pid, err = server.Db.SetNewPhoto(uid, body)
	if err != nil {
		consoleLogError(r, "/photo/upload/", "UpdateUser returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	consoleLogSuccess(r, "/photo/upload/", "user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+
		" was uploaded its photo successfully. photo id #"+BLUE+strconv.Itoa(pid)+NO_COLOR+". No response body")
}

// HTTP HANDLER FOR DOMAIN /user/update/status . IT HANDLES:
// UPDATE USER STATUS BY PATCH METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) HandlerPhotoUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		server.photoUpload(w, r)
	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)
		consoleLog(r, "/photo/upload/", "client wants to know what methods are allowed")
	} else {
		// ALL OTHERS METHODS
		consoleLogWarning(r, "/photo/upload/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
