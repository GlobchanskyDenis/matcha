package myDatabase

import (
	. "MatchaServer/config"
	"MatchaServer/handlers"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// USER REMOVE BY DELETE METHOD. NO REQUEST BODY. RESPONSE BODY IS JSON ONLY IN CASE OF ERROR.
// REQUEST SHOULD HAVE 'x-auth-token' HEADER
func (conn *ConnDB) userDelete(w http.ResponseWriter, r *http.Request) {
	var (
		message string
		err     error
		token   = r.Header.Get("x-auth-token")
		user    User
		request map[string]interface{}
		passwd  string
		uid     int
	)

	message = "request for DELETE was recieved"
	consoleLog(r, "/user/delete/", message)

	uid, err = handlers.TokenAuthDecode(token)
	if err != nil {
		consoleLogWarning(r, "/user/delete/", "TokenDecode returned error - " + err.Error())
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, `{"error":"` + err.Error() + `"}`)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/delete/", "request json decode failed - " + err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"` + "json decode failed" + `"}`)
		return
	}

	arg, isExist := request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/user/delete/", "password not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"` + "password not exist" + `"}`)
		return
	}
	passwd = handlers.PasswdHash(arg.(string))

	user, err = conn.GetUserByUid(uid)
	if err != nil {
		consoleLogError(r, "/user/delete/", "GetUserByUid returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"` + err.Error() + `"}`)
		return
	}

	if passwd != user.Passwd {
		consoleLogWarning(r, "/user/delete/", "password is incorrect")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"wrong password"+`"}`)
		return
	}

	conn.session.DeleteUserSessionByUid(user.Uid)

	err = conn.DeleteUser(user.Uid)
	if err != nil {
		consoleLogError(r, "/user/delete/", "DeleteUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"`+"database request returned error"+`"}`)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	consoleLogSuccess(r, "/user/delete/", "user #" + BLUE + strconv.Itoa(user.Uid) + NO_COLOR +
		" was removed successfully. No response body")
}

// HTTP HANDLER FOR DOMAIN /user/delete/
// DELETE USER BY DELETE METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (conn *ConnDB) HttpHandlerUserDelete(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "OPTIONS,DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)

		consoleLog(r, "/user/delete/", "client wants to know what methods are allowed")

	} else if r.Method == "DELETE" {

		conn.userDelete(w, r)

	} else {
		// ALL OTHERS METHODS

		consoleLogWarning(r, "/user/delete/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}