package apiServer

import (
	"MatchaServer/errors"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /interests/get/ . IT HANDLES:
// RETURNING ARRAY OF EXISTING INTERESTS
func (server *Server) InterestsGet(w http.ResponseWriter, r *http.Request) {
	var err error

	interests, err := server.Db.GetInterests()
	if err != nil {
		server.Logger.LogError(r, "database returned error - "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	jsonInterests, err := json.Marshal(interests)
	if err != nil {
		server.Logger.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errors.MarshalError)
		return
	}

	// This is my valid case.
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonInterests)
	server.Logger.LogSuccess(r, "Interests was returned to user. Amount "+strconv.Itoa(len(interests)))
}
