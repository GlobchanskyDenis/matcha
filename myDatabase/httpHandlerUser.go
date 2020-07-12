package myDatabase

import (
	"fmt"
	"net/http"
	"encoding/json"
	"MatchaServer/handlers"
	"MatchaServer/session"
	"log"
	. "MatchaServer/config"
	"strconv"
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

func fillUserStruct(request map[string]interface, user User) (User, error) {
	var usefullFieldsExists, ok, isExist bool
	var ok bool
	var message string

	message = "request for UPDATE was recieved: "

	arg, isExist := request["mail"]
	if isExist {
		usefullFieldsExists = true
		user.Mail, ok = arg.string
		if !ok {
			return user, fmt.Errorf("wrong type of param")
		}
		err = handlers.CheckMail(user.Mail)
		if err != nil {
			return user, err
		}
		message += " mail=" + BLUE + arg.string + NO_COLOR
	}
	arg, isExist = request["passwd"]
	if isExist {
		usefullFieldsExists = true
		user.Passwd, ok = arg.string
		if !ok {
			return user, fmt.Errorf("wrong type of param")
		}
		err = handlers.CheckPasswd(user.Passwd)
		if err != nil {
			return user, err
		}
		message += " password=hidden"
	}
	arg, isExist = request["fname"]
	if isExist {
		usefullFieldsExists = true
		user.Fname, ok = arg.string
		if !ok {
			return user, fmt.Errorf("wrong type of param")
		}
		err = handlers.CheckName(user.Fname)
		if err != nil {
			return user, err
		}
		message += " fname=" + BLUE + arg.string + NO_COLOR
	}
	arg, isExist = request["lname"]
	if isExist {
		usefullFieldsExists = true
		user.Lname, ok = arg.string
		if !ok {
			return user, fmt.Errorf("wrong type of param")
		}
		err = handlers.CheckName(user.Lname)
		if err != nil {
			return user, err
		}
		message += " lname=" + BLUE + arg.string + NO_COLOR
	}
	arg, isExist = request["age"]
	if isExist {
		usefullFieldsExists = true
		user.Age, ok = arg.int
		if !ok {
			return user, fmt.Errorf("wrong type of param")
		}
		if user.Age > 80 || user.Age < 14 {
			return user, fmt.Errorf("this age is forbidden")
		}
		message += " age=" + BLUE + strconv.Itoa(arg.int) + NO_COLOR
	}
	arg, isExist = request["gender"]
	if isExist {
		usefullFieldsExists = true
		user.Gender, ok = arg.string
		if !ok {
			return user, fmt.Errorf("wrong type of param")
		}
		err = handlers.CheckGender(user.Gender)
		if err != nil {
			return user, err
		}
		message += " gender=" + BLUE + arg.string + NO_COLOR
	}
	arg, isExist = request["orientation"]
	if isExist {
		usefullFieldsExists = true
		user.Orientation, ok = arg.string
		if !ok {
			return user, fmt.Errorf("wrong type of param")
		}
		err = handlers.CheckOrientation(user.Orientation)
		if err != nil {
			return user, err
		}
		message += " orientation=" + BLUE + arg.string + NO_COLOR
	}
	arg, isExist = request["biography"]
	if isExist {
		usefullFieldsExists = true
		user.Biography, ok = arg.string
		if !ok {
			return user, fmt.Errorf("wrong type of param")
		}
		err = handlers.CheckBiography(user.Biography)
		if err != nil {
			return user, err
		}
		message += " biography=" + BLUE + arg.string + NO_COLOR
	}
	arg, isExist = request["avaPhotoID"]
	if isExist {
		usefullFieldsExists = true
		user.AvaPhotoID, ok = arg.int
		if !ok {
			return user, fmt.Errorf("wrong type of param")
		}
		if user.AvaPhotoID < 0 {
			return user, fmt.Errorf("this age is forbidden")
		}
		message += " avaPhotoID=" + BLUE + strconv.Itoa(arg.int) + NO_COLOR
	}

	if !usefullFieldsExists {
		return user, fmt.Errorf("no usefull fields found")
	}
	consoleLog(r, "/user/", message)
	return user, nil
}

// USER REGISTRATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (conn *ConnDB) regUser(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, passwd string
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
	
	arg, isExist = request["mail"]
	if !isExist {
		consoleLogWarning(r, "/user/", "mail not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("mail not exist")
	}
	mail = arg.(string)

	arg, isExist = request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/user/", "password not exist")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("password not exist")
	}
	passwd = arg.(string)

	message = "request was recieved, mail: " + BLUE + mail + NO_COLOR +
		" password: hidden"
	consoleLog(r, "/user/", message)

	// Simple validation
	if mail == "" || passwd == "" {
		consoleLogWarning(r, "/user/", "mail or password is empty")
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("mail or password is empty")
	}

	err = handlers.CheckMail(mail)
	if err != nil {
		consoleLogWarning(r, "/user/", "mail - " + err.Error())
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		// CheckLogin is my own function, so I can not afraid of invalid runes in error
		panic("mail error - " + err.Error())
	}

	err = handlers.CheckPasswd(passwd)
	if err != nil {
		consoleLogWarning(r, "/user/", "password - " + err.Error())
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		// CheckLogin is my own function, so I can not afraid of invalid runes in error
		panic("password error - " + err.Error())
	}

	isUserExists, err := conn.IsUserExists(mail)
	if err != nil {
		consoleLogError(r, "/user/", "IsUserExists returned error " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("wrong request in database")
	}
	if isUserExists {
		consoleLogWarning(r, "/user/", "user " + BLUE + mail + NO_COLOR + " alredy exists")
		w.WriteHeader(http.StatusIMUsed) // 226
		panic("user " + login + " already exists")
	}

	err = conn.SetNewUser(mail, handlers.PasswdHash(passwd))
	if err != nil {
		consoleLogError(r, "/user/", "SetNewUser returned error " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("Cannot register this user")
	}
	w.WriteHeader(201)
	consoleLogSuccess(r, "/user/", "user " + BLUE + mail + NO_COLOR + " was created successfully. No response body")
}

// USER UPDATE BY PATCH METHOD. REQUEST AND RESPONSE DATA IS JSON.
// REQUEST SHOULD HAVE 'x-auth-token' HEADER
func (conn *ConnDB) updateUser(w http.ResponseWriter, r *http.Request) {
	var (
		uid int
		err error
		isExist bool
		user	User
		request map[string]interface{}
		token = r.Header.Get("x-auth-token")
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

	if token == "" {
		consoleLogWarning(r, "/user/", "token is empty")
		w.WriteHeader(http.StatusBadRequest) // 400
		panic("token is empty")
	}

	uid, err = handlers.TokenDecode(token)
	if err != nil {
		consoleLogWarning(r, "/user/", "token is empty")
		w.WriteHeader(http.StatusBadRequest) // 400
		panic("token is empty")
	}

	if !conn.session.IsUserLoggedByUid(uid) {
		consoleLogWarning(r, "/user/", "user #" + BLUE + strconv.Itoa(uid) + NO_COLOR + " is not logged")
		w.WriteHeader(http.StatusUnauthorized) // 401
		panic("user is not logged")
	}

	user, err = conn.GetUserByUid(uid)
	if err != nil {
		consoleLogError(r, "/user/", "GetUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("database request returned error")
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/", "request json decode failed - " + err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		panic("json decode failed")
	}

	user, err = fillUserStruct(request, user)
	if err != nil {
		consoleLogWarning(r, "/user/", err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		panic(err)
	}

	err = conn.UpdateUser(user)
	if err != nil {
		consoleLogError(r, "/user/", "UpdateUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("database request returned error")
	}

	// Проверить - принадлежит ли фото юзеру

	w.WriteHeader(http.StatusOK) // 200
	// Теперь не нужно обновлять токен аутентификации при изменении логина / почты
	consoleLogSuccess(r, "/user/", "user " + BLUE + user.Login + NO_COLOR + " was updated successfully. No response body")

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

	if passwd != sessionUser.User.Passwd {
		consoleLogWarning(r, "/user/", "password is incorrect")// + BLUE + passwd + " " + sessionUser.UserInfo.Passwd + NO_COLOR)
		w.WriteHeader(http.StatusNonAuthoritativeInfo) // 203
		panic("wrong password")
	}

	conn.session.DeleteUserSessionByUid(sessionUser.User.Uid)

	err = conn.DeleteUser(sessionUser.User.Uid)
	if err != nil {
		consoleLogError(r, "/user/", "DeleteUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		panic("database request returned error")
	}
	consoleLogSuccess(r, "/user/", "user " + BLUE + sessionUser.UserInfo.Mail + NO_COLOR + " was removed successfully. No response body")
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
