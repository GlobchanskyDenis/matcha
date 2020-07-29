package apiServer

import (
	. "MatchaServer/config"
	"encoding/json"
	"net/http"
	"strconv"
)

func (server *Server) searchAll(w http.ResponseWriter, r *http.Request) {
	var filter = r.URL.Query().Get("filter")

	users, err := server.Db.SearchUsersByOneFilter(filter)
	if err != nil {
		consoleLogError(r, "/users/", "SearchUsersByOneFilter returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "database request returned error" + `"}`))
		return
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		consoleLogError(r, "/users/", "Marshal returned error"+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "json convert error" + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte(jsonUsers))
	consoleLogSuccess(r, "/users/", "array of "+BLUE+"all"+NO_COLOR+
		" users was transmitted. Users amount "+strconv.Itoa(len(users)))
}

func (server *Server) searchLogged(w http.ResponseWriter, r *http.Request) {

	users, err := server.Db.GetLoggedUsers(server.session.GetLoggedUsersUidSlice())
	if err != nil {
		consoleLogError(r, "/users/", "GetLoggedUsers returned error"+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "database request failed" + `"}`))
		return
	}
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		consoleLogError(r, "/users/", "Marshal returned error"+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "json convert error" + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte(jsonUsers))
	consoleLogSuccess(r, "/users/", "array of "+BLUE+"logged"+NO_COLOR+
		" users was transmitted. Users amount "+strconv.Itoa(len(users)))
}

func (server *Server) search(w http.ResponseWriter, r *http.Request) {
	var filter = r.URL.Query().Get("filter")

	consoleLog(r, "/users/", "request was recieved with filter "+BLUE+filter+NO_COLOR)

	if filter != "all" && filter != "logged" {
		consoleLogWarning(r, "/users/", "filter "+BLUE+filter+NO_COLOR+" not exist")
		w.WriteHeader(http.StatusResetContent) // 205
		w.Write([]byte(`{"error":"` + "no such filter - " + filter + `"}`))
		return
	}

	if filter == "all" {
		server.searchAll(w, r)
	}
	if filter == "logged" {
		server.searchLogged(w, r)
	}
}

func (server *Server) HttpHandlerSearch(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "GET" {

		server.search(w, r)

	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)

		consoleLog(r, "/users/", "client wants to know what methods are allowed")

	} else {
		// ALL OTHERS METHODS

		consoleLogWarning(r, "/users/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}
