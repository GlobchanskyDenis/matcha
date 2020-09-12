package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"MatchaServer/errDef"
	"MatchaServer/handlers"
	"net/http"
	"context"
)

// HTTP HANDLER FOR DOMAIN /user/create/
func (server *Server) UserCreate(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, pass, token string
		err                        error
		requestParams              map[string]interface{}
		item					   interface{}
		ctx						   context.Context
		isExist, ok                bool
		user                       User
	)

	ctx = r.Context()
	requestParams = ctx.Value("requestParams").(map[string]interface{})

	item, isExist = requestParams["mail"]
	if !isExist {
		server.LogWarning(r, "mail not exist")
		server.error(w, errDef.NoArgument.WithArguments("Поле mail отсутствует", "mail field expected"))
		return
	}

	mail, ok = item.(string)
	if !ok {
		server.LogWarning(r, "mail has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле mail имеет неверный тип", "mail field has wrong type"))
		return
	}

	item, isExist = requestParams["pass"]
	if !isExist {
		server.LogWarning(r, "password not exist")
		server.error(w, errDef.NoArgument.WithArguments("Поле pass отсутствует", "pass field expected"))
		return
	}

	pass, ok = item.(string)
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

	isExist, err = server.Db.IsUserExistsByMail(mail)
	if err != nil {
		server.LogError(r, "IsUserExists returned error "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}
	if isExist {
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
