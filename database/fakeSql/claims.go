package fakeSql

import (
	"MatchaServer/common"
	"MatchaServer/errors"
)

func (conn *ConnFake) SetNewClaim(uidSender int, uidReceiver int) error {
	var newClaim = relationMock{uidSender: uidSender, uidReceiver: uidReceiver}

	if !conn.isUserExists(uidSender) {
		return errors.UserNotExist
	}
	if !conn.isUserExists(uidReceiver) {
		return errors.UserNotExist
	}
	if isSliceContainsElement(conn.claims, newClaim) {
		return errors.ImpossibleToExecute.WithArguments("Вы уже жаловались на этого пользователя",
			"You have already reported this user")
	}
	conn.claims = append(conn.claims, newClaim)
	return nil
}

func (conn *ConnFake) UnsetClaim(uidSender int, uidReceiver int) error {
	if !conn.isUserExists(uidSender) {
		return errors.UserNotExist
	}
	if !conn.isUserExists(uidReceiver) {
		return errors.UserNotExist
	}
	for i, claim := range conn.claims {
		if claim.uidSender == uidSender && claim.uidReceiver == uidReceiver {
			// Исключаю элемент номер i
			conn.claims = append(conn.claims[:i], conn.claims[i+1:]...)
			return nil
		}
	}
	return errors.ImpossibleToExecute.WithArguments("Вы не жаловались на этого пользователя",
		"You have not reported this user")
}

func (conn *ConnFake) DropUserClaims(uid int) error {
LOOP:
	for i, claim := range conn.claims {
		if claim.uidSender == uid || claim.uidReceiver == uid {
			// Исключаю элемент номер i
			conn.claims = append(conn.claims[:i], conn.claims[i+1:]...)
			break LOOP
		}
	}
	return nil
}

func (conn ConnFake) GetClaimedUsers(uidSender int) ([]common.User, error) {
	var users []common.User
	for _, claim := range conn.claims {
		if claim.uidSender == uidSender {
			for _, user := range conn.users {
				if user.Uid == claim.uidReceiver {
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
