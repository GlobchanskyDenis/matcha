package apiServer

import (
	. "MatchaServer/config"
	"MatchaServer/handlers"
	"encoding/json"
	"net/http"
)

// USER REGISTRATION BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) userReg(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, passwd, token string
		err                          error
		request                      map[string]interface{}
		isExist                      bool
		user						 User
	)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/reg/", "request json decode failed - "+err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "json decode failed" + `"}`))
		return
	}

	arg, isExist := request["mail"]
	if !isExist {
		consoleLogWarning(r, "/user/reg/", "mail not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "mail not exist" + `"}`))
		return
	}
	mail = arg.(string)

	arg, isExist = request["passwd"]
	if !isExist {
		consoleLogWarning(r, "/user/reg/", "password not exist")
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "password not exist" + `"}`))
		return
	}
	passwd = arg.(string)

	message = "request was recieved, mail: " + BLUE + mail + NO_COLOR +
		" password: hidden"
	consoleLog(r, "/user/reg/", message)

	if mail == "" || passwd == "" {
		consoleLogWarning(r, "/user/reg/", "mail or password is empty")
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		w.Write([]byte(`{"error":"` + "mail or password is empty" + `"}`))
		return
	}

	err = handlers.CheckMail(mail)
	if err != nil {
		consoleLogWarning(r, "/user/reg/", "mail - "+err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		// CheckMail is my own function, so I can not afraid of invalid runes in error
		w.Write([]byte(`{"error":"` + "mail error - " + err.Error() + `"}`))
		return
	}

	err = handlers.CheckPasswd(passwd)
	if err != nil {
		consoleLogWarning(r, "/user/reg/", "password - "+err.Error())
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		// CheckPasswd is my own function, so I can not afraid of invalid runes in error
		w.Write([]byte(`{"error":"` + "password error - " + err.Error() + `"}`))
		return
	}

	isUserExists, err := server.Db.IsUserExistsByMail(mail)
	if err != nil {
		consoleLogError(r, "/user/reg/", "IsUserExists returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "database request returned error" + `"}`))
		return
	}
	if isUserExists {
		consoleLogWarning(r, "/user/reg/", "user "+BLUE+mail+NO_COLOR+" alredy exists")
		w.WriteHeader(http.StatusNotAcceptable) // 406
		w.Write([]byte(`{"error":"` + "user " + mail + " already exists" + `"}`))
		return
	}

	user, err = server.Db.SetNewUser(mail, handlers.PasswdHash(passwd))
	if err != nil {
		consoleLogError(r, "/user/reg/", "SetNewUser returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "Cannot register this user" + `"}`))
		return
	}

	token, err = handlers.TokenMailEncode(mail)
	if err != nil {
		consoleLogError(r, "/user/reg/", "TokenMailEncode returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "Cannot create token for this user" + `"}`))
		return
	}

	// user, err = server.Db.GetUserByMail(mail)
	// if err != nil {
	// 	consoleLogError(r, "/user/reg/", "GetUserByMail returned error "+err.Error())
	// 	w.WriteHeader(http.StatusInternalServerError) // 500
	// 	w.Write([]byte(`{"error":"` + "Database returned error" + `"}`))
	// 	return
	// }

	err = server.Db.SetNewDevice(user.Uid, r.UserAgent())
	if err != nil {
		consoleLogError(r, "/user/reg/", "SetNewDevice returned error "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + "Database returned error" + `"}`))
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
func (server *Server) HttpHandlerUserReg(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "POST,PATCH,OPTIONS,DELETE")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type,x-auth-token")

	if r.Method == "POST" {

		server.userReg(w, r)

	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)

		consoleLog(r, "/user/reg/", "client wants to know what methods are allowed")

	} else {
		// ALL OTHERS METHODS

		consoleLogWarning(r, "/user/reg/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405

	}
}