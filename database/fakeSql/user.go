package fakeSql

import (
	"MatchaServer/common"
	"MatchaServer/errors"
)

func (conn *ConnFake) SetNewUser(mail string, encryptedPass string) (common.User, error) {
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

func (conn *ConnFake) GetUserByUid(uid int) (common.User, error) {
	user, isExists := conn.users[uid]
	if !isExists {
		return user, errors.RecordNotFound
	}
	if user.AvaID != nil {
		photo, _ := conn.GetPhotoByPid(*user.AvaID)
		user.Avatar = &photo.Src
	}
	return user, nil
}

func (conn *ConnFake) GetTargetUserByUid(myUid int, targetUid int) (common.TargetUser, error) {
	user, isExists := conn.users[myUid]
	if !isExists {
		return common.TargetUser{}, errors.RecordNotFound
	}
	if user.AvaID != nil {
		photo, _ := conn.GetPhotoByPid(*user.AvaID)
		user.Avatar = &photo.Src
	}
	return common.TargetUser{User: user}, nil
}

func (conn *ConnFake) GetUserByMail(mail string) (common.User, error) {
	for _, user := range conn.users {
		if user.Mail == mail {
			if user.AvaID != nil {
				photo, _ := conn.GetPhotoByPid(*user.AvaID)
				user.Avatar = &photo.Src
			}
			return user, nil
		}
	}
	return common.User{}, errors.RecordNotFound
}

/*
**	Данная функция возвращает абсолютно всех пользователей. Парсить запрос в мок объекте -
**	неблагодарное дело
 */
func (conn *ConnFake) GetUsersByQuery(query string, sourceUser common.User) ([]common.SearchUser, error) {
	var users []common.SearchUser

	for _, user := range conn.users {
		if sourceUser.Uid == user.Uid {
			continue
		}
		if user.AvaID != nil {
			photo, _ := conn.GetPhotoByPid(*user.AvaID)
			user.Avatar = &photo.Src
		}
		searchUser := common.SearchUser{User: user}
		users = append(users, searchUser)
	}
	return users, nil
}

func (conn *ConnFake) GetUserForAuth(mail string, encryptedPass string) (common.User, error) {
	for _, user := range conn.users {
		if user.Mail == mail && user.EncryptedPass == encryptedPass {
			if user.AvaID != nil {
				photo, _ := conn.GetPhotoByPid(*user.AvaID)
				user.Avatar = &photo.Src
			}
			return user, nil
		}
	}
	return common.User{}, errors.RecordNotFound
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

func (conn ConnFake) GetUserWithLikeInfo(targetUid int, myUid int) (common.SearchUser, error) {
	return common.SearchUser{}, nil
}
