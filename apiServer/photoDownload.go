package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /photo/download/ . IT HANDLES:
// DOWNLOAD USER PHOTOS BY UID
func (server *Server) PhotoDownload(w http.ResponseWriter, r *http.Request) {
	var (
		authorUid, myUid int
		tmpFloat64       float64
		err              error
		requestParams    map[string]interface{}
		item             interface{}
		ctx              context.Context
		isExist, ok      bool
	)

	ctx = r.Context()
	requestParams = ctx.Value("requestParams").(map[string]interface{})
	myUid = ctx.Value("uid").(int)

	item, isExist = requestParams["uid"]
	if !isExist {
		server.Logger.LogWarning(r, "uid not exist in request")
		server.error(w, errors.NoArgument.WithArguments("Поле uid отсутствует", "uid field expected"))
		return
	}

	tmpFloat64, ok = item.(float64)
	if !ok {
		server.Logger.LogWarning(r, "uid has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Поле uid имеет неверный тип", "uid field has wrong type"))
		return
	}
	authorUid = int(tmpFloat64)

	photos, err := server.Db.GetPhotosByUid(authorUid)
	if err != nil {
		server.Logger.LogError(r, "GetPhotosByUid returned error - "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	jsonPhotos, err := json.Marshal(photos)

	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "user #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
		" was downloaded photos of user #"+BLUE+strconv.Itoa(authorUid)+NO_COLOR+
		" successfully. Amount of photos: "+BLUE+strconv.Itoa(len(photos))+NO_COLOR)
	w.Write(jsonPhotos)
}
