package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /user/get/likedMe/ . IT HANDLES:
// IT RETURNS LIST OF USERS THAT YOU ADDED INTO IGNORE.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) UserGetLikedMe(w http.ResponseWriter, r *http.Request) {
	var (
		users []User
		uid   int
		err   error
		ctx   context.Context
	)

	ctx = r.Context()
	uid = ctx.Value("uid").(int)

	users, err = server.Db.GetUsersLikedMe(uid)
	if err != nil {
		server.Logger.LogError(r, "GetIgnoredUsers returned error - "+err.Error())
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
