package fakeSql

import (
	"MatchaServer/config"
)

func (conn ConnFake) SetNewUser(mail string, encryptedPass string) (config.User, error) {
	var user config.User

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

func (conn *ConnFake) UpdateUser(user config.User) error {
	conn.users[user.Uid] = user
	return nil
}

// !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
func (conn ConnFake) SearchUsersByOneFilter(filter string) ([]config.User, error) {
	return nil, nil
}

func (conn *ConnFake) GetUserByUid(uid int) (config.User, error) {
	user, _ := conn.users[uid]
	return user, nil
}

func (conn *ConnFake) GetUserByMail(mail string) (config.User, error) {
	for _, user := range conn.users {
		if user.Mail == mail {
			return user, nil
		}
	}
	return config.User{}, nil
}

func (conn *ConnFake) GetUserForAuth(mail string, encryptedPass string) (config.User, error) {
	for _, user := range conn.users {
		if user.Mail == mail && user.EncryptedPass == encryptedPass {
			return user, nil
		}
	}
	return config.User{}, nil
}

func (conn *ConnFake) GetLoggedUsers(uid []int) ([]config.User, error) {
	var users = []config.User{}
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
