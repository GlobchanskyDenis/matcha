package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /user/get/claimed/ . IT HANDLES:
// IT RETURNS LIST OF USERS THAT YOU ADDED INTO BLACK LIST.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) UserGetClaimed(w http.ResponseWriter, r *http.Request) {
	var (
		users []User
		uid   int
		err   error
		ctx   context.Context
	)

	ctx = r.Context()
	uid = ctx.Value("uid").(int)

	users, err = server.Db.GetClaimedUsers(uid)
	if err != nil {
		server.Logger.LogError(r, "GetClaimedUsers returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		server.Logger.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errors.MarshalError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonUsers)
	server.Logger.LogSuccess(r, "Users was found successfully. Total amount "+BLUE+strconv.Itoa(len(users))+NO_COLOR)
}
