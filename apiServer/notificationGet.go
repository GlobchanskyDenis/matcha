package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /notification/get/ . IT HANDLES:
// IT REMOVES NOTIFICATION BY ITS ID.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) NotificationGet(w http.ResponseWriter, r *http.Request) {
	var (
		notifs []Notif
		myUid  int
		err    error
		ctx    context.Context
	)

	ctx = r.Context()
	myUid = ctx.Value("uid").(int)

	notifs, err = server.Db.GetNotifByUidReceiver(myUid)
	if err != nil {
		server.Logger.LogError(r, "GetNotifByUidReceiver returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	jsonNotifs, err := json.Marshal(notifs)
	if err != nil {
		server.Logger.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errors.MarshalError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonNotifs)
	server.Logger.LogSuccess(r, "Notifications was handled successfully. Uid #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
		". Amount is #"+BLUE+strconv.Itoa(len(notifs))+NO_COLOR)
}
