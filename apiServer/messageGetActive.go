package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /message/get/active/ . IT HANDLES:
// IT RETURNS LIST OF MESSAGES THAT HAVE BEEN NOT SEEN YET.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) MessageGetActive(w http.ResponseWriter, r *http.Request) {
	var (
		messages        []Message
		myUid           int
		err             error
		ctx             context.Context
	)

	ctx = r.Context()
	myUid = ctx.Value("uid").(int)

	messages, err = server.Db.GetActiveMessages(myUid)
	if err != nil {
		server.Logger.LogError(r, "GetActiveMessages returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	jsonMessages, err := json.Marshal(messages)
	if err != nil {
		server.Logger.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errors.MarshalError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonMessages)
	server.Logger.LogSuccess(r, "Messages was handled successfully. Uid #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
		". Amount is #"+BLUE+strconv.Itoa(len(messages))+NO_COLOR)
}
