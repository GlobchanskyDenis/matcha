package fakeSql

import (
	"MatchaServer/common"
	"testing"
)

func TestUser(t *testing.T) {
	print(common.NO_COLOR)
	var repo = ConnFake{}
	var mail = "mail@mail.ru"
	var encryptedPass = "EncryptedPass"
	var mailNew = "email@mail.ru"
	var encryptedPassNew = "NewEncryptedPass"

	_ = repo.Connect()

	user1, _ := repo.SetNewUser(mail, encryptedPass)
	user2, _ := repo.SetNewUser(mailNew, encryptedPassNew)
	user1.Fname = "Denis"
	user1.Lname = "Globchansky"
	_ = repo.UpdateUser(user1)
	userTmp, _ := repo.GetUserByUid(user1.Uid)
	if userTmp.Uid != user1.Uid || userTmp.Fname != user1.Fname || userTmp.Lname != user1.Lname {
		t.Error(common.RED_BG + "ERROR: GetUserByUid" + common.NO_COLOR + "\n")
		return
	}
	userTmp, _ = repo.GetUserByMail(user1.Mail)
	if userTmp.Uid != user1.Uid || userTmp.Fname != user1.Fname || userTmp.Lname != user1.Lname {
		t.Error(common.RED_BG + "ERROR: GetUserByMail" + common.NO_COLOR + "\n")
		return
	}
	userTmp, _ = repo.GetUserForAuth(mail, encryptedPass)
	if userTmp.Uid != user1.Uid || userTmp.Fname != user1.Fname || userTmp.Lname != user1.Lname {
		t.Error(common.RED_BG + "ERROR: GetUserForAuth" + common.NO_COLOR + "\n")
		return
	}
	users, _ := repo.GetLoggedUsers([]int{user2.Uid})
	if users == nil || len(users) != 1 || users[0].Uid != user2.Uid || users[0].Fname != user2.Fname || users[0].Lname != user2.Lname {
		t.Error(common.RED_BG + "ERROR: GetLoggedUsers" + common.NO_COLOR + "\n")
		return
	}
	was, _ := repo.IsUserExistsByMail(user1.Mail)
	_ = repo.DeleteUser(user1.Uid)
	now, _ := repo.IsUserExistsByMail(user1.Mail)
	if was != true || now != false {
		t.Error(common.RED_BG + "ERROR: IsUserExistsByMail + DeleteUser" + common.NO_COLOR + "\n")
		return
	}
	was, _ = repo.IsUserExistsByUid(user2.Uid)
	_ = repo.DeleteUser(user2.Uid)
	now, _ = repo.IsUserExistsByUid(user2.Uid)
	if was != true || now != false {
		t.Error(common.RED_BG + "ERROR: IsUserExistsByUid + DeleteUser" + common.NO_COLOR + "\n")
		return
	}
	t.Log(common.GREEN_BG + "SUCCESS" + common.NO_COLOR + "\n")
	print(common.YELLOW)
}
