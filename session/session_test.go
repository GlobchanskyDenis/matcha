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
	var uid = 1

	sess := CreateSession()
	token, err := sess.AddUserToSession(uid)
	if err!= nil {
		t.Errorf(RED_BG + "FAILED: %s" + NO_COLOR + "\n", err.Error())
		return
	}
	user, err := sess.FindUserByToken(token)
	if err != nil {
		t.Errorf(RED_BG + "FAILED: %s" + NO_COLOR + "\n", err.Error())
		return
	}
	if user.Uid != uid {
		t.Errorf(RED_BG + "FAILED: not equal id %d %d" + NO_COLOR + "\n", uid, user.Uid)
		return
	}
	t.Logf(GREEN_BG + "SUCCESS" + NO_COLOR + "\n")
}

func TestAddUser_2(t *testing.T) {
	var uid1 = 1
	var uid2 = 5

	sess := CreateSession()
	token1, err := sess.AddUserToSession(uid1)
	if err!= nil {
		t.Errorf(RED_BG + "FAILED: %s" + NO_COLOR + "\n", err.Error())
		return
	}
	token2, err := sess.AddUserToSession(uid2)
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

	if user1.Uid != uid1 {
		t.Errorf(RED_BG + "FAILED: not equal id1 %d %d" + NO_COLOR + "\n", uid1, user1.Uid)
		return
	}
	if user2.Uid != uid2 {
		t.Errorf(RED_BG + "FAILED: not equal id %d %d" + NO_COLOR + "\n", uid2, user2.Uid)
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