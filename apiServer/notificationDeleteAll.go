package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /notification/delete/all/ . IT HANDLES:
// IT DELETES ALL NOTIFICATIONS OF USER.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) NotificationDeleteAll(w http.ResponseWriter, r *http.Request) {
	var (
		myUid int
		err   error
		ctx   context.Context
	)

	ctx = r.Context()
	myUid = ctx.Value("uid").(int)

	err = server.Db.DropReceiverNotifs(myUid)
	if err != nil {
		server.Logger.LogError(r, "DropReceiverNotifs returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "All notifications of user #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+" was deleted successfully.")
}
