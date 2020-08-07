package fakeSql

import (
	"MatchaServer/config"
	"testing"
)

// const (
// 	RED       = "\033[31m"
// 	GREEN     = "\033[32m"
// 	YELLOW    = "\033[33m"
// 	BLUE      = "\033[34m"
// 	RED_BG    = "\033[41;30m"
// 	GREEN_BG  = "\033[42;30m"
// 	YELLOW_BG = "\033[43;30m"
// 	BLUE_BG   = "\033[44;30m"
// 	NO_COLOR  = "\033[m"
// )

func TestTruncate(t *testing.T) {
	print(config.NO_COLOR)
	var repo = ConnFake{}

	_ = repo.Connect()
	user1, _ := repo.SetNewUser("mail@mail.ru", "EncryptedPass")
	user2, _ := repo.SetNewUser("new@mail.ru", "EncryptedPass")

	if user1.Uid == user2.Uid {
		t.Errorf(config.RED_BG + "ERROR: Uid is invalid " + config.NO_COLOR + "\n")
	}

	repo.TruncateAllTables()

	if _, isExists := repo.users[1]; !isExists {
		t.Log(config.GREEN_BG + "SUCCESS: user #1 was deleted" + config.NO_COLOR + "\n")
	}
	if _, isExists := repo.users[2]; !isExists {
		t.Log(config.GREEN_BG + "SUCCESS: user #2 was deleted" + config.NO_COLOR + "\n")
	}
	print(config.YELLOW)
}
