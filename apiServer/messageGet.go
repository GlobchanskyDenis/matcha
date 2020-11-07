package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /message/get/ . IT HANDLES:
// IT RETURNS OWN USER DATA IN RESPONSE BY POST METHOD.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) MessageGet(w http.ResponseWriter, r *http.Request) {
	var (
		uid64           float64
		messages        []Message
		myUid, otherUid int
		requestParams   map[string]interface{}
		item            interface{}
		ok, isExist     bool
		err             error
		ctx             context.Context
	)

	ctx = r.Context()
	myUid = ctx.Value("uid").(int)
	requestParams, ok = ctx.Value("requestParams").(map[string]interface{})
	if !ok {
		server.Logger.LogWarning(r, "Request params has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Параметры запроса имеют неверный тип", "request params has wrong type"))
		return
	}
	item, isExist = requestParams["otherUid"]
	if !isExist {
		server.Logger.LogWarning(r, "Param otherUid expected")
		server.error(w, errors.NoArgument.WithArguments("отсутствует параметр otherUid", "param otherUid expected"))
		return
	}
	uid64, ok = item.(float64)
	if !ok {
		server.Logger.LogWarning(r, "Id of another user has wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("otherUid имеет неверный тип", "otherUid has wrong type"))
		return
	}
	otherUid = int(uid64)

	messages, err = server.Db.GetMessagesFromChat(myUid, otherUid)
	if err != nil {
		server.Logger.LogError(r, "GetMessagesFromChat returned error - "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}

	jsonMessages, err := json.Marshal(messages)
	if err != nil {
		server.Logger.LogError(r, "Marshal returned error "+err.Error())
		server.error(w, errors.MarshalError)
		return
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonMessages)
	server.Logger.LogSuccess(r, "Messages was handled successfully. Uid #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
		". Amount is #"+BLUE+strconv.Itoa(len(messages))+NO_COLOR)
}
