package postgres

import (
	. "MatchaServer/common"
	"MatchaServer/config"
	"MatchaServer/errDef"
	"MatchaServer/handlers"
	"testing"
)

var (
	mail = "test@gmail.com"
	pass = "AsdVar34!A"

	mailNew = "test_new@gmail.com"
	passNew = "DFe2*FDsd"

	mailFail = "mail@gmail@yandex.ru"
	passFail = "12345678"
)

func TestUser(t *testing.T) {
	print(NO_COLOR)

	///////// INITIALIZE //////////

	conn := New()
	conf, err := config.Create("../../config/")
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot get config file - " + err.Error() + NO_COLOR)
		return
	}
	err = conn.Connect(&conf.Sql)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot connect to database - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: connection with database" + NO_COLOR)

	///////// USER CREATE //////////

	user, err := conn.SetNewUser(mail, pass)
	if err != nil {
		t.Errorf(RED_BG + "ERROR: Cannot create user - " + err.Error() + NO_COLOR)
		return
	}
	t.Log(GREEN_BG + "SUCCESS: user created" + NO_COLOR)
	defer conn.DeleteUser(user.Uid)

	user.Pass = pass
	user.EncryptedPass = handlers.PassHash(pass)
	user.Status = "confirmed"

	///////// TESTS //////////

	t.Run("valid update", func(t_ *testing.T) {
		err = conn.UpdateUser(user)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: Cannot update user - " + err.Error() + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: user updated" + NO_COLOR)
		}
	})

	t.Run("invalid GetUserByUid", func(t_ *testing.T) {
		_, err := conn.GetUserByUid(0)
		if !errDef.RecordNotFound.IsOverlapWithError(err) {
			t_.Errorf(RED_BG + "ERROR: it should be Record not found error but it dont" + NO_COLOR)
		} else if err == nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: record not found as it expected" + NO_COLOR)
		}
	})

	t.Run("valid GetUserByUid", func(t_ *testing.T) {
		tempUser, err := conn.GetUserByUid(user.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else if tempUser.Mail != user.Mail {
			t_.Errorf(RED_BG + "ERROR: returned wrong user" + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: GetUserByUid is fine" + NO_COLOR)
		}
	})

	t.Run("invalid GetUserByMail", func(t_ *testing.T) {
		_, err := conn.GetUserByMail(mailFail)
		if !errDef.RecordNotFound.IsOverlapWithError(err) {
			t_.Errorf(RED_BG + "ERROR: it should be Record not found error but it dont" + NO_COLOR)
		} else if err == nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: record not found as it expected" + NO_COLOR)
		}
	})

	t.Run("valid GetUserByMail", func(t_ *testing.T) {
		tempUser, err := conn.GetUserByMail(user.Mail)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else if tempUser.Mail != user.Mail {
			t_.Errorf(RED_BG + "ERROR: returned wrong user" + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: GetUserByMail is fine" + NO_COLOR)
		}
	})

	t.Run("invalid GetUserForAuth", func(t_ *testing.T) {
		_, err := conn.GetUserForAuth(user.Mail, passFail)
		if !errDef.RecordNotFound.IsOverlapWithError(err) {
			t_.Errorf(RED_BG + "ERROR: it should be Record not found error but it dont" + NO_COLOR)
		} else if err == nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: record not found as it expected" + NO_COLOR)
		}
	})

	t.Run("valid GetUserForAuth", func(t_ *testing.T) {
		tempUser, err := conn.GetUserForAuth(user.Mail, user.EncryptedPass)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else if tempUser.Mail != user.Mail {
			t_.Errorf(RED_BG + "ERROR: returned wrong user" + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: GetUserForAuth is fine" + NO_COLOR)
		}
	})

	t.Run("invalid GetLoggedUsers", func(t_ *testing.T) {
		users, err := conn.GetLoggedUsers([]int{0, 0})
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else if len(users) != 0 {
			t_.Errorf(RED_BG + "ERROR: GetLoggedUsers unexpected returned users" + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: no users returned as it expected" + NO_COLOR)
		}
	})

	t.Run("valid GetLoggedUsers", func(t_ *testing.T) {
		users, err := conn.GetLoggedUsers([]int{user.Uid})
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else if len(users) != 1 {
			t_.Errorf(RED_BG + "ERROR: GetLoggedUsers returned wrong number of users" + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: GetLoggedUsers is fine" + NO_COLOR)
		}
	})

	t.Run("invalid IsUserExistsByMail", func(t_ *testing.T) {
		isExists, err := conn.IsUserExistsByMail(mailFail)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else if isExists {
			t_.Errorf(RED_BG + "ERROR: IsUserExistsByMail returned true" + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: IsUserExistsByMail is fine" + NO_COLOR)
		}
	})

	t.Run("valid IsUserExistsByMail", func(t_ *testing.T) {
		isExists, err := conn.IsUserExistsByMail(user.Mail)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else if !isExists {
			t_.Errorf(RED_BG + "ERROR: IsUserExistsByMail returned false" + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: IsUserExistsByMail is fine" + NO_COLOR)
		}
	})

	t.Run("invalid IsUserExistsByUid", func(t_ *testing.T) {
		isExists, err := conn.IsUserExistsByUid(0)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else if isExists {
			t_.Errorf(RED_BG + "ERROR: IsUserExistsByUid returned true" + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: IsUserExistsByUid is fine" + NO_COLOR)
		}
	})

	t.Run("valid IsUserExistsByUid", func(t_ *testing.T) {
		isExists, err := conn.IsUserExistsByUid(user.Uid)
		if err != nil {
			t_.Errorf(RED_BG + "ERROR: " + err.Error() + NO_COLOR)
		} else if !isExists {
			t_.Errorf(RED_BG + "ERROR: IsUserExistsByUid returned false" + NO_COLOR)
		} else {
			t_.Log(GREEN_BG + "SUCCESS: IsUserExistsByUid is fine" + NO_COLOR)
		}
	})
	print(YELLOW)
}
