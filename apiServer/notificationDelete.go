package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /notification/get/ . IT HANDLES:
// IT RETURNS OWN USER DATA IN RESPONSE BY POST METHOD.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) NotificationDelete(w http.ResponseWriter, r *http.Request) {
	var (
		nid           int
		nid64         float64
		requestParams map[string]interface{}
		item          interface{}
		ok, isExist   bool
		err           error
		ctx           context.Context
	)

	ctx = r.Context()
	requestParams, ok = ctx.Value("requestParams").(map[string]interface{})
	if !ok {
		server.Logger.LogWarning(r, "Request params has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Параметры запроса имеют неверный тип", "request params has wrong type"))
		return
	}
	item, isExist = requestParams["nid"]
	if !isExist {
		server.Logger.LogWarning(r, "Param nid expected")
		server.error(w, errors.NoArgument.WithArguments("отсутствует параметр nid", "param nid expected"))
		return
	}
	nid64, ok = item.(float64)
	if !ok {
		server.Logger.LogWarning(r, "Id of photo has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("nid имеет неверный тип", "nid has wrong type"))
		return
	}
	nid = int(nid64)

	err = server.Db.DeleteNotif(nid)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "DeleteNotif returned record not found")
		server.error(w, errors.RecordNotFound)
		return
	} else if err != nil {
		server.Logger.LogError(r, "DeleteNotif returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "Notification #"+BLUE+strconv.Itoa(nid)+NO_COLOR+" was deleted successfully.")
}
