package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errDef"
	"MatchaServer/handlers"
	"net/http"
	"strconv"
	"context"
)

// HTTP HANDLER FOR DOMAIN /user/delete/
func (server *Server) UserDelete(w http.ResponseWriter, r *http.Request) {
	var (
		err                   error
		requestParams       map[string]interface{}
		item				interface{}
		ctx					context.Context
		pass, encryptedPass string
		uid                   int
		isExist, ok         bool
	)

	ctx = r.Context()
	requestParams = ctx.Value("requestParams").(map[string]interface{})
	uid = ctx.Value("uid").(int)

	item, isExist = requestParams["pass"]
	if !isExist {
		server.LogWarning(r, "password not exist")
		server.error(w, errDef.NoArgument.WithArguments("Поле pass отсутствует", "pass field expected"))
		return
	}

	pass, ok = item.(string)
	if !ok {
		server.LogWarning(r, "password have wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле pass имеет неверный тип", "pass field has wrong type"))
		return
	}

	encryptedPass = handlers.PassHash(pass)

	user, err := server.Db.GetUserByUid(uid)
	if err != nil {
		server.LogError(r, "GetUserByUid returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	if encryptedPass != user.EncryptedPass {
		server.LogWarning(r, "password is incorrect")
		server.error(w, errDef.InvalidArgument.WithArguments("неверный пароль", "password is wrong"))
		return
	}

	server.session.DeleteUserSessionByUid(user.Uid)

	err = server.Db.DeleteUser(user.Uid)
	if err != nil {
		server.LogError(r, "DeleteUser returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	server.LogSuccess(r, "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
		" was removed successfully. No response body")
}
