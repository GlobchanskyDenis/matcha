package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"MatchaServer/errors"
	"MatchaServer/handlers"
	"context"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /user/create/
// IT CREATES USER WITH MAIL AND PASSWORD
func (server *Server) UserCreate(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, pass, token string
		err                        error
		requestParams              map[string]interface{}
		item                       interface{}
		ctx                        context.Context
		isExist, ok                bool
		user                       User
	)

	ctx = r.Context()
	requestParams = ctx.Value("requestParams").(map[string]interface{})

	item, isExist = requestParams["mail"]
	if !isExist {
		server.Logger.LogWarning(r, "mail not exist")
		server.error(w, errors.NoArgument.WithArguments("Поле mail отсутствует", "mail field expected"))
		return
	}

	mail, ok = item.(string)
	if !ok {
		server.Logger.LogWarning(r, "mail has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Поле mail имеет неверный тип", "mail field has wrong type"))
		return
	}

	item, isExist = requestParams["pass"]
	if !isExist {
		server.Logger.LogWarning(r, "password not exist")
		server.error(w, errors.NoArgument.WithArguments("Поле pass отсутствует", "pass field expected"))
		return
	}

	pass, ok = item.(string)
	if !ok {
		server.Logger.LogWarning(r, "password has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Поле pass имеет неверный тип", "pass field has wrong type"))
		return
	}

	message = "request was recieved, mail: " + BLUE + mail + NO_COLOR + " password: hidden"
	server.Logger.Log(r, message)

	if mail == "" || pass == "" {
		server.Logger.LogWarning(r, "mail or password is empty")
		server.error(w, errors.InvalidArgument.WithArguments("логин или пароль пусты", "login or password is empty"))
		return
	}

	err = handlers.CheckMail(mail)
	if err != nil {
		server.Logger.LogWarning(r, "mail - "+err.Error())
		server.error(w, errors.InvalidArgument.WithArguments(err))
		return
	}

	err = handlers.CheckPass(pass)
	if err != nil {
		server.Logger.LogWarning(r, "password - "+err.Error())
		server.error(w, errors.InvalidArgument.WithArguments(err))
		return
	}

	isExist, err = server.Db.IsUserExistsByMail(mail)
	if err != nil {
		server.Logger.LogError(r, "IsUserExistsByMail returned error "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}
	if isExist {
		server.Logger.LogWarning(r, "user "+BLUE+mail+NO_COLOR+" alredy exists")
		server.error(w, errors.RegFailUserExists)
		return
	}

	user, err = server.Db.SetNewUser(mail, handlers.PassHash(pass))
	if err != nil {
		server.Logger.LogError(r, "SetNewUser returned error "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	token, err = handlers.TokenMailEncode(mail)
	if err != nil {
		server.Logger.LogError(r, "TokenMailEncode returned error "+err.Error())
		server.error(w, errors.MarshalError)
		return
	}

	err = server.Db.SetNewDevice(user.Uid, r.UserAgent())
	if err != nil {
		server.Logger.LogError(r, "SetNewDevice returned error "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	w.WriteHeader(201)
	w.Write([]byte(`{"uid":` + strconv.Itoa(user.Uid) + `}`))
	server.Logger.LogSuccess(r, "user "+BLUE+mail+NO_COLOR+" was created successfully. Uid #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR)

	go func(mail string, xRegToken string, r *http.Request, mailConf *config.Mail) {
		err := handlers.SendMail(mail, xRegToken, mailConf)
		if err != nil {
			server.Logger.LogError(r, "SendMail returned error "+err.Error())
		} else {
			server.Logger.LogSuccess(r, "Confirm mail for user "+BLUE+mail+NO_COLOR+" was send successfully")
		}
	}(mail, token, r, &server.mailConf)
}
