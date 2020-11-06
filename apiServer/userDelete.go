package apiServer

import (
	. "MatchaServer/common"
	"MatchaServer/errors"
	"MatchaServer/handlers"
	"context"
	"net/http"
	"strconv"
)

// HTTP HANDLER FOR DOMAIN /user/delete/
func (server *Server) UserDelete(w http.ResponseWriter, r *http.Request) {
	var (
		err                 error
		requestParams       map[string]interface{}
		item                interface{}
		ctx                 context.Context
		pass, encryptedPass string
		uid                 int
		isExist, ok         bool
	)

	ctx = r.Context()
	requestParams = ctx.Value("requestParams").(map[string]interface{})
	uid = ctx.Value("uid").(int)

	item, isExist = requestParams["pass"]
	if !isExist {
		server.Logger.LogWarning(r, "password not exist")
		server.error(w, errors.NoArgument.WithArguments("Поле pass отсутствует", "pass field expected"))
		return
	}

	pass, ok = item.(string)
	if !ok {
		server.Logger.LogWarning(r, "password have wrong type")
		server.error(w, errors.InvalidArgument.WithArguments("Поле pass имеет неверный тип", "pass field has wrong type"))
		return
	}

	encryptedPass = handlers.PassHash(pass)

	user, err := server.Db.GetUserByUid(uid)
	if err != nil {
		server.Logger.LogError(r, "GetUserByUid returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	if encryptedPass != user.EncryptedPass {
		server.Logger.LogWarning(r, "password is incorrect")
		server.error(w, errors.InvalidArgument.WithArguments("неверный пароль", "password is wrong"))
		return
	}

	//	Delete devices of user before user
	devices, err := server.Db.GetDevicesByUid(user.Uid)
	if err != nil {
		server.Logger.LogError(r, "GetDevicesByUid returned error - "+err.Error())
		server.error(w, errors.DatabaseError.WithArguments(err))
		return
	}
	for _, device := range devices {
		err = server.Db.DeleteDevice(device.Id)
		if err != nil {
			server.Logger.LogError(r, "Cannot delete user device - "+err.Error())
			server.error(w, errors.DatabaseError)
			return
		}
	}

	//	Delete notifs of user before user
	notifs, err := server.Db.GetNotifByUidReceiver(user.Uid)
	if err != nil {
		server.Logger.LogError(r, "GetNotifByUidReceiver returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}
	for _, notif := range notifs {
		err = server.Db.DeleteNotif(notif.Nid)
		if err != nil {
			server.Logger.LogError(r, "Cannot delete user notifications - "+err.Error())
			server.error(w, errors.DatabaseError)
			return
		}
	}

	//	Delete photos of user before user
	photos, err := server.Db.GetPhotosByUid(user.Uid)
	if err != nil {
		server.Logger.LogError(r, "GetPhotosByUid returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}
	for _, photo := range photos {
		err = server.Db.DeletePhoto(photo.Pid)
		if err != nil {
			server.Logger.LogError(r, "Cannot delete user photos - "+err.Error())
			server.error(w, errors.DatabaseError)
			return
		}
	}

	// Delete ignores of user before user
	err = server.Db.DropUserIgnores(user.Uid)
	if err != nil {
		server.Logger.LogError(r, "DropUserIgnores returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	// Delete claims of user before user
	err = server.Db.DropUserClaims(user.Uid)
	if err != nil {
		server.Logger.LogError(r, "DropUserClaimes returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	err = server.Db.DeleteUser(user.Uid)
	if err != nil {
		server.Logger.LogError(r, "DeleteUser returned error - "+err.Error())
		server.error(w, errors.DatabaseError)
		return
	}

	server.Session.DeleteUserSessionByUid(user.Uid)

	w.WriteHeader(http.StatusOK) // 200
	server.Logger.LogSuccess(r, "user #"+BLUE+strconv.Itoa(user.Uid)+NO_COLOR+
		" was removed successfully. No response body")
}
