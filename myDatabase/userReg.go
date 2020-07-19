package myDatabase

import (
	. "MatchaServer/config"
	"MatchaServer/handlers"
	"encoding/json"
	"fmt"
	"net/http"
)

// USER REGISTRATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (conn *ConnDB) userReg(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, passwd, token string
		err                   error
		request               map[string]interface{}
		isExist               bool
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/reg/", "request json decode failed - "+err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"json decode failed"+`"}`)
		return
	}

	arg, isExist := request["mail"]
	if !isExist {
		consoleLogWarning(r, "/user/reg/", "mail not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"mail not exist"+`"}`)
		return
	}
	mail = arg.(string)

	arg, isExist = request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/user/reg/", "password not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"password not exist"+`"}`)
		return
	}
	passwd = arg.(string)

	message = "request was recieved, mail: " + BLUE + mail + NO_COLOR +
		" password: hidden"
	consoleLog(r, "/user/reg/", message)

	if mail == "" || passwd == "" {
		consoleLogWarning(r, "/user/reg/", "mail or password is empty")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"mail or password is empty"+`"}`)
		return
	}

	err = handlers.CheckMail(mail)
	if err != nil {
		consoleLogWarning(r, "/user/reg/", "mail - "+err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		// CheckMail is my own function, so I can not afraid of invalid runes in error
		fmt.Fprintf(w, `{"error":"`+"mail error - "+err.Error()+`"}`)
		return
	}

	err = handlers.CheckPasswd(passwd)
	if err != nil {
		consoleLogWarning(r, "/user/reg/", "password - "+err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		// CheckPasswd is my own function, so I can not afraid of invalid runes in error
		fmt.Fprintf(w, `{"error":"`+"password error - "+err.Error()+`"}`)
		return
	}

	isUserExists, err := conn.IsUserExists(mail)
	if err != nil {
		consoleLogError(r, "/user/reg/", "IsUserExists returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"`+"database request returned error"+`"}`)
		return
	}
	if isUserExists {
		consoleLogWarning(r, "/user/reg/", "user "+BLUE+mail+NO_COLOR+" alredy exists")
		w.WriteHeader(http.StatusNotAcceptable) // 406
		fmt.Fprintf(w, `{"error":"`+"user "+mail+" already exists"+`"}`)
		return
	}

	err = conn.SetNewUser(mail, handlers.PasswdHash(passwd))
	if err != nil {
		consoleLogError(r, "/user/reg/", "SetNewUser returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"`+"Cannot register this user"+`"}`)
		return
	}

	token, err = handlers.TokenMailEncode(mail)
	if err != nil {
		consoleLogError(r, "/user/reg/", "TokenMailEncode returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"`+"Cannot create token for this user"+`"}`)
		return
	}

	w.WriteHeader(201)
	consoleLogSuccess(r, "/user/reg/", "user "+BLUE+mail+NO_COLOR+" was created successfully. No response body")

	go func(mail string, xRegToken string, r *http.Request) {
		err := handlers.SendMail(mail, xRegToken)
		if err != nil {
			consoleLogError(r, "/user/reg/", "SendMail returned error "+err.Error())
		} else {
			consoleLogSuccess(r, "/user/reg/", "Confirm mail for user "+BLUE+mail+NO_COLOR+" was send successfully")
		}
	}(mail, token, r)
}

// HTTP HANDLER FOR DOMAIN /user/reg
// REGISTRATE USER BY POST METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (conn *ConnDB) HttpHandlerUserReg(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST,PATCH,OPTIONS,DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "POST" {

		conn.userReg(w, r)

	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)

		consoleLog(r, "/user/reg/", "client wants to know what methods are allowed")

	} else {
		// ALL OTHERS METHODS

		consoleLogWarning(r, "/user/reg/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}