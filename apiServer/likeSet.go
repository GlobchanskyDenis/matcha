package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"strconv"
	"net/http"
)

// HTTP HANDLER FOR DOMAIN /message/get/ . IT HANDLES:
// IT RETURNS OWN USER DATA IN RESPONSE BY POST METHOD.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) LikeSet(w http.ResponseWriter, r *http.Request) {
	var (
		uid64			float64
		myUid, otherUid int
		requestParams	map[string]interface{}
		item			interface{}
		ok, isExist     bool
		err  error
		ctx  context.Context
	)

	ctx = r.Context()
	myUid = ctx.Value("uid").(int)
	requestParams, ok = ctx.Value("requestParams").(map[string]interface{})
	if !ok {
		server.Logger.LogWarning(r, "Request params has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Параметры запроса имеют неверный тип", "request params has wrong type"))
		return
	}
	item, isExist = requestParams["otherUid"]
	if !isExist {
		server.Logger.LogWarning(r, "Param otherUid expected")
		server.error(w, errors.NoArgument.WithArguments("отсутствует параметр otherUid", "param otherUid expected"))
		return
	}
	uid64, ok = item.(float64)
	if !ok {
		server.Logger.LogWarning(r, "Id of another user has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("otherUid имеет неверный тип", "otherUid has wrong type"))
		return
	}
	otherUid = int(uid64)

	err = server.Db.SetNewLike(myUid, otherUid)
	if errors.ImpossibleToExecute.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "Imposible to set like from user#"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
			" to user#"+BLUE+strconv.Itoa(otherUid)+NO_COLOR)
		server.error(w, errors.ImpossibleToExecute.WithArguments("Лайк уже существует", "like already exists"))
		return
	} else if err != nil {
		server.Logger.LogError(r, "SetNewLike returned error - "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "Like was set successfully from user #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
		" to user #"+BLUE+strconv.Itoa(otherUid)+NO_COLOR)
}