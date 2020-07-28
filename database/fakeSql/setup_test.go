package fakeSql

import (
	// "MatchaServer/config"
	"testing"
)

const (
	RED       = "\033[31m"
	GREEN     = "\033[32m"
	YELLOW    = "\033[33m"
	BLUE      = "\033[34m"
	RED_BG    = "\033[41;30m"
	GREEN_BG  = "\033[42;30m"
	YELLOW_BG = "\033[43;30m"
	BLUE_BG   = "\033[44;30m"
	NO_COLOR  = "\033[m"
)

func TestTruncate(t *testing.T) {
	var repo = ConnFake{}

	_ = repo.Connect()
	user1, _ := repo.SetNewUser("mail@mail.ru", "EncryptedPass")
	user2, _ := repo.SetNewUser("new@mail.ru", "EncryptedPass")

	println(user1.Uid)
	println(user2.Uid)
	
	if user1.Uid == user2.Uid {
		t.Errorf(RED_BG + "ERROR: Uid is invalid " + NO_COLOR + "\n")
	}

	repo.TruncateAllTables()

	// users, _ := repo.GetLoggedUsers([]int{user1.Uid, user1.Uid})

	// if users == nil || len(users) != 0 {
	// 	t.Errorf(RED_BG + "ERROR: after truncate repo is not empty " + NO_COLOR + "\n")
	// }

	if _, isExists := repo.users[1]; !isExists {
		t.Log(GREEN_BG + "SUCCESS: user #1 was deleted" + NO_COLOR + "\n")
	}
	if _, isExists := repo.users[2]; !isExists {
		t.Log(GREEN_BG + "SUCCESS: user #2 was deleted" + NO_COLOR + "\n")
	}
	print(YELLOW)
}