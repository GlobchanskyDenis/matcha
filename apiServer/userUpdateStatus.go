package apiServer

import (
	// . "MatchaServer/config"
	"MatchaServer/handlers"
	"MatchaServer/errDef"
	"encoding/json"
	"net/http"
	"strconv"
)

// USER MAIL CONFIRM BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) userUpdateStatus(w http.ResponseWriter, r *http.Request) {
	var (
		message, mail, token string
		err                  error
		// user                 User
		request              map[string]interface{}
		item                 interface{}
		isExist, ok          bool
	)

	message = "request for MAIL CONFIRM was recieved"
	consoleLog(r, "/user/update/status/", message)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/user/update/status/", "request json decode failed - "+err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "json decode failed" + `"}`))
		return
	}

	item, isExist = request["x-reg-token"]
	if !isExist {
		consoleLogError(r, "/user/update/status/", "x-reg-token not exist in request")
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "x-reg-token not exist in request" + `"}`))
		return
	}

	token, ok = item.(string)
	if !ok {
		consoleLogError(r, "/user/update/status/", "x-reg-token has wrong type")
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		w.Write([]byte(`{"error":"` + "x-reg-token has wrong type" + `"}`))
		return
	}

	if token == "" {
		consoleLogError(r, "/user/update/status/", "x-reg-token is empty")
		w.WriteHeader(http.StatusUnauthorized) // 401
		w.Write([]byte(`{"error":"` + "x-reg-token is empty" + `"}`))
		return
	}

	mail, err = handlers.TokenMailDecode(token)
	if err != nil {
		consoleLogWarning(r, "/user/update/status/", "TokenMailDecode returned error - "+err.Error())
		w.WriteHeader(http.StatusUnauthorized) // 401
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	user, err := server.Db.GetUserByMail(mail)
	if errDef.RecordNotFound.IsOverlapWithError(err) {
		consoleLogWarning(r, "/user/update/status/", "GetUserByMail - record not found")
		w.WriteHeader(http.StatusUnauthorized) // 401
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	} else if err != nil {
		consoleLogWarning(r, "/user/update/status/", "GetUserByMail returned error - "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + `database returned error` + `"}`))
		return
	}

	// if user.Uid == 0 {
	// 	// it means that no such iser in database
	// 	consoleLogWarning(r, "/user/update/status/", "Mail doesnt exists in database")
	// 	w.WriteHeader(http.StatusUnauthorized) // 401
	// 	w.Write([]byte(`{"error":"` + `Mail doesnt exists in database` + `"}`))
	// 	return
	// }

	user.Status = "confirmed"

	err = server.Db.UpdateUser(user)
	if err != nil {
		consoleLogWarning(r, "/user/update/status/", "UpdateUser returned error - "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + `database returned error` + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	consoleLogSuccess(r, "/user/update/status/", "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
		" was updated its status successfully. No response body")
}

// HTTP HANDLER FOR DOMAIN /user/update/status . IT HANDLES:
// UPDATE USER STATUS BY PATCH METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) HandlerUserUpdateStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "PATCH" {
		server.userUpdateStatus(w, r)
	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)
		consoleLog(r, "/user/update/status/", "client wants to know what methods are allowed")
	} else {
		// ALL OTHERS METHODS
		consoleLogWarning(r, "/user/update/status/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
