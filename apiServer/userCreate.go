package apiServer

import (
	. "MatchaServer/common"
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
		user                       User
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		server.LogError(r, "request json decode failed - "+err.Error())
		server.error(w, errDef.InvalidRequestBody)
		return
	}

	arg, isExist := request["mail"]
	if !isExist {
		server.LogWarning(r, "mail not exist")
		server.error(w, errDef.NoArgument.WithArguments("Поле mail отсутствует", "mail field expected"))
		return
	}

	mail, ok = arg.(string)
	if !ok {
		server.LogWarning(r, "mail has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле mail имеет неверный тип", "mail field has wrong type"))
		return
	}

	arg, isExist = request["pass"]
	if !isExist {
		server.LogWarning(r, "password not exist")
		server.error(w, errDef.NoArgument.WithArguments("Поле pass отсутствует", "pass field expected"))
		return
	}

	pass, ok = arg.(string)
	if !ok {
		server.LogWarning(r, "password has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле pass имеет неверный тип", "pass field has wrong type"))
		return
	}

	message = "request was recieved, mail: " + BLUE + mail + NO_COLOR + " password: hidden"
	server.Log(r, message)

	if mail == "" || pass == "" {
		server.LogWarning(r, "mail or password is empty")
		server.error(w, errDef.InvalidArgument.WithArguments("логин или пароль пусты", "login or password is empty"))
		return
	}

	err = handlers.CheckMail(mail)
	if err != nil {
		server.LogWarning(r, "mail - "+err.Error())
		// add apiErrorArgument in handlers functions
		server.error(w, errDef.InvalidArgument.WithArguments(err))
		return
	}

	err = handlers.CheckPass(pass)
	if err != nil {
		server.LogWarning(r, "password - "+err.Error())
		// add apiErrorArgument in handlers functions
		server.error(w, errDef.InvalidArgument.WithArguments(err))
		return
	}

	isUserExists, err := server.Db.IsUserExistsByMail(mail)
	if err != nil {
		server.LogError(r, "IsUserExists returned error "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}
	if isUserExists {
		server.LogWarning(r, "user "+BLUE+mail+NO_COLOR+" alredy exists")
		server.error(w, errDef.RegFailUserExists)
		return
	}

	user, err = server.Db.SetNewUser(mail, handlers.PassHash(pass))
	if err != nil {
		server.LogError(r, "SetNewUser returned error "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	token, err = handlers.TokenMailEncode(mail)
	if err != nil {
		server.LogError(r, "TokenMailEncode returned error "+err.Error())
		server.error(w, errDef.MarshalError)
		return
	}

	err = server.Db.SetNewDevice(user.Uid, r.UserAgent())
	if err != nil {
		server.LogError(r, "SetNewDevice returned error "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	w.WriteHeader(201)
	server.LogSuccess(r, "user "+BLUE+mail+NO_COLOR+" was created successfully. No response body")

	go func(mail string, xRegToken string, r *http.Request, mailConf *config.Mail) {
		err := handlers.SendMail(mail, xRegToken, mailConf)
		if err != nil {
			server.LogError(r, "SendMail returned error "+err.Error())
		} else {
			server.LogSuccess(r, "Confirm mail for user "+BLUE+mail+NO_COLOR+" was send successfully")
		}
	}(mail, token, r, &server.mailConf)
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

		server.Log(r, "client wants to know what methods are allowed")

	} else {
		// ALL OTHERS METHODS

		server.LogWarning(r, "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}
