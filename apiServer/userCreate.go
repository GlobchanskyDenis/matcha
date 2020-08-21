package apiServer

import (
	"MatchaServer/config"
	"MatchaServer/errDef"
	"MatchaServer/handlers"
	"encoding/json"
	"net/http"
)

// USER REGISTRATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) userReg(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, pass, token string
		err                        error
		request                    map[string]interface{}
		isExist, ok                bool
		user                       config.User
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/create/", "request json decode failed - "+err.Error())
		server.error(w, errDef.InvalidRequestBody)
		return
	}

	arg, isExist := request["mail"]
	if !isExist {
		consoleLogWarning(r, "/user/create/", "mail not exist")
		server.error(w, errDef.NoArgument.WithArguments("Поле mail отсутствует", "mail field expected"))
		return
	}

	mail, ok = arg.(string)
	if !ok {
		consoleLogWarning(r, "/user/create/", "mail has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле mail имеет неверный тип", "mail field has wrong type"))
		return
	}

	arg, isExist = request["pass"]
	if !isExist {
		consoleLogWarning(r, "/user/create/", "password not exist")
		server.error(w, errDef.NoArgument.WithArguments("Поле pass отсутствует", "pass field expected"))
		return
	}

	pass, ok = arg.(string)
	if !ok {
		consoleLogWarning(r, "/user/create/", "password has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле pass имеет неверный тип", "pass field has wrong type"))
		return
	}

	message = "request was recieved, mail: "+BLUE+mail+NO_COLOR+" password: hidden"
	consoleLog(r, "/user/create/", message)

	if mail == "" || pass == "" {
		consoleLogWarning(r, "/user/create/", "mail or password is empty")
		server.error(w, errDef.InvalidArgument.WithArguments("логин или пароль пусты", "login or password is empty"))
		return
	}

	err = handlers.CheckMail(mail)
	if err != nil {
		consoleLogWarning(r, "/user/create/", "mail - "+err.Error())
		// add apiErrorArgument in handlers functions
		server.error(w, errDef.InvalidArgument.WithArguments(err))
		return
	}

	err = handlers.CheckPass(pass)
	if err != nil {
		consoleLogWarning(r, "/user/create/", "password - "+err.Error())
		// add apiErrorArgument in handlers functions
		server.error(w, errDef.InvalidArgument.WithArguments(err))
		return
	}

	isUserExists, err := server.Db.IsUserExistsByMail(mail)
	if err != nil {
		consoleLogError(r, "/user/create/", "IsUserExists returned error "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}
	if isUserExists {
		consoleLogWarning(r, "/user/create/", "user "+BLUE+mail+NO_COLOR+" alredy exists")
		server.error(w, errDef.RegFailUserExists)
		return
	}

	user, err = server.Db.SetNewUser(mail, handlers.PassHash(pass))
	if err != nil {
		consoleLogError(r, "/user/create/", "SetNewUser returned error "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	token, err = handlers.TokenMailEncode(mail)
	if err != nil {
		consoleLogError(r, "/user/create/", "TokenMailEncode returned error "+err.Error())
		server.error(w, errDef.MarshalError)
		return
	}

	err = server.Db.SetNewDevice(user.Uid, r.UserAgent())
	if err != nil {
		consoleLogError(r, "/user/create/", "SetNewDevice returned error "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	w.WriteHeader(201)
	consoleLogSuccess(r, "/user/create/", "user "+BLUE+mail+NO_COLOR+" was created successfully. No response body")

	go func(mail string, xRegToken string, r *http.Request) {
		err := handlers.SendMail(mail, xRegToken)
		if err != nil {
			consoleLogError(r, "/user/create/", "SendMail returned error "+err.Error())
		} else {
			consoleLogSuccess(r, "/user/create/", "Confirm mail for user "+BLUE+mail+NO_COLOR+" was send successfully")
		}
	}(mail, token, r)
}

// HTTP HANDLER FOR DOMAIN /user/reg
// REGISTRATE USER BY POST METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) HandlerUserCreate(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST,PATCH,OPTIONS,DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "POST" {

		server.userReg(w, r)

	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)

		consoleLog(r, "/user/create/", "client wants to know what methods are allowed")

	} else {
		// ALL OTHERS METHODS

		consoleLogWarning(r, "/user/create/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}
