package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"MatchaServer/handlers"
	"context"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /user/update/status/
// USER MAIL CONFIRM. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) userUpdateStatusPatch(w http.ResponseWriter, r *http.Request) {
	var (
		mail, token   string
		err           error
		requestParams map[string]interface{}
		item          interface{}
		ctx           context.Context
		isExist, ok   bool
	)

	ctx = r.Context()
	requestParams = ctx.Value("requestParams").(map[string]interface{})

	item, isExist = requestParams["x-reg-token"]
	if !isExist {
		server.Logger.LogError(r, "x-reg-token not exist in request")
		server.error(w, errors.NoArgument.WithArguments("Поле x-reg-token отсутствует", "x-reg-token field expected"))
		return
	}

	token, ok = item.(string)
	if !ok {
		server.Logger.LogError(r, "x-reg-token has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Поле x-reg-token имеет неверный тип",
			"x-reg-token field has wrong type"))
		return
	}

	if token == "" {
		server.Logger.LogError(r, "x-reg-token is empty")
		server.error(w, errors.UserNotLogged)
		return
	}

	mail, err = handlers.TokenMailDecode(token)
	if err != nil {
		server.Logger.LogWarning(r, "TokenMailDecode returned error - "+err.Error())
		server.error(w, errors.UserNotLogged)
		return
	}

	user, err := server.Db.GetUserByMail(mail)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "GetUserByMail - record not found")
		server.error(w, errors.UserNotExist)
		return
	} else if err != nil {
		server.Logger.LogWarning(r, "GetUserByMail returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	user.Status = "confirmed"

	err = server.Db.UpdateUser(user)
	if err != nil {
		server.Logger.LogWarning(r, "UpdateUser returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
		" was updated its status successfully. No response body")
}

// HTTP HANDLER FOR DOMAIN /user/update/status/
// USER MAIL CONFIRM. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) userUpdateStatusPost(w http.ResponseWriter, r *http.Request) {
	var (
		mail, token   string
		slice         []string
		err           error
		requestParams map[string]interface{}
		item          interface{}
		ctx           context.Context
		isExist, ok   bool
	)

	ctx = r.Context()
	requestParams = ctx.Value("requestParams").(map[string]interface{})

	item, isExist = requestParams["x-reg-token"]
	if !isExist {
		server.Logger.LogError(r, "x-reg-token not exist in request. Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	}

	slice, ok = item.([]string)
	if !ok {
		server.Logger.LogError(r, "x-reg-token slice has wrong type. Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	}

	if len(slice) != 1 {
		server.Logger.LogError(r, "x-reg-token slice has wrong length. Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	}

	token = slice[0]

	if token == "" {
		server.Logger.LogError(r, "x-reg-token is empty. Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	}

	mail, err = handlers.TokenMailDecode(token)
	if err != nil {
		server.Logger.LogWarning(r, "TokenMailDecode returned error - "+err.Error()+". Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	}

	user, err := server.Db.GetUserByMail(mail)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "GetUserByMail - record not found. Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	} else if err != nil {
		server.Logger.LogWarning(r, "GetUserByMail returned error - "+err.Error()+". Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	}

	user.Status = "confirmed"

	err = server.Db.UpdateUser(user)
	if err != nil {
		server.Logger.LogWarning(r, "UpdateUser returned error - "+err.Error()+". Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	}

	http.Redirect(w, r, "http://localhost:3001/confirm/mail/success", http.StatusFound)
	server.Logger.LogSuccess(r, "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
		" was updated its status successfully. Redirect to front page success")
}

// HTTP HANDLER FOR DOMAIN /user/update/status/
// USER MAIL CONFIRM. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) userUpdateStatusGet(w http.ResponseWriter, r *http.Request) {
	var (
		mail, token string
		err         error
	)

	token = r.URL.Query().Get("x-reg-token")

	if token == "" {
		server.Logger.LogError(r, "x-reg-token is empty. Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	}

	mail, err = handlers.TokenMailDecode(token)
	if err != nil {
		server.Logger.LogWarning(r, "TokenMailDecode returned error - "+err.Error()+". Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	}

	user, err := server.Db.GetUserByMail(mail)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "GetUserByMail - record not found. Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	} else if err != nil {
		server.Logger.LogWarning(r, "GetUserByMail returned error - "+err.Error()+". Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	}

	user.Status = "confirmed"

	err = server.Db.UpdateUser(user)
	if err != nil {
		server.Logger.LogWarning(r, "UpdateUser returned error - "+err.Error()+". Redirect to front page fail")
		http.Redirect(w, r, "http://localhost:3001/confirm/mail/fail", http.StatusFound)
		return
	}

	http.Redirect(w, r, "http://localhost:3001/confirm/mail/success", http.StatusFound)
	server.Logger.LogSuccess(r, "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
		" was updated its status successfully. Redirect to front page success")
}

// HTTP HANDLER FOR DOMAIN /user/update/status/
// USER MAIL CONFIRM. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) UserUpdateStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PATCH" {
		server.userUpdateStatusPatch(w, r)
	} else if r.Method == "POST" {
		server.userUpdateStatusPost(w, r)
	} else if r.Method == "GET" {
		server.userUpdateStatusGet(w, r)
	}
}
