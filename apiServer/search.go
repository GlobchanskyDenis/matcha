package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errDef"
	"encoding/json"
	"net/http"
	"strconv"
	// "context"
)

func (server *Server) searchAll(w http.ResponseWriter, r *http.Request) {
	var filter = r.URL.Query().Get("filter")

	users, err := server.Db.SearchUsersByOneFilter(filter)
	if err != nil {
		server.LogError(r, "SearchUsersByOneFilter returned error "+err.Error())
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
	w.Write([]byte(jsonUsers))
	server.LogSuccess(r, "array of "+BLUE+"all"+NO_COLOR+
		" users was transmitted. Users amount "+strconv.Itoa(len(users)))
}

func (server *Server) searchLogged(w http.ResponseWriter, r *http.Request) {

	users, err := server.Db.GetLoggedUsers(server.session.GetLoggedUsersUidSlice())
	if err != nil {
		server.LogError(r, "GetLoggedUsers returned error"+err.Error())
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
	server.LogSuccess(r, "array of "+BLUE+"logged"+NO_COLOR+
		" users was transmitted. Users amount "+strconv.Itoa(len(users)))
}

// Фильтры по онлайну, возрасту, рейтингу, локации (или радиус от заданной точки), интересам, 
// обязательный фильтр по соответствию пол/ориентация
// если поля не заполнены - показываем всех

func (server *Server) Search(w http.ResponseWriter, r *http.Request) {
	// var {
	// 	requestParams	map[string]interface{}
	// 	ctx				context.Context
	// }
	var filter = r.URL.Query().Get("filter")
	// ctx = r.Context()
	// requestParams = ctx.Value("requestParams").(map[string]interface{})
	

	server.Log(r, "request was recieved with filter "+BLUE+filter+NO_COLOR)

	if filter != "all" && filter != "logged" {
		server.LogWarning(r, "filter "+BLUE+filter+NO_COLOR+" not exist")
		errDef.InvalidArgument.WithArguments("Значение поля filter недопустимо", "filter field has wrong value")
		return
	}

	if filter == "all" {
		server.searchAll(w, r)
	}
	if filter == "logged" {
		server.searchLogged(w, r)
	}
}
