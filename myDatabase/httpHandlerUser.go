package myDatabase

import (
	"fmt"
	"errors"
	"net/http"
	"encoding/json"
	"MatchaServer/handlers"
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

func fillUserStruct(request map[string]interface{}, user User) (User, string, error) {
	var usefullFieldsExists, ok, isExist bool
	var message string
	var err error
	var tmpFloat float64

	message = "request for UPDATE was recieved: "

	arg, isExist := request["mail"]
	if isExist {
		usefullFieldsExists = true
		user.Mail, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckMail(user.Mail)
		if err != nil {
			return user, message, err
		}
		message += " mail=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["passwd"]
	if isExist {
		usefullFieldsExists = true
		tmp, ok := arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		user.Passwd = handlers.PasswdHash(tmp)
		err = handlers.CheckPasswd(tmp)
		if err != nil {
			return user, message, err
		}
		message += " password=hidden"
	}
	arg, isExist = request["fname"]
	if isExist {
		usefullFieldsExists = true
		user.Fname, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckName(user.Fname)
		if err != nil {
			return user, message, err
		}
		message += " fname=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["lname"]
	if isExist {
		usefullFieldsExists = true
		user.Lname, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckName(user.Lname)
		if err != nil {
			return user, message, err
		}
		message += " lname=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["age"]
	if isExist {
		usefullFieldsExists = true
		tmpFloat, ok = arg.(float64)
		user.Age = int(tmpFloat)
		if !ok {
			// fmt.Printf("%V %T", request["age"],request["age"])
			return user, message, errors.New("wrong type of param")
		}
		if user.Age > 80 || user.Age < 14 {
			return user, message, errors.New("this age is forbidden")
		}
		message += " age=" + BLUE + strconv.Itoa(user.Age) + NO_COLOR
	}
	arg, isExist = request["gender"]
	if isExist {
		usefullFieldsExists = true
		user.Gender, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckGender(user.Gender)
		if err != nil {
			return user, message, err
		}
		message += " gender=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["orientation"]
	if isExist {
		usefullFieldsExists = true
		user.Orientation, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckOrientation(user.Orientation)
		if err != nil {
			return user, message, err
		}
		message += " orientation=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["biography"]
	if isExist {
		usefullFieldsExists = true
		user.Biography, ok = arg.(string)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		err = handlers.CheckBiography(user.Biography)
		if err != nil {
			return user, message, err
		}
		message += " biography=" + BLUE + arg.(string) + NO_COLOR
	}
	arg, isExist = request["avaPhotoID"]
	if isExist {
		usefullFieldsExists = true
		user.AvaPhotoID, ok = arg.(int)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		if user.AvaPhotoID < 0 {
			return user, message, errors.New("this age is forbidden")
		}
		message += " avaPhotoID=" + BLUE + strconv.Itoa(arg.(int)) + NO_COLOR
	}

	if !usefullFieldsExists {
		return user, message, errors.New("no usefull fields found")
	}
	return user, message, nil
}

// USER REGISTRATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (conn *ConnDB) regUser(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, passwd string
		err error
		request map[string]interface{}
		isExist bool
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/", "request json decode failed - " + err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"` + "json decode failed" + `"}`)
		return
	}
	
	arg, isExist := request["mail"]
	if !isExist {
		consoleLogWarning(r, "/user/", "mail not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"` + "mail not exist" + `"}`)
		return
	}
	mail = arg.(string)

	arg, isExist = request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/user/", "password not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"` + "password not exist" + `"}`)
		return
	}
	passwd = arg.(string)

	message = "request was recieved, mail: " + BLUE + mail + NO_COLOR +
		" password: hidden"
	consoleLog(r, "/user/", message)

	if mail == "" || passwd == "" {
		consoleLogWarning(r, "/user/", "mail or password is empty")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"` + "mail or password is empty" + `"}`)
		return
	}

	err = handlers.CheckMail(mail)
	if err != nil {
		consoleLogWarning(r, "/user/", "mail - " + err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		// CheckMail is my own function, so I can not afraid of invalid runes in error
		fmt.Fprintf(w, `{"error":"` + "mail error - " + err.Error() + `"}`)
		return
	}

	err = handlers.CheckPasswd(passwd)
	if err != nil {
		consoleLogWarning(r, "/user/", "password - " + err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		// CheckPasswd is my own function, so I can not afraid of invalid runes in error
		fmt.Fprintf(w, `{"error":"` + "password error - " + err.Error() + `"}`)
		return
	}

	isUserExists, err := conn.IsUserExists(mail)
	if err != nil {
		consoleLogError(r, "/user/", "IsUserExists returned error " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"` + "database request returned error" + `"}`)
		return
	}
	if isUserExists {
		consoleLogWarning(r, "/user/", "user " + BLUE + mail + NO_COLOR + " alredy exists")
		w.WriteHeader(http.StatusNotAcceptable) // 406
		fmt.Fprintf(w, `{"error":"` + "user " + mail + " already exists" + `"}`)
		return
	}

	err = conn.SetNewUser(mail, handlers.PasswdHash(passwd))
	if err != nil {
		consoleLogError(r, "/user/", "SetNewUser returned error " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"` + "Cannot register this user" + `"}`)
		return
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
		user	User
		message string
		request map[string]interface{}
		token = r.Header.Get("x-auth-token")
	)

	if token == "" {
		consoleLogWarning(r, "/user/", "token is empty")
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, `{"error":"` + "token is empty" + `"}`)
		return
	}

	uid, err = handlers.TokenDecode(token)
	if err != nil {
		consoleLogWarning(r, "/user/", "TokenDecode returned error - " + err.Error())
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, `{"error":"` + "token decoding error" + `"}`)
		return
	}

	if !conn.session.IsUserLoggedByUid(uid) {
		consoleLogWarning(r, "/user/", "user #" + BLUE + strconv.Itoa(uid) + NO_COLOR + " is not logged")
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, `{"error":"` + "user is not logged" + `"}`)
		return
	}

	user, err = conn.GetUserByUid(uid)
	if err != nil {
		consoleLogError(r, "/user/", "GetUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"` + "database request returned error" + `"}`)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/", "request json decode failed - " + err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"` + "json decode failed" + `"}`)
		return
	}

	user, message, err = fillUserStruct(request, user)
	if err != nil {
		consoleLogWarning(r, "/user/", err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"` + err.Error() + `"}`)
		return
	}

	consoleLog(r, "/user/", message)

	err = conn.UpdateUser(user)
	if err != nil {
		consoleLogError(r, "/user/", "UpdateUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"` + "database request returned error" + `"}`)
		return
	}

	// Проверить - принадлежит ли фото юзеру

	w.WriteHeader(http.StatusOK) // 200
	consoleLogSuccess(r, "/user/", "user #" + BLUE + strconv.Itoa(user.Uid) + NO_COLOR + " was updated successfully. No response body")
}

// USER REMOVE BY DELETE METHOD. NO REQUEST DATA. RESPONSE DATA IS JSON ONLY IN CASE OF ERROR.
// REQUEST SHOULD HAVE 'x-auth-token' HEADER
func (conn *ConnDB) deleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		message string
		err error
		token = r.Header.Get("x-auth-token")
		user User
		request map[string]interface{}
		passwd string
		uid int
	)

	message = "request for DELETE was recieved"//: token=" + BLUE + token + NO_COLOR
	consoleLog(r, "/user/", message)

	uid, err = handlers.TokenDecode(token)
	if err != nil {
		consoleLogWarning(r, "/user/", "TokenDecode returned error - " + err.Error())
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, `{"error":"` + err.Error() + `"}`)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/", "request json decode failed - " + err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"` + "json decode failed" + `"}`)
		return
	}

	arg, isExist := request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/user/", "password not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"` + "password not exist" + `"}`)
		return
	}
	passwd = handlers.PasswdHash(arg.(string))

	user, err = conn.GetUserByUid(uid)
	if err != nil {
		consoleLogError(r, "/user/", "GetUserByUid returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"` + err.Error() + `"}`)
		return
	}

	if passwd != user.Passwd {
		consoleLogWarning(r, "/user/", "password is incorrect")// + BLUE + passwd + " " + sessionUser.UserInfo.Passwd + NO_COLOR)
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"` + "wrong password" + `"}`)
		return
	}

	conn.session.DeleteUserSessionByUid(user.Uid)

	err = conn.DeleteUser(user.Uid)
	if err != nil {
		consoleLogError(r, "/user/", "DeleteUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"` + "database request returned error" + `"}`)
		return
	}
	
	w.WriteHeader(http.StatusOK) // 200
	consoleLogSuccess(r, "/user/", "user #" + BLUE + strconv.Itoa(user.Uid) + NO_COLOR + " was removed successfully. No response body")
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
