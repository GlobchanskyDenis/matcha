package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /history/myViews/ . IT HANDLES:
// IT RETURNS LIST OF PREVIOUS VIEWS OF USERS BY POST METHOD.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) HistoryViews(w http.ResponseWriter, r *http.Request) {
	var (
		history []HistoryReference
		uid     int
		err     error
		ctx     context.Context
	)

	ctx = r.Context()
	uid = ctx.Value("uid").(int)

	history, err = server.Db.GetHistoryReferencesByUid(uid)
	if err != nil {
		server.Logger.LogError(r, "GetHistoryReferencesByUid returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	jsonHistory, err := json.Marshal(history)
	if err != nil {
		server.Logger.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errors.MarshalError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonHistory)
	server.Logger.LogSuccess(r, "History of user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+
		" was found and transmitted successfully. Views amount is #"+BLUE+strconv.Itoa(len(history))+NO_COLOR)
}
