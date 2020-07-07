package myDatabase

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
)

func (conn *ConnDB) getUsersAll(w http.ResponseWriter, r *http.Request) {
	var (
		filter = r.URL.Query().Get("filter")
	)

	// all errors will be send to panic. This is recovery function
	defer func(w http.ResponseWriter) {
		if err := recover(); err != nil {
			fmt.Fprintf(w, "{\"error\":\"%s\"}", err)
		}
	}(w)

	users, err := conn.SearchUsersByOneFilter(filter)
	if err != nil {
		consoleLogError(r, "/users/", "Error: SearchUsersByOneFilter returned error " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("database request returned error")
	} else {
		jsonUsers, err := json.Marshal(users)
		if err != nil {
			consoleLogError(r, "/users/", "Error: Marshal returned error" + fmt.Sprintf("%s", err))
			w.WriteHeader(http.StatusInternalServerError) // 500
			panic("json convert error")
		} else {
			w.Write([]byte(jsonUsers))
			consoleLog(r, "/users/", "response was transmitted. Users amount " +  strconv.Itoa(len(users)))
		}
	}
}

func (conn *ConnDB) getUsersLogged(w http.ResponseWriter, r *http.Request) {
	// all errors will be send to panic. This is recovery function
	defer func(w http.ResponseWriter) {
		if err := recover(); err != nil {
			fmt.Fprintf(w, "{\"error\":\"%s\"}", err)
		}
	}(w)

	users := conn.session.GetLoggedUsersInfo()
	jsonUsers, err := json.Marshal(users)
	if err != nil {
		consoleLogError(r, "/users/", "Error: Marshal returned error" + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("json convert error")
	} else {
		w.Write([]byte(jsonUsers))
		consoleLog(r, "/users/", "response was transmitted. Users amount " +  strconv.Itoa(len(users)))
	}
}

func (conn *ConnDB) getUsers(w http.ResponseWriter, r *http.Request) {
	var (
		filter = r.URL.Query().Get("filter")
	)

	// all errors will be send to panic. This is recovery function
	defer func(w http.ResponseWriter) {
		if err := recover(); err != nil {
			fmt.Fprintf(w, "{\"error\":\"%s\"}", err)
		}
	}(w)

	consoleLog(r, "/users/", "request was recieved with filter \033[34m" + filter)

	if filter != "all" && filter != "logged" {
		consoleLogWarning(r, "/users/", "Warning: filter \033[34m" + filter + "\033[33m not exist")
		w.WriteHeader(http.StatusResetContent) // 205
		panic("no such filter - " + filter)
	}

	if filter == "all" { conn.getUsersAll(w, r) }
	if filter == "logged" { conn.getUsersLogged(w, r) }
}

func (conn *ConnDB) HttpHandlerUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "GET" {

		conn.getUsers(w, r)

	} else if r.Method == "OPTIONS" {
	// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)

		consoleLog(r, "/users/", "OPTIONS: client wants to know what methods are allowed")

	} else {
	// ALL OTHERS METHODS

		consoleLogWarning(r, "/users/", "Warning: wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}