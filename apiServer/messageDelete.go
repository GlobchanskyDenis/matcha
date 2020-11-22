package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /message/delete/ . IT HANDLES:
// IT REMOVES MESSAGE BY ITS ID IF IT BELONGS TO THIS USER.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) MessageDelete(w http.ResponseWriter, r *http.Request) {
	var (
		myUid, mid    int
		mid64         float64
		message       Message
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
	item, isExist = requestParams["mid"]
	if !isExist {
		server.Logger.LogWarning(r, "Param mid expected")
		server.error(w, errors.NoArgument.WithArguments("отсутствует параметр mid", "param mid expected"))
		return
	}
	mid64, ok = item.(float64)
	if !ok {
		server.Logger.LogWarning(r, "Message id has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("nid имеет неверный тип", "nid has wrong type"))
		return
	}
	mid = int(mid64)

	message, err = server.Db.GetMessageByMid(mid)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "message not found. mid #"+BLUE+strconv.Itoa(mid)+NO_COLOR)
		server.error(w, errors.RecordNotFound)
		return
	} else if err != nil {
		server.Logger.LogError(r, "GetMessageByMid returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}
	if message.UidSender != myUid {
		server.Logger.LogError(r, "Message belongs to another user. Your uid #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
			" message owner uid #"+BLUE+strconv.Itoa(message.UidSender)+NO_COLOR)
		server.error(w, errors.ImpossibleToExecute.WithArguments("Сообщение не принадлежит вам", "Message isnt yours"))
		return
	}

	err = server.Db.DeleteMessage(mid)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "message not found. mid #"+BLUE+strconv.Itoa(mid)+NO_COLOR)
		server.error(w, errors.RecordNotFound)
		return
	} else if err != nil {
		server.Logger.LogError(r, "DeleteMessage returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "Message #"+BLUE+strconv.Itoa(mid)+NO_COLOR+" was deleted successfully.")
}
