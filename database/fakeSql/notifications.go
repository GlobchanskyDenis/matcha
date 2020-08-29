package fakeSql

import (
	"MatchaServer/common"
)

func (conn ConnFake) SetNewNotif(uidReceiver int, uidSender int, body string) (int, error) {
	var notif common.Notif

	notif.UidSender = uidSender
	notif.UidReceiver = uidReceiver
	notif.Body = body

	for key := 1; ; key++ {
		if _, isExists := conn.notif[key]; !isExists {
			notif.Nid = key
			break
		}
	}

	conn.notif[notif.Nid] = notif
	return notif.Nid, nil
}

func (conn ConnFake) DeleteNotif(nid int) error {
	delete(conn.notif, nid)
	return nil
}

func (conn ConnFake) GetNotifByUidReceiver(uid int) ([]common.Notif, error) {
	var notifs = []common.Notif{}
	var notif common.Notif

	for _, notif = range conn.notif {
		if notif.UidReceiver == uid {
			notifs = append(notifs, notif)
		}
	}
	return notifs, nil
}
