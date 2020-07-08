package myDatabase

import (
	"fmt"
	"net/http"
	"encoding/json"
	"MatchaServer/handlers"
	"MatchaServer/session"
	"log"
)

func consoleLog(r *http.Request, section string, message string) {
	log.Printf("%s %7s %7s %s\n", r.RemoteAddr, r.Method, section, "\033[32m"+message+"\033[m")
}

func consoleLogWarning(r *http.Request, section string, message string) {
	log.Printf("%s %7s %7s %s\n", r.RemoteAddr, r.Method, section, "\033[33m"+message+"\033[m")
}

func consoleLogError(r *http.Request, section string, message string) {
	log.Printf("%s %7s %7s %s\n", r.RemoteAddr, r.Method, section, "\033[31m"+message+"\033[m")
}


// USER REGISTRATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (conn *ConnDB) regUser(w http.ResponseWriter, r *http.Request) {
	var (
		message, login, passwd, mail, phone string
		err error
		request map[string]interface{}
		isExist bool
	)

	// all errors will be send to panic. This is recovery function
	defer func(w http.ResponseWriter) {
		if err := recover(); err != nil {
			fmt.Fprintf(w, "{\"error\":\"%s\"}", err)
		}
	}(w)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/", "Error: request decode error")
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("decode error")
	}
	
	arg, isExist := request["login"]
	if !isExist {
		consoleLogWarning(r, "/user/", "Warning: login not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("login not exist")
	}
	login = arg.(string)

	arg, isExist = request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/user/", "Warning: password not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("password not exist")
	}
	passwd = arg.(string)

	arg, isExist = request["mail"]
	if !isExist {
		consoleLogWarning(r, "/user/", "Warning: mail not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("mail not exist")
	}
	mail = arg.(string)

	arg, isExist = request["phone"]
	if !isExist {
		consoleLogWarning(r, "/user/", "Warning: phone number not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("phone number not exist")
	}
	phone = arg.(string)

	message = "request was recieved, login: \033[34m" + login +
		"\033[32m mail: \033[34m" + mail +
		"\033[32m phone: \033[34m" + phone +
		"\033[32m password: hidden"
	consoleLog(r, "/user/", message)

	// Simple validation
	if login == "" || mail == "" || passwd == "" || phone == "" {
		consoleLogWarning(r, "/user/", "Warning: login or password or mail or phone is empty")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("login or password or mail or phone is empty")
	}

	err = handlers.CheckLogin(login)
	if err != nil {
		consoleLogWarning(r, "/user/", "Warning: login - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		// CheckLogin is my own function, so I can not afraid of invalid runes in error
		panic( fmt.Errorf( "login error - %s", err ) )
	}

	err = handlers.CheckPasswd(passwd)
	if err != nil {
		consoleLogWarning(r, "/user/", "Warning: password - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		// CheckLogin is my own function, so I can not afraid of invalid runes in error
		panic( fmt.Errorf( "password error - %s", err ) )
	}

	err = handlers.CheckMail(mail)
	if err != nil {
		consoleLogWarning(r, "/user/", "Warning: mail - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		// CheckLogin is my own function, so I can not afraid of invalid runes in error
		panic( fmt.Errorf( "mail error - %s", err ) )
	}

	err = handlers.CheckPhone(phone)
	if err != nil {
		consoleLogWarning(r, "/user/", "Warning: phone - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		// CheckLogin is my own function, so I can not afraid of invalid runes in error
		panic( fmt.Errorf( "phone error - %s", err ) )
	}

	isUserExists, err := conn.IsUserExists(login)
	if err != nil {
		consoleLogError(r, "/user/", "IsUserExists returned error " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("wrong request in database")
	}
	if isUserExists {
		consoleLogWarning(r, "/user/", "Warning: user \033[34m" + login + "\033[33m alredy exists")
		w.WriteHeader(http.StatusAlreadyReported) // 208
		panic("user " + login + " already exists")
	}

	err = conn.SetNewUser(login, handlers.PasswdHash(passwd), mail, phone)
	if err != nil {
		consoleLogError(r, "/user/", "SetNewUser returned error " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("Cannot register this user")
	} else {
		consoleLog(r, "/user/", "Request handled fine. Response will be with empty error field")
		// here I send no response in valid case - just response status 201
		w.WriteHeader(201)
	}
}

// USER UPDATE BY PATCH METHOD. REQUEST AND RESPONSE DATA IS JSON.
// REQUEST SHOULD HAVE 'x-auth-token' HEADER
func (conn *ConnDB) updateUser(w http.ResponseWriter, r *http.Request) {
	var (
		message string
		err error
		request map[string]interface{}
		update = map[string]string{}
		isExist bool
		token = r.Header.Get("x-auth-token")
		sessionUser session.SessionItem
		user	UserStruct
		// wg = &sync.WaitGroup{}
	)

	// all errors will be send to panic. This is recovery function
	defer func(w http.ResponseWriter) {
		if err := recover(); err != nil {
			fmt.Fprintf(w, "{\"error\":\"%s\"}", err)
		}
	}(w)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/", "Error: request decode error - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("decode error")
	}

	message = "request for UPDATE was recieved: token=\033[34m" + token + "\033[32m"

	arg, isExist := request["login"]
	if isExist {
		update["login"] =  arg.(string)
		message += " login=\033[34m" + update["login"] + "\033[32m"
	}

	arg, isExist = request["passwd"]
	if isExist {
		update["passwd"] =  arg.(string)
		message += " password=hidden"
	}

	arg, isExist = request["mail"]
	if isExist {
		update["mail"] =  arg.(string)
		message += " mail=\033[34m" + update["mail"] + "\033[32m"
	}

	arg, isExist = request["phone"]
	if isExist {
		update["phone"] =  arg.(string)
		message += " phone=\033[34m" + update["phone"] + "\033[32m"
	}

	consoleLog(r, "/user/", message)

	sessionUser, err = conn.session.FindUserByToken(token)
	if err != nil {
		consoleLogWarning(r, "/user/", "Warning: FindUserByToken returned error - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic(err)
	}

	if (sessionUser.UserInfo.Login == "" && sessionUser.UserInfo.Id == 0) {
		consoleLogWarning(r, "/user/", "Warning: FindUserByToken returned empty struct - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic(err)
	}

	_, isExist = update["login"]
	if isExist {
		sessionUser.UserInfo.Login = update["login"]
		err = handlers.CheckLogin(update["login"])
		if err != nil {
			consoleLogWarning(r, "/user/", "Warning: login - " + fmt.Sprintf("%s", err))
			w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
			// CheckLogin is my own function, so I can not afraid of invalid runes in error
			panic( fmt.Errorf( "login error - %s", err ) )
		}
	}

	_, isExist = update["passwd"]
	if isExist {
		sessionUser.UserInfo.Passwd = handlers.PasswdHash(update["passwd"])
		err = handlers.CheckPasswd(update["passwd"])
		if err != nil {
			consoleLogWarning(r, "/user/", "Warning: password - " + fmt.Sprintf("%s", err))
			w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
			// CheckLogin is my own function, so I can not afraid of invalid runes in error
			panic( fmt.Errorf( "password error - %s", err ) )
		}
	}

	_, isExist = update["mail"]
	if isExist {
		sessionUser.UserInfo.Mail = update["mail"]
		err = handlers.CheckMail(update["mail"])
		if err != nil {
			consoleLogWarning(r, "/user/", "Warning: mail - " + fmt.Sprintf("%s", err))
			w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
			// CheckLogin is my own function, so I can not afraid of invalid runes in error
			panic( fmt.Errorf( "mail error - %s", err ) )
		}
	}

	_, isExist = update["phone"]
	if isExist {
		sessionUser.UserInfo.Phone = update["phone"]
		err = handlers.CheckPhone(update["phone"])
		if err != nil {
			consoleLogWarning(r, "/user/", "Warning: phone number - " + fmt.Sprintf("%s", err))
			w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
			// CheckLogin is my own function, so I can not afraid of invalid runes in error
			panic( fmt.Errorf( "phone number error - %s", err ) )
		}
	}

	if token == "" {
		consoleLogWarning(r, "/user/", "Warning: token is empty")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("token is empty")
	}

	user, err = conn.GetUserById(sessionUser.UserInfo.Id)
	if err != nil {
		consoleLogError(r, "/user/", "Error: GetUser returned error - " + fmt.Sprintf("%s", err))
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

	_, isExist = update["phone"]
	if isExist {
		user.Phone = update["phone"]
	}

	err = conn.session.UpdateSessionUser(token, sessionUser.UserInfo)
	if err != nil {
		consoleLogError(r, "/user/", "Error: UpdateSessionUser returned error - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("token decode error")
	}

	err = conn.UpdateUser(user)
	if err != nil {
		consoleLogError(r, "/user/", "Error: UpdateUser returned error - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("database request returned error")
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
			fmt.Fprintf(w, "{\"error\":\"%s\"}", err)
		}
	}(w)

	message = "request for DELETE was recieved: token=\033[34m" + token
	consoleLog(r, "/user/", message)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/", "Error: request decode error - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("decode error")
	}

	arg, isExist := request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/user/", "Warning: password not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("password not exist")
	}
	passwd = handlers.PasswdHash(arg.(string))
	// passwd = arg.(string)

	sessionUser, err = conn.session.FindUserByToken(token)
	if err != nil {
		consoleLogWarning(r, "/user/", "Warning: FindUserByToken returned error - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic(err)
	}

	if passwd != sessionUser.UserInfo.Passwd {
		consoleLogWarning(r, "/user/", "Warning: password is incorrect "  + passwd + " " + sessionUser.UserInfo.Passwd)
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("wrong password")
	}

	message = "removing user \033[34m" + sessionUser.UserInfo.Login + "\033[32m token=\033[34m" + token
	consoleLog(r, "/user/", message)

	conn.session.DeleteUserSessionByLogin(sessionUser.UserInfo.Login)

	err = conn.DeleteUser(sessionUser.UserInfo.Id)
	if err != nil {
		consoleLogError(r, "/user/", "Error: DeleteUser returned error - " + fmt.Sprintf("%s", err))
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("database request returned error")
	}

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

		consoleLog(r, "/user/", "OPTIONS: client wants to know what methods are allowed")

	} else if r.Method == "PATCH" {

		conn.updateUser(w, r)

	} else if r.Method == "DELETE" {

		conn.deleteUser(w, r)

	} else {
	// ALL OTHERS METHODS

		consoleLogWarning(r, "/user/", "Warning: wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}

