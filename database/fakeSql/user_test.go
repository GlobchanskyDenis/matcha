package fakeSql

import (
	"testing"
	// "fmt"
)

func TestUser(t *testing.T) {
	var repo = ConnFake{}
	var mail = "mail@mail.ru"
	var pass = "EncryptedPass"
	var mailNew = "email@mail.ru"
	var passNew = "NewEncryptedPass"

	_ = repo.Connect()

	user1, _ := repo.SetNewUser(mail, pass)
	user2, _ := repo.SetNewUser(mailNew, passNew)
	user1.Fname = "Denis"
	user1.Lname = "Globchansky"
	_ = repo.UpdateUser(user1)
	userTmp, _ := repo.GetUserByUid(user1.Uid)
	if userTmp != user1 {
		t.Error(RED_BG + "ERROR: GetUserByUid" + NO_COLOR + "\n")
		return
	}
	userTmp, _ = repo.GetUserByMail(user1.Mail)
	if userTmp != user1 {
		t.Error(RED_BG + "ERROR: GetUserByMail" + NO_COLOR + "\n")
		return
	}
	userTmp, _ = repo.GetUserForAuth(mail, pass)
	if userTmp != user1 {
		t.Error(RED_BG + "ERROR: GetUserForAuth" + NO_COLOR + "\n")
		return
	}
	users, _ := repo.GetLoggedUsers([]int{user2.Uid})
	if users == nil || len(users) != 1 || users[0] != user2 {
		t.Error(RED_BG + "ERROR: GetLoggedUsers" + NO_COLOR + "\n")
		return
	}
	was, _ := repo.IsUserExistsByMail(user1.Mail)
	_ = repo.DeleteUser(user1.Uid)
	now, _ := repo.IsUserExistsByMail(user1.Mail)
	if was != true || now != false {
		t.Error(RED_BG + "ERROR: IsUserExistsByMail + DeleteUser" + NO_COLOR + "\n")
		return
	}
	was, _ = repo.IsUserExistsByUid(user2.Uid)
	_ = repo.DeleteUser(user2.Uid)
	now, _ = repo.IsUserExistsByUid(user2.Uid)
	if was != true || now != false {
		t.Error(RED_BG + "ERROR: IsUserExistsByUid + DeleteUser" + NO_COLOR + "\n")
		return
	}
	t.Log(GREEN_BG + "SUCCESS" + NO_COLOR + "\n")
}
