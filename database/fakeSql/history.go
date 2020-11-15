package fakeSql

import (
	"MatchaServer/common"
)

func (conn ConnFake) SetNewHistoryReference(uid int, targetUid int) error {
	return nil
}

func (conn ConnFake) GetHistoryReferencesByUid(uid int) ([]common.HistoryReference, error) {
	return nil, nil
}

func (conn ConnFake) GetHistoryReferencesByTargetUid(targetUid int) ([]common.HistoryReference, error) {
	return nil, nil
}

func (conn ConnFake) DropUserHistory(uid int) error {
	return nil
}
