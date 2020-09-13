package apiServer

import (
	"MatchaServer/errDef"
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
		server.LogError(r, "database returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	jsonInterests, err := json.Marshal(interests)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		server.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errDef.MarshalError)
		return
	}

	// This is my valid case.
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonInterests)
	server.LogSuccess(r, "Interests was returned to user. Amount "+strconv.Itoa(len(interests)))
}
