package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/apiServer/searchFilters"
	"MatchaServer/errDef"
	"encoding/json"
	"net/http"
	"strconv"
	"context"
)

// Фильтры по онлайну, возрасту, рейтингу, локации (или радиус от заданной точки), интересам, 
// обязательный фильтр по соответствию пол/ориентация
// если поля не заполнены - показываем всех

func (server *Server) Search(w http.ResponseWriter, r *http.Request) {
	var (
		requestParams	map[string]interface{}
		ctx				context.Context
		filters			*searchFilters.Filters
		uid				int
		err				error
		user			User
		users			[]User
		sexRestrictions string
	)

	ctx = r.Context()
	uid = ctx.Value("uid").(int)
	requestParams = ctx.Value("requestParams").(map[string]interface{})
	filters = searchFilters.New()
	err = filters.Parse(requestParams, uid, server.Db, &server.Session)
	if err != nil {
		server.LogWarning(r, "Cannot parse filter: "+BLUE+err.Error()+NO_COLOR)
		server.error(w, errDef.InvalidArgument.WithArguments(err))
		return
	}

	server.Log(r, "search filters: "+BLUE+filters.Print()+NO_COLOR)

	user, err = server.Db.GetUserByUid(uid)
	if errDef.RecordNotFound.IsOverlapWithError(err) {
		server.LogWarning(r, "User with uid #"+BLUE+strconv.Itoa(uid)+NO_COLOR+" not found")
		server.error(w, errDef.UserNotExist)
		return
	} else if err != nil {
		server.LogError(r, "GetUserByUid returned error "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}
	sexRestrictions = searchFilters.PrepareSexRestrictions(user)
	query := filters.PrepareQuery(sexRestrictions)
	users, err = server.Db.GetUsersByQuery(query)
	if err != nil {
		server.LogError(r, "GetUsersByQuery returned error "+err.Error())
		server.error(w, errDef.DatabaseError)
		return
	}

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		server.LogError(r, "Marshal returned error"+err.Error())
		server.error(w, errDef.MarshalError)
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonUsers)
	server.LogSuccess(r, "array of users was transmitted. Users amount "+strconv.Itoa(len(users)))
}
