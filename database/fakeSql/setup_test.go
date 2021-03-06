package fakeSql

import (
	"MatchaServer/common"
	"MatchaServer/config"
	"testing"
)

func TestTruncate(t *testing.T) {
	print(common.NO_COLOR)
	var repo = ConnFake{}

	conf, err := config.Create("../../config/")
	if err != nil {
		t.Errorf(common.RED_BG + "ERROR: Cannot get config file - " + err.Error() + common.NO_COLOR)
		return
	}
	err = repo.Connect(&conf.Sql)
	if err != nil {
		t.Errorf(common.RED_BG + "ERROR: Cannot connect to database - " + err.Error() + common.NO_COLOR)
		return
	}
	user1, _ := repo.SetNewUser("mail@mail.ru", "EncryptedPass")
	user2, _ := repo.SetNewUser("new@mail.ru", "EncryptedPass")

	if user1.Uid == user2.Uid {
		t.Errorf(common.RED_BG + "ERROR: Uid is invalid " + common.NO_COLOR + "\n")
	}

	repo.TruncateAllTables()

	if _, isExists := repo.users[1]; !isExists {
		t.Log(common.GREEN_BG + "SUCCESS: user #1 was deleted" + common.NO_COLOR + "\n")
	}
	if _, isExists := repo.users[2]; !isExists {
		t.Log(common.GREEN_BG + "SUCCESS: user #2 was deleted" + common.NO_COLOR + "\n")
	}
	print(common.YELLOW)
}
