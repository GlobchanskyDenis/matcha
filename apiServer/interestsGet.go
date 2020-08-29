package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errDef"
	"encoding/json"
	"net/http"
	"strconv"
)

// USER AUTHORISATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) interestsGet(w http.ResponseWriter, r *http.Request) {
	var (
		message string
		err     error
	)

	defer func() {
		if err := recover(); err != nil {
			println(RED_BG + "PANIC!!!!! " + err.(error).Error() + NO_COLOR)
		}
	}()

	message = "request for interests array was recieved"
	server.Log(r, "/interests/get/", message)

	interests, err := server.Db.GetInterests()
	if err != nil {
		server.LogError(r, "/interests/get/", "database returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	jsonInterests, err := json.Marshal(interests)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		server.LogError(r, "/interests/get/", "Marshal returned error "+err.Error())
		server.error(w, errDef.MarshalError)
		return
	}

	// This is my valid case.
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonInterests)
	server.LogSuccess(r, "/interests/get/", "Interests was returned to user. Amount "+strconv.Itoa(len(interests)))
}

// HTTP HANDLER FOR DOMAIN /interests/get/ . IT HANDLES:
// RETURNING ARRAY OF EXISTING INTERESTS BY GET METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) HandlerInterestsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "GET" {
		server.interestsGet(w, r)
	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)
		server.Log(r, "/interests/get/", "client wants to know what methods are allowed")
	} else {
		// ALL OTHERS METHODS
		server.LogWarning(r, "/interests/get/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
