package session

import (
	"testing"
	. "MatchaServer/config"
)

func TestCreateSession(t *testing.T) {
	sess := CreateSession()
	if (sess.session == nil) || sess.mu == nil {
		t.Errorf(RED_BG + "FAILED: empty session" + NO_COLOR + "\n")
	} else {
		t.Logf(GREEN_BG + "SUCCESS" + NO_COLOR + "\n")
	}
}

func TestAddUser_1(t *testing.T) {
	var login = "admin"
	var passwd = "adsdasdsadsad"
	var mail = "mail@gmail.com"
	var id = 1

	sess := CreateSession()
	token, err := sess.AddUserToSession(id, login, passwd, mail)
	if err!= nil {
		t.Errorf(RED_BG + "FAILED: %s" + NO_COLOR + "\n", err.Error())
		return
	}
	user, err := sess.FindUserByToken(token)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s" + NO_COLOR + "\n", err.Error())
		return
	}
	if user.UserInfo.Id != id {
		t.Errorf(RED_BG + "FAILED: not equal id %d %d" + NO_COLOR + "\n", id, user.UserInfo.Id)
		return
	}
	if user.UserInfo.Login != login {
		t.Errorf(RED_BG + "FAILED: not equal login %s %s" + NO_COLOR + "\n", login, user.UserInfo.Login)
		return
	}
	if user.UserInfo.Passwd != passwd {
		t.Errorf(RED_BG + "FAILED: not equal passwd %s %s" + NO_COLOR + "\n", passwd, user.UserInfo.Passwd)
		return
	}
	if user.UserInfo.Mail != mail {
		t.Errorf(RED_BG + "FAILED: not equal mail %s %s" + NO_COLOR + "\n", mail, user.UserInfo.Mail)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS" + NO_COLOR + "\n")
}

func TestAddUser_2(t *testing.T) {
	var login1 = "admin"
	var login2 = "bsabre"
	var passwd1 = "asdassrsda"
	var mail1 = "mail1@gmail.com"
	var passwd2 = "aasddasw3wwv"
	var mail2 = "mail2@gmail.com"
	var id1 = 1
	var id2 = 23

	sess := CreateSession()
	token1, err := sess.AddUserToSession(id1, login1, passwd1, mail1)
	if err!= nil {
		t.Errorf(RED_BG + "FAILED: %s" + NO_COLOR + "\n", err.Error())
		return
	}
	token2, err := sess.AddUserToSession(id2, login2, passwd2, mail2)
	if err!= nil {
		t.Errorf(RED_BG + "FAILED: %s" + NO_COLOR + "\n", err.Error())
		return
	}
	user1, err := sess.FindUserByToken(token1)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s" + NO_COLOR + "\n", err.Error())
		return
	}
	user2, err := sess.FindUserByToken(token2)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s" + NO_COLOR + "\n", err.Error())
		return
	}

	if user1.UserInfo.Id != id1 {
		t.Errorf(RED_BG + "FAILED: not equal id1 %d %d" + NO_COLOR + "\n", id1, user1.UserInfo.Id)
		return
	}
	if user1.UserInfo.Login != login1 {
		t.Errorf(RED_BG + "FAILED: not equal login1 %s %s" + NO_COLOR + "\n", login1, user1.UserInfo.Login)
		return
	}
	if user1.UserInfo.Passwd != passwd1 {
		t.Errorf(RED_BG + "FAILED: not equal passwd1 %s %s" + NO_COLOR + "\n", passwd1, user1.UserInfo.Passwd)
		return
	}
	if user1.UserInfo.Mail != mail1 {
		t.Errorf(RED_BG + "FAILED: not equal mail1 %s %s" + NO_COLOR + "\n", mail1, user1.UserInfo.Mail)
		return
	}
	if user2.UserInfo.Id != id2 {
		t.Errorf(RED_BG + "FAILED: not equal id %d %d" + NO_COLOR + "\n", id2, user2.UserInfo.Id)
		return
	}
	if user2.UserInfo.Login != login2 {
		t.Errorf(RED_BG + "FAILED: not equal login %s %s" + NO_COLOR + "\n", login2, user2.UserInfo.Login)
		return
	}
	if user2.UserInfo.Passwd != passwd2 {
		t.Errorf(RED_BG + "FAILED: not equal passwd %s %s" + NO_COLOR + "\n", passwd2, user2.UserInfo.Passwd)
		return
	}
	if user2.UserInfo.Mail != mail2 {
		t.Errorf(RED_BG + "FAILED: not equal mail %s %s" + NO_COLOR + "\n", mail2, user2.UserInfo.Mail)
		return
	}

	t.Logf(GREEN_BG + "SUCCESS" + NO_COLOR + "\n")
}

func TestAddUserInvalid_1(t *testing.T) {
	sess := CreateSession()
	_, err := sess.FindUserByToken("0000000")

	if err == nil {
		t.Errorf(RED_BG + "FAILED: expected error. Something goes wrong" + NO_COLOR + "\n")
	} else {
		t.Logf(GREEN_BG + "SUCCESS: %s - as it expected" + NO_COLOR + "\n", err.Error())
	}
}