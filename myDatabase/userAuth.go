package myDatabase

import (
	. "MatchaServer/config"
	"MatchaServer/handlers"
	"encoding/json"
	"fmt"
	"net/http"
)

// USER AUTHORISATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (conn *ConnDB) userAuth(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, passwd, token, tokenWS, response string
		user                                            User
		err                                             error
		request                                         map[string]interface{}
		isExist                                         bool
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/auth/", "request decode error")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"decode error"+`"}`)
		return
	}

	arg, isExist := request["mail"]
	if !isExist {
		consoleLogWarning(r, "/auth/", "mail not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"mail not exist"+`"}`)
		return
	}
	mail = arg.(string)

	arg, isExist = request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/auth/", "password not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"password not exist"+`"}`)
		return
	}
	passwd = arg.(string)

	message = "request was recieved, mail: " + BLUE + mail + NO_COLOR + " password: hidden "
	consoleLog(r, "/auth/", message)

	// Simple validation
	if mail == "" || passwd == "" {
		consoleLogWarning(r, "/auth/", "mail or password is empty")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"mail or password is empty"+`"}`)
		return
	}

	user, err = conn.GetUserDataForAuth(mail, handlers.PasswdHash(passwd))
	if err != nil {
		consoleLogError(r, "/auth/", "GetUserDataForAuth returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"`+"database request failed"+`"}`)
		return
	}

	if (user == User{}) {
		consoleLogWarning(r, "/auth/", "wrong mail or password")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"wrong mail or password"+`"}`)
		return
	}

	if user.AccType == "not confirmed" {
		consoleLogWarning(r, "/auth/", "user " + BLUE + user.Mail + NO_COLOR + " should confirm its email")
		w.WriteHeader(http.StatusAccepted) // 202
		fmt.Fprintf(w, `{"error":"`+"confirm email first"+`"}`)
		return
	}

	token, err = conn.session.AddUserToSession(user.Uid)
	if err != nil {
		consoleLogError(r, "/auth/", "SetNewUser returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"`+"Cannot authenticate this user"+`"}`)
		return
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		consoleLogWarning(r, "/auth/", "Marshal returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"`+"cannot convert to json"+`"}`)
		return
	}

	tokenWS, err = conn.session.CreateTokenWS(user.Uid) //handlers.TokenWebSocketAuth(mail)
	if err != nil {
		// удалить пользователя из сессии (потом - когда решится вопрос со множественностью веб сокетов)
		consoleLogError(r, "/auth/", "cannot create web socket token - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"` + "cannot create web socket token" + `"}`)
		return
	}

	// This is my valid case. Response status will be set automaticly to 200.
	w.WriteHeader(http.StatusOK) // 200
	response = `{"x-auth-token":"` + token + `","ws-auth-token":"` + tokenWS + `",` + string(jsonUser[1:])
	fmt.Fprintf(w, response)
	consoleLogSuccess(r, "/auth/", "User " + BLUE + mail + NO_COLOR + " was authenticated successfully")
}

// HTTP HANDLER FOR DOMAIN /auth/ . IT HANDLES:
// AUTHENTICATE USER BY POST METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (conn *ConnDB) HttpHandlerUserAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "POST" {
		conn.userAuth(w, r)
	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)
		consoleLog(r, "/auth/", "client wants to know what methods are allowed")
	} else {
		// ALL OTHERS METHODS
		consoleLogWarning(r, "/auth/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
