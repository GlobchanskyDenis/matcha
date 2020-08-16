package apiServer

import (
	"MatchaServer/errDef"
	"MatchaServer/handlers"
	"encoding/json"
	"net/http"
	"strconv"
)

// USER REMOVE BY DELETE METHOD. NO REQUEST BODY. RESPONSE BODY IS JSON ONLY IN CASE OF ERROR.
// REQUEST SHOULD HAVE 'x-auth-token' HEADER
func (server *Server) userDelete(w http.ResponseWriter, r *http.Request) {
	var (
		message, token      string
		err                 error
		request             map[string]interface{}
		pass, encryptedPass string
		uid                 int
		isLogged, isExist, ok bool
	)

	message = "request for DELETE was recieved"
	consoleLog(r, "/user/delete/", message)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/delete/", "request json decode failed - "+err.Error())
		server.error(w, errDef.InvalidRequestBody)
		return
	}

	arg, isExist := request["x-auth-token"]
	if !isExist {
		consoleLogWarning(r, "/user/delete/", "x-auth-token not exist")
		server.error(w, errDef.NoArgument.WithArguments("Поле x-auth-token отсутствует", "x-auth-token field expected"))
		return
	}

	token, ok = arg.(string)
	if !ok {
		consoleLogWarning(r, "/user/delete/", "token have wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле x-auth-token имеет неверный тип", "x-auth-token field has wrong type"))
		return
	}

	if token == "" {
		consoleLogWarning(r, "/user/delete/", "token is empty")
		server.error(w, errDef.UserNotLogged)
		return
	}

	uid, err = handlers.TokenUidDecode(token)
	if err != nil {
		consoleLogWarning(r, "/user/delete/", "TokenDecode returned error - "+err.Error())
		server.error(w, errDef.UserNotLogged)
		return
	}

	isLogged = server.session.IsUserLoggedByUid(uid)
	if !isLogged {
		consoleLogWarning(r, "/photo/upload/", "User #"+strconv.Itoa(uid)+" is not logged")
		server.error(w, errDef.UserNotLogged)
		return
	}

	arg, isExist = request["pass"]
	if !isExist {
		consoleLogWarning(r, "/user/delete/", "password not exist")
		server.error(w, errDef.NoArgument.WithArguments("Поле pass отсутствует", "pass field expected"))
		return
	}

	pass, ok = arg.(string)
	if !ok {
		consoleLogWarning(r, "/user/delete/", "password have wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле pass имеет неверный тип", "pass field has wrong type"))
		return
	}

	encryptedPass = handlers.PassHash(pass)

	user, err := server.Db.GetUserByUid(uid)
	if err != nil {
		consoleLogError(r, "/user/delete/", "GetUserByUid returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	if encryptedPass != user.EncryptedPass {
		consoleLogWarning(r, "/user/delete/", "password is incorrect")
		server.error(w, errDef.InvalidArgument.WithArguments("неверный пароль", "password is wrong"))
		return
	}

	server.session.DeleteUserSessionByUid(user.Uid)

	err = server.Db.DeleteUser(user.Uid)
	if err != nil {
		consoleLogError(r, "/user/delete/", "DeleteUser returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	consoleLogSuccess(r, "/user/delete/", "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
		" was removed successfully. No response body")
}

// HTTP HANDLER FOR DOMAIN /user/delete/
// DELETE USER BY DELETE METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) HandlerUserDelete(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "OPTIONS,DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)

		consoleLog(r, "/user/delete/", "client wants to know what methods are allowed")

	} else if r.Method == "DELETE" {

		server.userDelete(w, r)

	} else {
		// ALL OTHERS METHODS

		consoleLogWarning(r, "/user/delete/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}
