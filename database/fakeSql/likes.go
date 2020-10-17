package fakeSql

import (
	"MatchaServer/common"
)

func (conn ConnFake) SetNewLike(uidSender int, uidReceiver int) error {
	return nil
}

func (conn ConnFake) UnsetLike(uidSender int, uidReceiver int) error {
	return nil
}

func (conn ConnFake) GetUsersThatICanSpeak(myUid int) ([]common.User, error) {
	return nil, nil
}

func (conn ConnFake) IsICanSpeakWithUser(myUid, otherUid int) (bool, error) {
	return false, nil
}