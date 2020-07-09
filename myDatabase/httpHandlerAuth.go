package myDatabase

import (
	"fmt"
	"net/http"
	"encoding/json"
	"MatchaServer/handlers"
)

// USER AUTHORISATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (conn *ConnDB) authUser(w http.ResponseWriter, r *http.Request) {
	var (
		message, login, passwd, token, tokenWS, response string
		user UserStruct
		err error
		request map[string]interface{}
		isExist bool
	)

	// All errors will be send to panic. This is recovery function
	defer func(w http.ResponseWriter) {
		if err := recover(); err != nil {
			fmt.Fprintf(w, `{"error":"%s"}`, fmt.Sprintf("%s", err))
		}
	}(w)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/auth/", "Error: request decode error")
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("decode error")
	}

	arg, isExist := request["login"]
	if !isExist {
		consoleLogWarning(r, "/auth/", "Warning: login not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("login not exist")
	}

	login = arg.(string)
	arg, isExist = request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/auth/", "Warning: password not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("password not exist")
	}

	passwd = arg.(string)
	message = "request was recieved, login: \033[34m" + login + "\033[32m password: hidden "
	consoleLog(r, "/auth/", message)

	// Simple validation
	if login == "" || passwd == "" {
		consoleLogWarning(r, "/auth/", "Warning: login or password is empty")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("login or password is empty")
	}

	// Look for user in database
	user, err = conn.GetUserDataForAuth(login, handlers.PasswdHash(passwd))
	if err != nil {
		consoleLogError(r, "/auth/", "GetUserDataForAuth returned error " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("wrong request in database")
	}

	if (user == UserStruct{}) {
		consoleLogWarning(r, "/auth/", "Warning: wrong login or password")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		// w.WriteHeader(http.StatusNoContent) // 204 - With this status my json data will not add to response
		panic("wrong login or password")
	} else {
		token, err = conn.session.AddUserToSession(user.Id, user.Login, user.Passwd, user.Mail)
		if err != nil {
			consoleLogError(r, "/auth/", "SetNewUser returned error " + fmt.Sprintf("%s", err))
			w.WriteHeader(http.StatusInternalServerError) // 500
			panic("Cannot authenticate this user")
		}
		consoleLog(r, "/auth/", "User \033[34m" + login + "\033[32m was found successfully. Token is \033[34m" + token)
		jsonUser, err := json.Marshal(user)
		if err != nil {
			consoleLogWarning(r, "/auth/", "Marshal returned error " + fmt.Sprintf("%s", err))
			w.WriteHeader(http.StatusInternalServerError) // 500
			panic("cannot convert to json")
		}
		tokenWS, err = conn.session.CreateTokenWS(login) //handlers.TokenWebSocketAuth(login)
		if err != nil {
			
		}
		// This is my valid case. Response status will be set automaticly to 200.
		response = `{"x-auth-token":"` + token + `","ws-auth-token":"`+tokenWS+`",` + string(jsonUser[1:])
		fmt.Fprintf(w, response)
	}
}

// HTTP HANDLER FOR DOMAIN /auth/ . IT HANDLES:
// AUTHENTICATE USER BY POST METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (conn *ConnDB) HttpHandlerAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "POST" {
		conn.authUser(w, r)
	} else if r.Method == "OPTIONS" {
	// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)
		consoleLog(r, "/auth/", "OPTIONS: client wants to know what methods are allowed")
	} else {
	// ALL OTHERS METHODS
		consoleLogWarning(r, "/auth/", "Warning: wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
