package httpHandlers

import (
	. "MatchaServer/config"
	"MatchaServer/handlers"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// USER MAIL CONFIRM BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (conn *ConnAll) userUpdateStatus(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, token string
		err     error
		user    User
		request map[string]interface{}
		requestItem interface{}
		isExist, ok bool
	)

	message = "request for MAIL CONFIRM was recieved"
	consoleLog(r, "/user/update/status/", message)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/update/status/", "request json decode failed - " + err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"json decode failed"+`"}`)
		return
	}

	requestItem, isExist = request["x-reg-token"]
	if !isExist {
		consoleLogError(r, "/user/update/status/", "x-reg-token not exist in request")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"x-reg-token not exist in request"+`"}`)
		return
	}

	token, ok = requestItem.(string)
	if !ok {
		consoleLogError(r, "/user/update/status/", "x-reg-token has wrong type")
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, `{"error":"`+"x-reg-token has wrong type"+`"}`)
		return
	}

	if token == "" {
		consoleLogError(r, "/user/update/status/", "x-reg-token is empty")
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, `{"error":"`+"x-reg-token is empty"+`"}`)
		return
	}

	mail, err = handlers.TokenMailDecode(token)
	if err != nil {
		consoleLogWarning(r, "/user/update/status/", "TokenMailDecode returned error - " + err.Error())
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, `{"error":"`+err.Error()+`"}`)
		return
	}

	user, err = conn.Db.GetUserByMail(mail)
	if err != nil {
		consoleLogWarning(r, "/user/update/status/", "GetUserByMail returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"`+`database returned error`+`"}`)
		return
	}

	if user == (User{}) {
		consoleLogWarning(r, "/user/update/status/", "Mail doesnt exists in database")
		w.WriteHeader(http.StatusUnauthorized) // 401
		fmt.Fprintf(w, `{"error":"`+`Mail doesnt exists in database`+`"}`)
		return
	}

	user.AccType = "confirmed"

	err = conn.Db.UpdateUser(user)
	if err != nil {
		consoleLogWarning(r, "/user/update/status/", "UpdateUser returned error - " + err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		fmt.Fprintf(w, `{"error":"`+`database returned error`+`"}`)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	consoleLogSuccess(r, "/user/update/status/", "user #" + BLUE + strconv.Itoa(user.Uid) + NO_COLOR +
		" was updated its status successfully. No response body")
}

// HTTP HANDLER FOR DOMAIN /user/update/status . IT HANDLES:
// UPDATE USER STATUS BY PATCH METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (conn *ConnAll) HttpHandlerUserUpdateStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "PATCH" {
		conn.userUpdateStatus(w, r)
	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)
		consoleLog(r, "/auth/", "client wants to know what methods are allowed")
	} else {
		// ALL OTHERS METHODS
		consoleLogWarning(r, "/auth/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}