package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /message/set/inactive/ . IT HANDLES:
// IT SET STATUS OF MESSAGE INACTIVE (IT MEANS THAT YOU ARE ALREADY SEEN IT).
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) MessageSetInactive(w http.ResponseWriter, r *http.Request) {
	var (
		uid64         float64
		message       Message
		myUid, mid    int
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
		server.error(w, errors.InvalidArgument.WithArguments("Параметры запроса имеют неверный тип", "request params has wrong type"))
		return
	}
	item, isExist = requestParams["mid"]
	if !isExist {
		server.Logger.LogWarning(r, "Param mid expected")
		server.error(w, errors.NoArgument.WithArguments("отсутствует параметр mid", "param mid expected"))
		return
	}
	uid64, ok = item.(float64)
	if !ok {
		server.Logger.LogWarning(r, "Message id has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("id сообщения имеет неверный тип", "message id has wrong type"))
		return
	}
	mid = int(uid64)

	message, err = server.Db.GetMessageByMid(mid)
	if err != nil {
		server.Logger.LogError(r, "GetMessagesFromChat returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}
	if message.UidReceiver != myUid {
		server.Logger.LogError(r, "message #"+BLUE+strconv.Itoa(mid)+NO_COLOR+" belongs to another user. Your uid #"+
			BLUE+strconv.Itoa(myUid)+NO_COLOR+" message owner uid #"+BLUE+strconv.Itoa(message.UidReceiver)+NO_COLOR)
		server.error(w, errors.ImpossibleToExecute.WithArguments("это фото было отправлено не вам",
			"this message was adressed not to you"))
		return
	}

	err = server.Db.SetMessageInactive(mid)
	if err != nil {
		server.Logger.LogError(r, "SetMessageInactive returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "Messages was handled successfully. Uid #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
		". Mid #"+BLUE+strconv.Itoa(mid)+NO_COLOR)
}
