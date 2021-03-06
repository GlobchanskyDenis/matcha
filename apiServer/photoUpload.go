package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /photo/upload/ . IT HANDLES:
// UPLOAD USER PHOTO
func (server *Server) PhotoUpload(w http.ResponseWriter, r *http.Request) {
	var (
		src           string
		uid, pid      int
		err           error
		requestParams map[string]interface{}
		item          interface{}
		ctx           context.Context
		isExist, ok   bool
	)

	ctx = r.Context()
	requestParams = ctx.Value("requestParams").(map[string]interface{})
	uid = ctx.Value("uid").(int)

	item, isExist = requestParams["src"]
	if !isExist {
		server.Logger.LogWarning(r, "src not exist in request")
		server.error(w, errors.NoArgument.WithArguments("Поле src отсутствует", "src field expected"))
		return
	}

	src, ok = item.(string)
	if !ok {
		server.Logger.LogWarning(r, "src has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Поле src имеет неверный тип", "src field has wrong type"))
		return
	}

	pid, err = server.Db.SetNewPhoto(uid, src)
	if err != nil {
		server.Logger.LogError(r, "SetNewPhoto returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+
		" was uploaded its photo successfully. photo id #"+BLUE+strconv.Itoa(pid)+NO_COLOR+". No response body")
}
