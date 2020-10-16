package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"strconv"
	"net/http"
)

// HTTP HANDLER FOR DOMAIN /user/get/ . IT HANDLES:
// IT RETURNS OWN USER DATA IN RESPONSE BY POST METHOD.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) UserGet(w http.ResponseWriter, r *http.Request) {
	var (
		user User
		uid  int
		err  error
		ctx  context.Context
	)

	ctx = r.Context()
	uid = ctx.Value("uid").(int)

	user, err = server.Db.GetUserByUid(uid)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "GetUserByUid - record not found")
		server.error(w, errors.UserNotExist)
		return
	} else if err != nil {
		server.Logger.LogError(r, "GetUser returned error - "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		server.Logger.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errors.MarshalError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonUser)
	server.Logger.LogSuccess(r, "User #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+" was found and transmitted successfully")
}
