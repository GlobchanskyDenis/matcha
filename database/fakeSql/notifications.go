package fakeSql

import (
	"MatchaServer/config"
)

func (conn ConnFake) SetNewNotif(uidReceiver int, uidSender int, body string) (int, error) {
	var notif config.Notif

	notif.UidSender = uidSender
	notif.UidReceiver = uidReceiver
	notif.Body = body

	for key := 1; ; key++ {
		if _, isExists := conn.notif[key]; !isExists {
			notif.Nid = key
			break;
		}
	}

	conn.notif[notif.Nid] = notif
	return notif.Nid, nil
}

func (conn ConnFake) DeleteNotif(nid int) error {
	delete(conn.notif, nid)
	return nil
}

func (conn ConnFake) GetNotifByUidReceiver(uid int) ([]config.Notif, error) {
	var notifs = []config.Notif{}
	var notif config.Notif

	for _, notif = range conn.notif {
		if notif.UidReceiver == uid {
			notifs = append(notifs, notif)
		}
	}
	return notifs, nil
}