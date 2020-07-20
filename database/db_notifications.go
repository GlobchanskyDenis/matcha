package database

import (
	"MatchaServer/config"
	"errors"
)

func (conn ConnDB) SetNewNotif(uid int, body string) error {
	stmt, err := conn.db.Prepare("INSERT INTO notif (uid, body) VALUES ($1, $2)")
	if err != nil {
		return errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, body)
	if err != nil {
		return errors.New(err.Error() + " in executing")
	}
	return nil
}

func (conn ConnDB) DeleteNotif(nid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM notif WHERE nid=$1")
	if err != nil {
		return errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	_, err = stmt.Exec(nid)
	if err != nil {
		return errors.New(err.Error() + " in executing")
	}
	return nil
}

func (conn ConnDB) GetNotifByUid(uid int) ([]config.Notif, error) {
	var notifs = []config.Notif{}
	var notif config.Notif

	stmt, err := conn.db.Prepare("SELECT * FROM notif WHERE uid=$1")
	if err != nil {
		return notifs, errors.New(err.Error() + " in preparing")
	}
	rows, err := stmt.Query(uid)
	if err != nil {
		return notifs, errors.New(err.Error() + " in executing")
	}
	for rows.Next() {
		err = rows.Scan(&(notif.Nid), &(notif.Uid), &(notif.Body))
		if err != nil {
			return notifs, errors.New(err.Error() + " in rows")
		}
		notifs = append(notifs, notif)
	}
	return notifs, nil
}