package fakeSql

import (
	"MatchaServer/common"
	"MatchaServer/errors"
)

func (conn *ConnFake) SetNewIgnore(uidSender int, uidReceiver int) error {
	var newIgnore = relationMock{uidSender: uidSender, uidReceiver: uidReceiver}

	if !conn.isUserExists(uidSender) {
		return errors.UserNotExist
	}
	if !conn.isUserExists(uidReceiver) {
		return errors.UserNotExist
	}
	if isSliceContainsElement(conn.ignores, newIgnore) {
		return errors.ImpossibleToExecute.WithArguments("Вы уже игнорируете этого пользователя",
			"You are already ignoring this user")
	}
	conn.ignores = append(conn.ignores, newIgnore)
	conn.changeUserRating(uidReceiver, 3)
	return nil
}

func (conn *ConnFake) UnsetIgnore(uidSender int, uidReceiver int) error {
	if !conn.isUserExists(uidSender) {
		return errors.UserNotExist
	}
	if !conn.isUserExists(uidReceiver) {
		return errors.UserNotExist
	}
	for i, ignore := range conn.ignores {
		if ignore.uidSender == uidSender && ignore.uidReceiver == uidReceiver {
			// Исключаю элемент номер i
			conn.ignores = append(conn.ignores[:i], conn.ignores[i+1:]...)
			conn.changeUserRating(uidReceiver, -3)
			return nil
		}
	}
	return errors.ImpossibleToExecute.WithArguments("Вы не игнорируете этого пользователя",
		"You are not ignoring this user")
}

func (conn *ConnFake) DropUserIgnores(uid int) error {
LOOP:
	for i, ignore := range conn.ignores {
		if ignore.uidSender == uid || ignore.uidReceiver == uid {
			// Исключаю элемент номер i
			conn.ignores = append(conn.ignores[:i], conn.ignores[i+1:]...)
			break LOOP
		}
	}
	return nil
}

func (conn ConnFake) GetIgnoredUsers(uidSender int) ([]common.User, error) {
	var users []common.User
	for _, ignore := range conn.ignores {
		if ignore.uidSender == uidSender {
			for _, user := range conn.users {
				if user.Uid == ignore.uidReceiver {
					if user.AvaID != 0 {
						photo, _ := conn.GetPhotoByPid(user.AvaID)
						user.Avatar = &photo.Src
					}
					users = append(users, user)
					break
				}
			}
		}
	}
	return users, nil
}
