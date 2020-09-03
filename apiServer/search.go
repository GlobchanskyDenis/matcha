package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errDef"
	"encoding/json"
	"net/http"
	"strconv"
)

func (server *Server) searchAll(w http.ResponseWriter, r *http.Request) {
	var filter = r.URL.Query().Get("filter")

	users, err := server.Db.SearchUsersByOneFilter(filter)
	if err != nil {
		server.LogError(r, "SearchUsersByOneFilter returned error "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		server.LogError(r, "Marshal returned error"+err.Error())
		server.error(w, errDef.MarshalError)
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte(jsonUsers))
	server.LogSuccess(r, "array of "+BLUE+"all"+NO_COLOR+
		" users was transmitted. Users amount "+strconv.Itoa(len(users)))
}

func (server *Server) searchLogged(w http.ResponseWriter, r *http.Request) {

	users, err := server.Db.GetLoggedUsers(server.session.GetLoggedUsersUidSlice())
	if err != nil {
		server.LogError(r, "GetLoggedUsers returned error"+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		server.LogError(r, "Marshal returned error"+err.Error())
		server.error(w, errDef.MarshalError)
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonUsers)
	server.LogSuccess(r, "array of "+BLUE+"logged"+NO_COLOR+
		" users was transmitted. Users amount "+strconv.Itoa(len(users)))
}

func (server *Server) search(w http.ResponseWriter, r *http.Request) {
	var filter = r.URL.Query().Get("filter")

	server.Log(r, "request was recieved with filter "+BLUE+filter+NO_COLOR)

	if filter != "all" && filter != "logged" {
		server.LogWarning(r, "filter "+BLUE+filter+NO_COLOR+" not exist")
		errDef.InvalidArgument.WithArguments("Значение поля filter недопустимо", "filter field has wrong value")
		return
	}

	if filter == "all" {
		server.searchAll(w, r)
	}
	if filter == "logged" {
		server.searchLogged(w, r)
	}
}

func (server *Server) HandlerSearch(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "GET" {

		server.search(w, r)

	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)

		server.Log(r, "client wants to know what methods are allowed")

	} else {
		// ALL OTHERS METHODS

		server.LogWarning(r, "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}
