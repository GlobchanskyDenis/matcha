package fakeSql

import (
	"MatchaServer/common"
	"MatchaServer/errors"
)

func (conn *ConnFake) SetNewLike(uidSender int, uidReceiver int) error {
	var newLike = relationMock{uidSender: uidSender, uidReceiver: uidReceiver}

	if !conn.isUserExists(uidSender) {
		return errors.UserNotExist
	}
	if !conn.isUserExists(uidReceiver) {
		return errors.UserNotExist
	}
	if isSliceContainsElement(conn.likes, newLike) {
		return errors.ImpossibleToExecute.WithArguments("Вы уже игнорируете этого пользователя",
			"You are already ignoring this user")
	}
	conn.likes = append(conn.likes, newLike)
	conn.changeUserRating(uidReceiver, 1)
	return nil
}

func (conn *ConnFake) UnsetLike(uidSender int, uidReceiver int) error {
	if !conn.isUserExists(uidSender) {
		return errors.UserNotExist
	}
	if !conn.isUserExists(uidReceiver) {
		return errors.UserNotExist
	}
	for i, like := range conn.likes {
		if like.uidSender == uidSender && like.uidReceiver == uidReceiver {
			// Исключаю элемент номер i
			conn.likes = append(conn.likes[:i], conn.likes[i+1:]...)
			conn.changeUserRating(uidReceiver, -1)
			return nil
		}
	}
	return errors.ImpossibleToExecute.WithArguments("Вы не игнорируете этого пользователя",
		"You are not ignoring this user")
}

func (conn *ConnFake) DropUserLikes(uid int) error {
LOOP:
	for i, like := range conn.likes {
		if like.uidSender == uid || like.uidReceiver == uid {
			// Исключаю элемент номер i
			conn.likes = append(conn.likes[:i], conn.likes[i+1:]...)
			break LOOP
		}
	}
	return nil
}

func (conn ConnFake) GetFriendUsers(uidSender int) ([]common.FriendUser, error) {
	var users []common.FriendUser
	for _, like := range conn.likes {
		if like.uidSender == uidSender {
			for _, user := range conn.users {
				if user.Uid == like.uidReceiver {
					if user.AvaID != 0 {
						photo, _ := conn.GetPhotoByPid(user.AvaID)
						user.Avatar = &photo.Src
					}
					var friendUser = common.FriendUser{User: user}
					// тут нужно присоединить последнее письмо между пользователями. Тестов нет
					for i := len(conn.messages) - 1; i >= 0; i-- {
						message := conn.messages[i]
						if (message.UidSender == uidSender && message.UidReceiver == like.uidReceiver) ||
							(message.UidSender == like.uidReceiver && message.UidReceiver == uidSender) {
							friendUser.UidSender = &message.UidSender
							friendUser.UidReceiver = &message.UidReceiver
							friendUser.LastMessageBody = &message.Body
							break
						}
					}
					users = append(users, friendUser)
					break
				}
			}
		}
	}
	return users, nil
}

func (conn ConnFake) IsICanSpeakWithUser(myUid, otherUid int) (bool, error) {
	for _, like := range conn.likes {
		if like.uidSender == myUid && like.uidReceiver == otherUid {
			for _, like := range conn.likes {
				if like.uidSender == otherUid && like.uidReceiver == myUid {
					return true, nil
				}
			}
		}
	}
	return false, nil
}
