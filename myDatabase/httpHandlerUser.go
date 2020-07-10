package myDatabase

import (
	"fmt"
	"net/http"
	"encoding/json"
	"MatchaServer/handlers"
	"MatchaServer/session"
	"log"
	. "MatchaServer/config"
)

func consoleLog(r *http.Request, section string, message string) {
	log.Printf("%s %7s %7s %s\n", r.RemoteAddr, r.Method, section, message)
}

func consoleLogSuccess(r *http.Request, section string, message string) {
	log.Printf("%s %7s %7s %s\n", r.RemoteAddr, r.Method, section, GREEN_BG + "SUCCESS: " + NO_COLOR + message)
}

func consoleLogWarning(r *http.Request, section string, message string) {
	log.Printf("%s %7s %7s %s\n", r.RemoteAddr, r.Method, section, YELLOW_BG + "WARNING: " + NO_COLOR + message)
}

func consoleLogError(r *http.Request, section string, message string) {
	log.Printf("%s %7s %7s %s\n", r.RemoteAddr, r.Method, section, RED_BG + "ERROR: " + NO_COLOR + message)
}

// USER REGISTRATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (conn *ConnDB) regUser(w http.ResponseWriter, r *http.Request) {
	var (
		message, login, passwd, mail string
		err error
		request map[string]interface{}
		isExist bool
	)

	// all errors will be send to panic. This is recovery function
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
		consoleLogError(r, "/user/", "request decode error")
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("decode error")
	}
	
	arg, isExist := request["login"]
	if !isExist {
		consoleLogWarning(r, "/user/", "login not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("login not exist")
	}
	login = arg.(string)

	arg, isExist = request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/user/", "password not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("password not exist")
	}
	passwd = arg.(string)

	arg, isExist = request["mail"]
	if !isExist {
		consoleLogWarning(r, "/user/", "mail not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("mail not exist")
	}
	mail = arg.(string)

	message = "request was recieved, login: " + BLUE + login + NO_COLOR +
		" mail: " + BLUE + mail + NO_COLOR +
		" password: hidden"
	consoleLog(r, "/user/", message)

	// Simple validation
	if login == "" || mail == "" || passwd == "" {
		consoleLogWarning(r, "/user/", "login or password or mail is empty")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("login or password or mail is empty")
	}

	err = handlers.CheckLogin(login)
	if err != nil {
		consoleLogWarning(r, "/user/", "login - " + err.Error())
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		// CheckLogin is my own function, so I can not afraid of invalid runes in error
		panic("login error - " + err.Error())
	}

	err = handlers.CheckPasswd(passwd)
	if err != nil {
		consoleLogWarning(r, "/user/", "password - " + err.Error())
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		// CheckLogin is my own function, so I can not afraid of invalid runes in error
		panic("password error - " + err.Error())
	}

	err = handlers.CheckMail(mail)
	if err != nil {
		consoleLogWarning(r, "/user/", "mail - " + err.Error())
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		// CheckLogin is my own function, so I can not afraid of invalid runes in error
		panic("mail error - " + err.Error())
	}

	isUserExists, err := conn.IsUserExists(login)
	if err != nil {
		consoleLogError(r, "/user/", "IsUserExists returned error " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("wrong request in database")
	}
	if isUserExists {
		consoleLogWarning(r, "/user/", "user " + BLUE + login + NO_COLOR + " alredy exists")
		w.WriteHeader(http.StatusAlreadyReported) // 208
		panic("user " + login + " already exists")
	}

	err = conn.SetNewUser(login, handlers.PasswdHash(passwd), mail)
	if err != nil {
		consoleLogError(r, "/user/", "SetNewUser returned error " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("Cannot register this user")
	}
	w.WriteHeader(201)
	consoleLogSuccess(r, "/user/", "user " + BLUE + login + NO_COLOR + " was created successfully. No response body")
}

// USER UPDATE BY PATCH METHOD. REQUEST AND RESPONSE DATA IS JSON.
// REQUEST SHOULD HAVE 'x-auth-token' HEADER
func (conn *ConnDB) updateUser(w http.ResponseWriter, r *http.Request) {
	var (
		message, newToken string
		err error
		request map[string]interface{}
		update = map[string]string{}
		isExist, usefullFieldsExists bool
		token = r.Header.Get("x-auth-token")
		sessionUser session.SessionItem
		user	UserStruct
	)

	// all errors will be send to panic. This is recovery function
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
		consoleLogError(r, "/user/", "request decode error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("decode error")
	}

	message = "request for UPDATE was recieved: "

	arg, isExist := request["login"]
	if isExist {
		usefullFieldsExists = true
		update["login"] =  arg.(string)
		message += " login=" + BLUE + update["login"] + NO_COLOR
	}

	arg, isExist = request["passwd"]
	if isExist {
		usefullFieldsExists = true
		update["passwd"] =  arg.(string)
		message += " password=hidden"
	}

	arg, isExist = request["mail"]
	if isExist {
		usefullFieldsExists = true
		update["mail"] =  arg.(string)
		message += " mail=" + BLUE + update["mail"] + NO_COLOR
	}

	consoleLog(r, "/user/", message)

	if !usefullFieldsExists {
		consoleLogWarning(r, "/user/", "no usefull fields found")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("no usefull fields")
	}

	if token == "" {
		consoleLogWarning(r, "/user/", "token is empty")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("token is empty")
	}

	sessionUser, err = conn.session.FindUserByToken(token)
	if err != nil {
		consoleLogWarning(r, "/user/", "FindUserByToken returned error - " + err.Error())
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic(err)
	}

	_, isExist = update["login"]
	if isExist {
		sessionUser.UserInfo.Login = update["login"]
		err = handlers.CheckLogin(update["login"])
		if err != nil {
			consoleLogWarning(r, "/user/", "login - " + err.Error())
			w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
			// CheckLogin is my own function, so I can not afraid of invalid runes in error
			panic("login error - " + err.Error())
		}
	}

	_, isExist = update["passwd"]
	if isExist {
		sessionUser.UserInfo.Passwd = handlers.PasswdHash(update["passwd"])
		err = handlers.CheckPasswd(update["passwd"])
		if err != nil {
			consoleLogWarning(r, "/user/", "password - " + err.Error())
			w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
			// CheckLogin is my own function, so I can not afraid of invalid runes in error
			panic("password error - " + err.Error())
		}
	}

	_, isExist = update["mail"]
	if isExist {
		sessionUser.UserInfo.Mail = update["mail"]
		err = handlers.CheckMail(update["mail"])
		if err != nil {
			consoleLogWarning(r, "/user/", "mail - " + err.Error())
			w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
			// CheckLogin is my own function, so I can not afraid of invalid runes in error
			panic("mail error - " + err.Error())
		}
	}

	user, err = conn.GetUserById(sessionUser.UserInfo.Id)
	if err != nil {
		consoleLogError(r, "/user/", "GetUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("database request returned error")
	}

	_, isExist = update["login"]
	if isExist {
		user.Login = update["login"]
	}

	_, isExist = update["passwd"]
	if isExist {
		user.Passwd = handlers.PasswdHash(update["passwd"])
	}

	_, isExist = update["mail"]
	if isExist {
		user.Mail = update["mail"]
	}

	err = conn.session.UpdateSessionUser(token, sessionUser.UserInfo)
	if err != nil {
		consoleLogError(r, "/user/", "UpdateSessionUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("token decode error")
	}

	err = conn.UpdateUser(user)
	if err != nil {
		consoleLogError(r, "/user/", "UpdateUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("database request returned error")
	}

	_, isExist = update["login"]
	if isExist {
		newToken, err = handlers.TokenEncode(update["login"])
		if err != nil {
			consoleLogError(r, "/user/", "Token encode returned error - " + err.Error())
			w.WriteHeader(http.StatusInternalServerError) // 500
			panic("error in token encoding")
		}
		fmt.Fprintf(w, `{"x-auth-token":"` + newToken + `"}`)
		consoleLogSuccess(r, "/user/", "user " + BLUE + user.Login + NO_COLOR + " was updated successfully. x-auth-token in response body")
	} else {
		consoleLogSuccess(r, "/user/", "user " + BLUE + user.Login + NO_COLOR + " was updated successfully. No response body")
	}
}

// USER REMOVE BY DELETE METHOD. NO REQUEST DATA. RESPONSE DATA IS JSON ONLY IN CASE OF ERROR.
// REQUEST SHOULD HAVE 'x-auth-token' HEADER
func (conn *ConnDB) deleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		message string
		err error
		token = r.Header.Get("x-auth-token")
		sessionUser session.SessionItem
		request map[string]interface{}
		passwd string
	)

	// all errors will be send to panic. This is recovery function
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

	message = "request for DELETE was recieved: token=" + BLUE + token + NO_COLOR
	consoleLog(r, "/user/", message)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/", "request decode error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("decode error")
	}

	arg, isExist := request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/user/", "password not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("password not exist")
	}
	passwd = handlers.PasswdHash(arg.(string))

	sessionUser, err = conn.session.FindUserByToken(token)
	if err != nil {
		consoleLogWarning(r, "/user/", "FindUserByToken returned error - " + err.Error())
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic(err)
	}

	if passwd != sessionUser.UserInfo.Passwd {
		consoleLogWarning(r, "/user/", "password is incorrect " + BLUE + passwd + " " + sessionUser.UserInfo.Passwd + NO_COLOR)
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("wrong password")
	}

	conn.session.DeleteUserSessionByLogin(sessionUser.UserInfo.Login)

	err = conn.DeleteUser(sessionUser.UserInfo.Id)
	if err != nil {
		consoleLogError(r, "/user/", "DeleteUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("database request returned error")
	}
	consoleLogSuccess(r, "/user/", "user " + BLUE + sessionUser.UserInfo.Login + NO_COLOR + " was removed successfully. No response body")
}

// HTTP HANDLER FOR DOMAIN /user/ . IT HANDLES:
// REGISTRATE USER BY POST METHOD
// UPDATE USER BY PATCH METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (conn *ConnDB) HttpHandlerUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST,PATCH,OPTIONS,DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "POST" {

		conn.regUser(w, r)

	} else if r.Method == "OPTIONS" {
	// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)

		consoleLog(r, "/user/", "client wants to know what methods are allowed")

	} else if r.Method == "PATCH" {

		conn.updateUser(w, r)

	} else if r.Method == "DELETE" {

		conn.deleteUser(w, r)

	} else {
	// ALL OTHERS METHODS

		consoleLogWarning(r, "/user/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}
