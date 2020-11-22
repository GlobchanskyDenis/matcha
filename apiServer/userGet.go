package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /user/get/ . IT HANDLES:
// IT RETURNS TARGET USER DATA IN RESPONSE BY POST METHOD.
// IT MAKES RECORD IN USER HISTORY IF TARGET USER IS NOT YOU.
// REQUEST AND RESPONSE DATA IS JSON
func (server *Server) UserGet(w http.ResponseWriter, r *http.Request) {
	var (
		myUser, user    User
		otherUser       SearchUser
		myUid, otherUid int
		err             error
		ctx             context.Context
		requestParams   map[string]interface{}
		item            interface{}
		isExist, ok     bool
		uid64           float64
		jsonUser        []byte
	)

	ctx = r.Context()
	myUid = ctx.Value("uid").(int)
	requestParams = ctx.Value("requestParams").(map[string]interface{})
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

	if myUid != otherUid {
		// Its a private field
		user.Mail = ""

		otherUser, err = server.Db.GetUserWithLikeInfo(otherUid, myUid)
		if errors.RecordNotFound.IsOverlapWithError(err) {
			server.Logger.LogWarning(r, "GetUserWithLikeInfo - record not found")
			server.error(w, errors.UserNotExist)
			return
		} else if err != nil {
			server.Logger.LogError(r, "GetUser returned error - "+err.Error())
			server.error(w, errors.DatabaseError)
			return
		}

		// Make record in users history
		err = server.Db.SetNewHistoryReference(myUid, otherUid)
		if err != nil {
			server.Logger.LogError(r, "SetNewHistoryReference returned error - "+err.Error())
			server.error(w, errors.DatabaseError)
			return
		}

		myUser, err = server.Db.GetUserByUid(myUid)
		if errors.RecordNotFound.IsOverlapWithError(err) {
			server.Logger.LogWarning(r, "Your user#"+BLUE+strconv.Itoa(myUid)+NO_COLOR+" not exists")
			server.error(w, errors.ImpossibleToExecute.WithArguments("Вашего пользователя не существует", "Your user isnt exist"))
			return
		} else if err != nil {
			server.Logger.LogError(r, "GetUserByUid returned error - "+err.Error())
			server.error(w, errors.DatabaseError)
			return
		}

		// Create notification to target user
		nid, err := server.Db.SetNewNotif(myUid, otherUid, myUser.Fname+" "+myUser.Lname+" watched your account")
		if err != nil {
			server.Logger.LogError(r, "SetNewNotif returned error - "+err.Error())
			server.error(w, errors.DatabaseError)
			return
		}
		if server.Session.IsUserLoggedByUid(otherUid) {
			err = server.Session.SendNotifToLoggedUser(nid, otherUid, myUid, myUser.Fname+" "+myUser.Lname+" watched your account")
			if err != nil {
				server.Logger.LogError(r, "SendNotifToLoggedUser returned error - "+err.Error())
				server.error(w, errors.UnknownInternalError)
				return
			}
		}
		jsonUser, err = json.Marshal(otherUser)
		if err != nil {
			server.Logger.LogError(r, "Marshal returned error "+err.Error())
			server.error(w, errors.MarshalError)
			return
		}
	} else {
		user, err = server.Db.GetUserByUid(otherUid)
		if errors.RecordNotFound.IsOverlapWithError(err) {
			server.Logger.LogWarning(r, "GetUserByUid - record not found")
			server.error(w, errors.UserNotExist)
			return
		} else if err != nil {
			server.Logger.LogError(r, "GetUserByUid returned error - "+err.Error())
			server.error(w, errors.DatabaseError)
			return
		}

		jsonUser, err = json.Marshal(user)
		if err != nil {
			server.Logger.LogError(r, "Marshal returned error "+err.Error())
			server.error(w, errors.MarshalError)
			return
		}
	}

	// This is my valid case
	w.WriteHeader(http.StatusOK) // 200
	w.Write(jsonUser)
	server.Logger.LogSuccess(r, "User #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+" was found and transmitted successfully")
}
