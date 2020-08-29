package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errDef"
	"MatchaServer/handlers"
	"encoding/json"
	"net/http"
	"strconv"
)

// USER MAIL CONFIRM BY POST METHOD. REQUEST AND RESPONSE DATA IS JSON
func (server *Server) photoDownload(w http.ResponseWriter, r *http.Request) {
	var (
		message, token        string
		myUid, authorUid      int
		tmpFloat64            float64
		err                   error
		request               map[string]interface{}
		item                  interface{}
		isExist, isLogged, ok bool
	)

	message = "request for PHOTO DOWNLOAD was recieved"
	server.Log(r, "/photo/download/", message)

	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		server.LogError(r, "/photo/download/", "request json decode failed - "+err.Error())
		server.error(w, errDef.InvalidRequestBody)
		return
	}

	item, isExist = request["uid"]
	if !isExist {
		server.LogWarning(r, "/photo/download/", "uid not exist in request")
		server.error(w, errDef.NoArgument.WithArguments("Поле uid отсутствует", "uid field expected"))
		return
	}

	tmpFloat64, ok = item.(float64)
	if !ok {
		server.LogWarning(r, "/photo/download/", "uid has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле uid имеет неверный тип", "uid field has wrong type"))
		return
	}
	authorUid = int(tmpFloat64)

	item, isExist = request["x-auth-token"]
	if !isExist {
		server.LogWarning(r, "/photo/download/", "x-auth-token not exist in request")
		server.error(w, errDef.NoArgument.WithArguments("Поле x-auth-token отсутствует", "x-auth-token field expected"))
		return
	}

	token, ok = item.(string)
	if !ok {
		server.LogWarning(r, "/photo/download/", "x-auth-token has wrong type")
		server.error(w, errDef.InvalidArgument.WithArguments("Поле x-auth-token имеет неверный тип", "x-auth-token field has wrong type"))
		return
	}

	if token == "" {
		server.LogWarning(r, "/photo/download/", "x-auth-token is empty")
		server.error(w, errDef.UserNotLogged)
		return
	}

	myUid, err = handlers.TokenUidDecode(token)
	if err != nil {
		server.LogWarning(r, "/photo/download/", "TokenUidDecode returned error - "+err.Error())
		server.error(w, errDef.UserNotLogged)
		return
	}

	isLogged = server.session.IsUserLoggedByUid(myUid)
	if !isLogged {
		server.LogWarning(r, "/photo/download/", "User #"+strconv.Itoa(myUid)+" is not logged")
		server.error(w, errDef.UserNotLogged)
		return
	}

	photos, err := server.Db.GetPhotosByUid(authorUid)
	if err != nil {
		server.LogError(r, "/photo/download/", "GetPhotosByUid returned error - "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	jsonPhotos, err := json.Marshal(photos)

	w.WriteHeader(http.StatusOK) // 200
	server.LogSuccess(r, "/photo/download/", "user #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
		" was downloaded photos of user #"+BLUE+strconv.Itoa(authorUid)+NO_COLOR+
		" successfully. Amount of photos: "+BLUE+strconv.Itoa(len(photos))+NO_COLOR)
	w.Write(jsonPhotos)
}

// HTTP HANDLER FOR DOMAIN /user/update/status . IT HANDLES:
// UPDATE USER STATUS BY PATCH METHOD
// SEND HTTP OPTIONS IN CASE OF OPTIONS METHOD
func (server *Server) HandlerPhotoDownload(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "PATCH,OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		server.photoDownload(w, r)
	} else if r.Method == "OPTIONS" {
		// OPTIONS METHOD (CLIENT WANTS TO KNOW WHAT METHODS AND HEADERS ARE ALLOWED)
		server.Log(r, "/photo/download/", "client wants to know what methods are allowed")
	} else {
		// ALL OTHERS METHODS
		server.LogWarning(r, "/photo/download/", "wrong request method")
		w.WriteHeader(http.StatusMethodNotAllowed) // 405
	}
}
