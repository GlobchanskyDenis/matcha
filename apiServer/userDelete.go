package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"MatchaServer/handlers"
	"context"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /user/delete/
func (server *Server) UserDelete(w http.ResponseWriter, r *http.Request) {
	var (
		err                 error
		requestParams       map[string]interface{}
		item                interface{}
		ctx                 context.Context
		pass, encryptedPass string
		uid                 int
		isExist, ok         bool
	)

	ctx = r.Context()
	requestParams = ctx.Value("requestParams").(map[string]interface{})
	uid = ctx.Value("uid").(int)

	item, isExist = requestParams["pass"]
	if !isExist {
		server.Logger.LogWarning(r, "password not exist")
		server.error(w, errors.NoArgument.WithArguments("Поле pass отсутствует", "pass field expected"))
		return
	}

	pass, ok = item.(string)
	if !ok {
		server.Logger.LogWarning(r, "password have wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Поле pass имеет неверный тип", "pass field has wrong type"))
		return
	}

	encryptedPass = handlers.PassHash(pass)

	user, err := server.Db.GetUserByUid(uid)
	if err != nil {
		server.Logger.LogError(r, "GetUserByUid returned error - "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	if encryptedPass != user.EncryptedPass {
		server.Logger.LogWarning(r, "password is incorrect")
		server.error(w, errors.InvalidArgument.WithArguments("неверный пароль", "password is wrong"))
		return
	}

	server.Session.DeleteUserSessionByUid(user.Uid)

	err = server.Db.DeleteUser(user.Uid)
	if err != nil {
		server.Logger.LogError(r, "DeleteUser returned error - "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
		" was removed successfully. No response body")
}
