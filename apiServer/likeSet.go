package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /like/set/ . IT HANDLES:
// IT SET LIKE TO TARGET USER.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) LikeSet(w http.ResponseWriter, r *http.Request) {
	var (
		uid64             float64
		myUid, otherUid   int
		requestParams     map[string]interface{}
		item              interface{}
		ok, isExist       bool
		err               error
		ctx               context.Context
		myUser, otherUser User
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

	// Вот тут проверить чтобы у этого юзера были фотки
	myUser, err = server.Db.GetUserByUid(myUid)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "Your user#"+BLUE+strconv.Itoa(myUid)+NO_COLOR+" not exists")
		server.error(w, errors.ImpossibleToExecute.WithArguments("Вашего пользователя не существует", "Your user isnt exist"))
		return
	} else if err != nil {
		server.Logger.LogError(r, "GetUserByUid returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	} else if myUser.AvaID == nil {
		server.Logger.LogWarning(r, "Your user#"+BLUE+strconv.Itoa(otherUid)+NO_COLOR+" have no avatar")
		server.error(w, errors.ImpossibleToExecute.WithArguments("Нельзя лайкать не имея аватарки",
			"Forbidden to like without avatar"))
		return
	}

	otherUser, err = server.Db.GetUserByUid(otherUid)
	if errors.RecordNotFound.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "User#"+BLUE+strconv.Itoa(otherUid)+NO_COLOR+" not exists")
		server.error(w, errors.ImpossibleToExecute.WithArguments("Такого пользователя не существует", "This user isnt exist"))
		return
	} else if err != nil {
		server.Logger.LogError(r, "GetUserByUid returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	} else if otherUser.AvaID == nil {
		server.Logger.LogWarning(r, "User#"+BLUE+strconv.Itoa(otherUid)+NO_COLOR+" have no photos")
		server.error(w, errors.ImpossibleToExecute.WithArguments("Нельзя лайкать пользователей без фото",
			"Forbidden to like users without photo"))
		return
	}

	err = server.Db.SetNewLike(myUid, otherUid)
	if errors.ImpossibleToExecute.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "Imposible to set like from user#"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
			" to user#"+BLUE+strconv.Itoa(otherUid)+NO_COLOR)
		server.error(w, err.(errors.ApiError))
		return
	} else if errors.UserNotExist.IsOverlapWithError(err) {
		server.Logger.LogWarning(r, "User not exist - "+BLUE+err.Error()+NO_COLOR)
		server.error(w, errors.ImpossibleToExecute.WithArguments("Такого пользователя не существует", "User not exist"))
		return
	} else if err != nil {
		server.Logger.LogError(r, "SetNewLike returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	// Create notification to target user
	nid, err := server.Db.SetNewNotif(myUid, otherUid, myUser.Fname+" "+myUser.Lname+" liked you")
	if err != nil {
		server.Logger.LogError(r, "SetNewNotif returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}
	if server.Session.IsUserLoggedByUid(otherUid) {
		err = server.Session.SendNotifToLoggedUser(nid, otherUid, myUid, myUser.Fname+" "+myUser.Lname+" liked you")
		if err != nil {
			server.Logger.LogError(r, "SendNotifToLoggedUser returned error - "+err.Error())
			server.error(w, errors.UnknownInternalError)
			return
		}
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "Like was set successfully from user #"+BLUE+strconv.Itoa(myUid)+NO_COLOR+
		" to user #"+BLUE+strconv.Itoa(otherUid)+NO_COLOR)
}
