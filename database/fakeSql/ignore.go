package fakeSql

import (
	"MatchaServer/common"
)

func (conn ConnFake) SetNewIgnore(uidSender int, uidReceiver int) error {
	return nil
}

func (conn ConnFake) UnsetIgnore(uidSender int, uidReceiver int) error {
	return nil
}

func (conn ConnFake) GetIgnoredUsers(uidSender int) ([]common.User, error) {
	return nil, nil
}
