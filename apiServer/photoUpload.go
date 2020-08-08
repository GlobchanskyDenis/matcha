package apiServer

import (
	. "MatchaServer/config"
	"MatchaServer/handlers"
	"MatchaServer/errDef"
	"encoding/json"
	"net/http"
	"strconv"
)

// USER MAIL CONFIRM BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) photoUpload(w http.ResponseWriter, r *http.Request) {
	var (
		message, body, token  string
		uid, pid			int
		err          error
		request      map[string]interface{}
		item         interface{}
		isExist, ok  bool
	)

	message = "request for PHOTO UPLOAD was recieved"
	consoleLog(r, "/photo/upload/", message)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		consoleLogError(r, "/photo/upload/", "request json decode failed - "+err.Error())
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "json decode failed" + `"}`))
		return
	}

	item, isExist = request["x-auth-token"]
	if !isExist {
		consoleLogError(r, "/photo/upload/", "x-auth-token not exist in request")
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "x-auth-token not exist in request" + `"}`))
		return
	}

	token, ok = item.(string)
	if !ok {
		consoleLogError(r, "/photo/upload/", "x-auth-token has wrong type")
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		w.Write([]byte(`{"error":"` + "x-auth-token has wrong type" + `"}`))
		return
	}

	if token == "" {
		consoleLogError(r, "/photo/upload/", "x-auth-token is empty")
		w.WriteHeader(http.StatusUnauthorized) // 401
		w.Write([]byte(`{"error":"` + "x-auth-token is empty" + `"}`))
		return
	}

	uid, err = handlers.TokenUidDecode(token)
	if err != nil {
		consoleLogWarning(r, "/photo/upload/", "TokenUidDecode returned error - "+err.Error())
		w.WriteHeader(http.StatusUnauthorized) // 401
		w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		return
	}

	item, isExist = request["photo"]
	if !isExist {
		consoleLogError(r, "/photo/upload/", "x-auth-token not exist in request")
		w.WriteHeader(http.StatusBadRequest) // 400
		w.Write([]byte(`{"error":"` + "x-auth-token not exist in request" + `"}`))
		return
	}

	body, ok = item.(string)
	if !ok {
		consoleLogError(r, "/photo/upload/", "photo has wrong type")
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		w.Write([]byte(`{"error":"` + "photo has wrong type" + `"}`))
		return
	}

	isExist, err = server.Db.IsUserExistsByUid(uid)
	if err != nil {
		consoleLogWarning(r, "/photo/upload/", "IsUserExistsByUid returned error - "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + `database returned error` + `"}`))
		return
	}
	if !isExist {
		consoleLogWarning(r, "/photo/upload/", "user record not found")
		w.WriteHeader(http.StatusUnauthorized) // 401
		w.Write([]byte(`{"error":"` + errDef.RecordNotFound.Error() + `"}`))
		return
	}

	pid, err = server.Db.SetNewPhoto(uid, []byte(body))
	if err != nil {
		consoleLogWarning(r, "/photo/upload/", "UpdateUser returned error - "+err.Error())
		w.WriteHeader(http.StatusInternalServerError) // 500
		w.Write([]byte(`{"error":"` + `database returned error` + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	consoleLogSuccess(r, "/photo/upload/", "user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+
		" was uploaded its photo successfully. photo id #"+BLUE+strconv.Itoa(pid)+NO_COLOR+". No response body")
}

// HTTP HANDLER FOR DOMAIN /user/update/status . IT HANDLES:
// UPDATE USER STATUS BY PATCH METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) HandlerPhotoUpload(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		server.photoUpload(w, r)
	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)
		consoleLog(r, "/photo/upload/", "client wants to know what methods are allowed")
	} else {
		// ALL OTHERS METHODS
		consoleLogWarning(r, "/photo/upload/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
