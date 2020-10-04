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
func (server *Server) UserUpdateStatus(w http.ResponseWriter, r *http.Request) {
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
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	user.Status = "confirmed"

	err = server.Db.UpdateUser(user)
	if err != nil {
		server.Logger.LogWarning(r, "UpdateUser returned error - "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
		" was updated its status successfully. No response body")
}
