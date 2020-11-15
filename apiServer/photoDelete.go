package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /photo/delete/ . IT HANDLES:
// IT DELETES SELECTED PHOTO.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) PhotoDelete(w http.ResponseWriter, r *http.Request) {
	var (
		myUid, pid    int
		pid64         float64
		photo         Photo
		requestParams map[string]interface{}
		item          interface{}
		ok, isExist   bool
		err           error
		ctx           context.Context
	)

	ctx = r.Context()
	myUid = ctx.Value("uid").(int)
	requestParams, ok = ctx.Value("requestParams").(map[string]interface{})
	if !ok {
		server.Logger.LogWarning(r, "Request params has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Параметры запроса имеют неверный тип",
			"request params has wrong type"))
		return
	}
	item, isExist = requestParams["pid"]
	if !isExist {
		server.Logger.LogWarning(r, "Param pid expected")
		server.error(w, errors.NoArgument.WithArguments("отсутствует параметр pid", "param pid expected"))
		return
	}
	pid64, ok = item.(float64)
	if !ok {
		server.Logger.LogWarning(r, "Photo id has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("pid имеет неверный тип", "pid has wrong type"))
		return
	}
	pid = int(pid64)

	photo, err = server.Db.GetPhotoByPid(pid)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "photo not found. pid #"+BLUE+strconv.Itoa(pid)+NO_COLOR)
		server.error(w, errors.RecordNotFound)
		return
	} else if err != nil {
		server.Logger.LogError(r, "GetPhotoByPid returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}
	if photo.Uid != myUid {
		server.Logger.LogError(r, "Photo belongs to another user. Your uid #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
			" photo owner uid #"+BLUE+strconv.Itoa(photo.Uid)+NO_COLOR)
		server.error(w, errors.ImpossibleToExecute.WithArguments("Фото не принадлежит вам", "Photo isnt yours"))
		return
	}

	err = server.Db.DeletePhoto(pid)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "photo not found. mid #"+BLUE+strconv.Itoa(pid)+NO_COLOR)
		server.error(w, errors.RecordNotFound)
		return
	} else if err != nil {
		server.Logger.LogError(r, "DeletePhoto returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "Photo #"+BLUE+strconv.Itoa(pid)+NO_COLOR+" was deleted successfully.")
}
