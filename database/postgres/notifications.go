package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errDef"
)

func (conn ConnDB) SetNewNotif(uidReceiver int, uidSender int, body string) (int, error) {
	var nid int
	stmt, err := conn.db.Prepare("INSERT INTO notifs (uidSender, uidReceiver, body) VALUES ($1, $2, $3) RETURNING nid")
	if err != nil {
		return nid, errDef.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(uidSender, uidReceiver, body).Scan(&nid)
	if err != nil {
		return nid, errDef.DatabaseQueryError.AddOriginalError(err)
	}
	return nid, nil
}

func (conn ConnDB) DeleteNotif(nid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM notifs WHERE nid=$1")
	if err != nil {
		return errDef.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(nid)
	if err != nil {
		return errDef.DatabaseExecutingError.AddOriginalError(err)
	}
	return nil
}

func (conn ConnDB) GetNotifByUidReceiver(uid int) ([]common.Notif, error) {
	var notifs = []common.Notif{}
	var notif common.Notif

	stmt, err := conn.db.Prepare("SELECT * FROM notifs WHERE uidReceiver=$1")
	if err != nil {
		return nil, errDef.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query(uid)
	if err != nil {
		return nil, errDef.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&(notif.Nid), &(notif.UidSender), &(notif.UidReceiver), &(notif.Body))
		if err != nil {
			return notifs, errDef.DatabaseScanError.AddOriginalError(err)
		}
		notifs = append(notifs, notif)
	}
	return notifs, nil
}
