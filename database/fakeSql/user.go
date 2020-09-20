package fakeSql

import (
	"MatchaServer/common"
	"MatchaServer/errDef"
)

func (conn ConnFake) SetNewUser(mail string, encryptedPass string) (common.User, error) {
	var user common.User

	user.Mail = mail
	user.EncryptedPass = encryptedPass

	for key := 1; ; key++ {
		if _, isExists := conn.users[key]; !isExists {
			user.Uid = key
			break
		}
	}

	conn.users[user.Uid] = user
	return user, nil
}

func (conn *ConnFake) DeleteUser(uid int) error {
	delete(conn.users, uid)
	return nil
}

func (conn *ConnFake) UpdateUser(user common.User) error {
	conn.users[user.Uid] = user
	return nil
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (conn ConnFake) SearchUsersByOneFilter(filter string) ([]common.User, error) {
	return nil, nil
}

func (conn *ConnFake) GetUserByUid(uid int) (common.User, error) {
	user, isExists := conn.users[uid]
	if !isExists {
		return user, errDef.RecordNotFound
	}
	return user, nil
}

func (conn *ConnFake) GetUserByMail(mail string) (common.User, error) {
	for _, user := range conn.users {
		if user.Mail == mail {
			return user, nil
		}
	}
	return common.User{}, errDef.RecordNotFound
}

func (conn *ConnFake) GetUsersByQuery(query string) ([]common.User, error) {
	return nil, nil
}

func (conn *ConnFake) GetUserForAuth(mail string, encryptedPass string) (common.User, error) {
	for _, user := range conn.users {
		if user.Mail == mail && user.EncryptedPass == encryptedPass {
			return user, nil
		}
	}
	return common.User{}, errDef.RecordNotFound
}

func (conn *ConnFake) GetLoggedUsers(uid []int) ([]common.User, error) {
	var users = []common.User{}
	for _, user := range conn.users {
		for _, id := range uid {
			if user.Uid == id {
				users = append(users, user)
			}
		}
	}
	return users, nil
}

func (conn ConnFake) IsUserExistsByMail(mail string) (bool, error) {
	for _, user := range conn.users {
		if user.Mail == mail {
			return true, nil
		}
	}
	return false, nil
}

func (conn ConnFake) IsUserExistsByUid(uid int) (bool, error) {
	if _, isExists := conn.users[uid]; isExists {
		return true, nil
	}
	return false, nil
}
