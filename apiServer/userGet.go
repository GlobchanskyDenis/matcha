package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errDef"
	"context"
	"encoding/json"
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
	if errDef.RecordNotFound.IsOverlapWithError(err) {
		server.LogWarning(r, "GetUserByUid - record not found")
		server.error(w, errDef.UserNotExist)
		return
	} else if err != nil {
		server.LogError(r, "GetUser returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		server.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errDef.MarshalError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonUser)
	server.LogSuccess(r, "User "+BLUE+mail+NO_COLOR+" was authenticated successfully")
}
