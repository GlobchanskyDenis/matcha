package myDatabase

import (
	"fmt"
	"net/http"
	"encoding/json"
	"MatchaServer/handlers"
	. "MatchaServer/config"
)

// USER AUTHORISATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (conn *ConnDB) authUser(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, passwd, token, tokenWS, response string
		user User // config.User
		err error
		request map[string]interface{}
		isExist bool
	)

	// All errors will be send to panic. This is recovery function
	defer func(w http.ResponseWriter) {
		if err := recover(); err != nil {
			switch err.(type) {
			case error:
				fmt.Fprintf(w, `{"error":"` + err.(error).Error() + `"}`)
			case string:
				fmt.Fprintf(w, `{"error":"` + err.(string) + `"}`)
			}
		}
	}(w)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/auth/", "request decode error")
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("decode error")
	}

	arg, isExist := request["mail"]
	if !isExist {
		consoleLogWarning(r, "/auth/", "mail not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("mail not exist")
	}
	mail = arg.(string)

	arg, isExist = request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/auth/", "password not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("password not exist")
	}
	passwd = arg.(string)

	message = "request was recieved, mail: " + BLUE + mail + NO_COLOR + " password: hidden "
	consoleLog(r, "/auth/", message)

	// Simple validation
	if mail == "" || passwd == "" {
		consoleLogWarning(r, "/auth/", "mail or password is empty")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("mail or password is empty")
	}

	// Look for user in database
	user, err = conn.GetUserDataForAuth(mail, handlers.PasswdHash(passwd))
	if err != nil {
		consoleLogError(r, "/auth/", "GetUserDataForAuth returned error " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("wrong request in database")
	}
	if (user == User{}) {
		consoleLogWarning(r, "/auth/", "wrong mail or password")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		// w.WriteHeader(http.StatusNoContent) // 204 - With this status my json data will not add to response
		panic("wrong mail or password")
	}

	token, err = conn.session.AddUserToSession(user)
	if err != nil {
		consoleLogError(r, "/auth/", "SetNewUser returned error " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("Cannot authenticate this user")
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		consoleLogWarning(r, "/auth/", "Marshal returned error " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("cannot convert to json")
	}

	tokenWS, err = conn.session.CreateTokenWS(mail) //handlers.TokenWebSocketAuth(mail)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		consoleLogError(r, "/auth/", "cannot create web socket token - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("cannot create web socket token")
	}

	// This is my valid case. Response status will be set automaticly to 200.
	response = `{"x-auth-token":"` + token + `","ws-auth-token":"` + tokenWS + `",` + string(jsonUser[1:])
	fmt.Fprintf(w, response)
	consoleLogSuccess(r, "/auth/", "User " + BLUE + mail + NO_COLOR + " was authenticated successfully")
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
		consoleLog(r, "/auth/", "client wants to know what methods are allowed")
	} else {
	// ALL OTHERS METHODS
		consoleLogWarning(r, "/auth/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
