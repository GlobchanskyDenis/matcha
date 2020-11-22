package fakeSql

import (
	"MatchaServer/common"
	"MatchaServer/errors"
)

func (conn *ConnFake) SetNewNotif(uidSender int, uidReceiver int, body string) (int, error) {
	var notif common.Notif

	notif.UidSender = uidSender
	notif.UidReceiver = uidReceiver
	notif.Body = body

	for key := 1; ; key++ {
		if _, isExists := conn.notifs[key]; !isExists {
			notif.Nid = key
			break
		}
	}
	conn.notifs[notif.Nid] = notif
	return notif.Nid, nil
}

func (conn *ConnFake) DeleteNotif(nid int) error {
	_, isExists := conn.notifs[nid]
	if !isExists {
		return errors.RecordNotFound
	}
	delete(conn.notifs, nid)
	return nil
}

func (conn *ConnFake) DropUserNotifs(uid int) error {
	for nid, notif := range conn.notifs {
		if notif.UidSender == uid || notif.UidReceiver == uid {
			delete(conn.notifs, nid)
		}
	}
	return nil
}

func (conn *ConnFake) DropReceiverNotifs(uid int) error {
	for nid, notif := range conn.notifs {
		if notif.UidReceiver == uid {
			delete(conn.notifs, nid)
		}
	}
	return nil
}

func (conn ConnFake) GetNotifByNid(nid int) (common.Notif, error) {
	var notif common.Notif

	for _, notif = range conn.notifs {
		if notif.Nid == nid {
			return notif, nil
		}
	}
	return notif, errors.RecordNotFound
}

func (conn ConnFake) GetNotifByUidReceiver(uid int) ([]common.Notif, error) {
	var notifs []common.Notif

	for _, notif := range conn.notifs {
		if notif.UidReceiver == uid {
			notifs = append(notifs, notif)
		}
	}
	return notifs, nil
}
