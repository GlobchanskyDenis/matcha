package apiServer

import (
	"MatchaServer/apiServer/searchFilters"
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// Фильтры по онлайну, возрасту, рейтингу, локации (или радиус от заданной точки), интересам,
// обязательный фильтр по соответствию пол/ориентация
// если поля не заполнены - показываем всех
// обязательный фильтр - исключать пользователей из черного списка и игноров

func (server *Server) Search(w http.ResponseWriter, r *http.Request) {
	var (
		requestParams   map[string]interface{}
		ctx             context.Context
		filters         *searchFilters.Filters
		uid             int
		err             error
		user            User
		searchUsers     []SearchUser
		sexRestrictions string
	)

	ctx = r.Context()
	uid = ctx.Value("uid").(int)
	requestParams = ctx.Value("requestParams").(map[string]interface{})

	user, err = server.Db.GetUserByUid(uid)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "User with uid #"+BLUE+strconv.Itoa(uid)+NO_COLOR+" not found")
		server.error(w, errors.UserNotExist)
		return
	} else if err != nil {
		server.Logger.LogError(r, "GetUserByUid returned error "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}
	if user.AvaID == nil || user.Fname == "" || user.Lname == "" {
		server.Logger.LogWarning(r, "Your user #"+BLUE+strconv.Itoa(uid)+NO_COLOR+" cannot make search")
		server.error(w, errors.ImpossibleToExecute.WithArguments("Сначала нужно заполнить ваши имя, фамилию и аватар",
			"You need to fill name, surname and avatar first"))
		return
	}

	filters = searchFilters.New()
	err = filters.Parse(requestParams, uid, server.Db, &server.Session)
	if err != nil {
		server.Logger.LogWarning(r, "Cannot parse filter: "+BLUE+err.Error()+NO_COLOR)
		server.error(w, errors.InvalidArgument.WithArguments(err))
		return
	}

	server.Logger.Log(r, "search filters: "+BLUE+filters.Print()+NO_COLOR)

	sexRestrictions = searchFilters.PrepareSexRestrictions(user)
	query := filters.PrepareQuery(sexRestrictions, &server.Logger)
	searchUsers, err = server.Db.GetUsersByQuery(query, user)
	if err != nil {
		server.Logger.LogError(r, "GetUsersByQuery returned error "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	jsonUsers, err := json.Marshal(searchUsers)
	if err != nil {
		server.Logger.LogError(r, "Marshal returned error"+err.Error())
		server.error(w, errors.MarshalError)
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonUsers)
	server.Logger.LogSuccess(r, "array of users was transmitted. Users amount "+strconv.Itoa(len(searchUsers)))
}
