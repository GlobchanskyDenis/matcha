package fakeSql

import (
	"MatchaServer/common"
)

func (conn ConnFake) SetNewClaim(uidSender int, uidReceiver int) error {
	return nil
}

func (conn ConnFake) UnsetClaim(uidSender int, uidReceiver int) error {
	return nil
}

func (conn ConnFake) DropUserClaims(uid int) error {
	return nil
}

func (conn ConnFake) GetClaimedUsers(uidSender int) ([]common.User, error) {
	return nil, nil
}
