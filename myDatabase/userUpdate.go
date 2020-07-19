package myDatabase

import (
	. "MatchaServer/config"
	"MatchaServer/handlers"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

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
		tmpFloat, ok = arg.(float64)
		user.AvaPhotoID = int(tmpFloat)
		if !ok {
			return user, message, errors.New("wrong type of param")
		}
		if user.AvaPhotoID < 0 {
			return user, message, errors.New("this id is forbidden")
		}
		message += " avaPhotoID=" + BLUE + strconv.Itoa(user.AvaPhotoID) + NO_COLOR
	}

	if !usefullFieldsExists {
		return user, message, errors.New("no usefull fields found")
	}
	return user, message, nil
}

// USER UPDATE BY PATCH METHOD
// REQUEST BODY IS JSON
// REQUEST SHOULD HAVE 'x-auth-token' HEADER
// RESPONSE BODY IS JSON ONLY IN CASE OF ERROR. IN OTHER CASE - NO RESPONSE BODY
func (conn *ConnDB) userUpdate(w http.ResponseWriter, r *http.Request) {
	var (
		uid     int
		err     error
		user    User
		message string
		request map[string]interface{}
		token   = r.Header.Get("x-auth-token")
	)

	if token == "" {
		consoleLogWarning(r, "/user/update/", "token is empty")
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, `{"error":"`+"token is empty"+`"}`)
		return
	}

	uid, err = handlers.TokenUidDecode(token)
	if err != nil {
		consoleLogWarning(r, "/user/update/", "TokenUidDecode returned error - "+err.Error())
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, `{"error":"`+"token decoding error"+`"}`)
		return
	}

	if !conn.session.IsUserLoggedByUid(uid) {
		consoleLogWarning(r, "/user/update/", "user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+" is not logged")
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, `{"error":"`+"user is not logged"+`"}`)
		return
	}

	user, err = conn.GetUserByUid(uid)
	if err != nil {
		consoleLogError(r, "/user/update/", "GetUser returned error - "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"`+"database request returned error"+`"}`)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/update/", "request json decode failed - "+err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"json decode failed"+`"}`)
		return
	}

	user, message, err = fillUserStruct(request, user)
	if err != nil {
		consoleLogWarning(r, "/user/update/", err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+err.Error()+`"}`)
		return
	}

	consoleLog(r, "/user/update/", message)

	err = conn.UpdateUser(user)
	if err != nil {
		consoleLogError(r, "/user/update/", "UpdateUser returned error - "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"`+"database request returned error"+`"}`)
		return
	}

	// Проверить - принадлежит ли фото юзеру

	w.WriteHeader(http.StatusOK) // 200
	consoleLogSuccess(r, "/user/update/", "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+" was updated successfully. No response body")
}

// HTTP HANDLER FOR DOMAIN /user/update/
// UPDATE USER BY PATCH METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (conn *ConnDB) HttpHandlerUserUpdate(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST,PATCH,OPTIONS,DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)

		consoleLog(r, "/user/update/", "client wants to know what methods are allowed")

	} else if r.Method == "PATCH" {

		conn.userUpdate(w, r)

	} else {
		// ALL OTHER METHODS

		consoleLogWarning(r, "/user/update/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}
