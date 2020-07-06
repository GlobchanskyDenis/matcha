package session

import (
	"testing"
	"fmt"
)

func TestCreateSession(t *testing.T) {
	sess := CreateSession()
	if (sess.session == nil) || sess.mu == nil {
		t.Errorf("\033[31mFAILED\033[m - empty session\n\n")
	} else {
		t.Logf("\033[32mDONE\033[m\n\n")
	}
}

func TestAddUser_1(t *testing.T) {
	var login = "admin"
	var id = 1

	sess := CreateSession()
	token := sess.AddUserToSession(login, id)
	user, err := sess.FindUserByToken(token)

	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s\n\n", fmt.Sprintf("%s", err))
	} else if user.Id != id || user.Login != login {
		t.Errorf("\033[31mFAILED\033[m - not equal %s %s %d %d\n\n", login, user.Login, id, user.Id)
	} else {
		t.Logf("\033[32mDONE\033[m\n\n")
	}
}

func TestAddUser_2(t *testing.T) {
	var login1 = "admin"
	var login2 = "bsabre"
	var id1 = 1
	var id2 = 23

	sess := CreateSession()
	token1 := sess.AddUserToSession(login1, id1)
	token2 := sess.AddUserToSession(login2, id2)
	user1, err := sess.FindUserByToken(token1)

	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s\n\n", fmt.Sprintf("%s", err))
		return
	}

	user2, err := sess.FindUserByToken(token2)

	if err != nil {
		t.Errorf("\033[31mFAILED\033[m - %s\n\n", fmt.Sprintf("%s", err))
	} else if user1.Id != id1 || user1.Login != login1 {
		t.Errorf("\033[31mFAILED\033[m - not equal %s %s %d %d\n\n", login1, user1.Login, id1, user1.Id)
	} else if user2.Id != id2 || user2.Login != login2 {
		t.Errorf("\033[31mFAILED\033[m - not equal %s %s %d %d\n\n", login2, user2.Login, id2, user2.Id)
	} else {
		t.Logf("\033[32mDONE\033[m\n\n")
	}
}

func TestAddUserInvalid_1(t *testing.T) {
	sess := CreateSession()
	_, err := sess.FindUserByToken("0000000")

	if err == nil {
		t.Errorf("\033[31mFAILED\033[m - expected error. Something goes wrong\n\n")
	} else {
		t.Logf("\033[32mDONE\033[m - %s - as it expected\n\n", fmt.Sprintf("%s", err))
	}
}