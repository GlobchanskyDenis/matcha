package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errDef"
	"errors"
)

func (conn ConnDB) SetNewPhoto(uid int, src string) (int, error) {
	var pid int
	stmt, err := conn.db.Prepare("INSERT INTO photos (uid, src) VALUES ($1, $2) RETURNING pid")
	if err != nil {
		return pid, errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	row, err := stmt.Query(uid, src)
	if err != nil {
		return pid, errors.New(err.Error() + " in executing")
	}
	if row.Next() {
		err = row.Scan(&pid)
		if err != nil {
			return pid, errors.New(err.Error() + " in rows")
		}
	}
	return pid, nil
}

func (conn ConnDB) DeletePhoto(pid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM photos WHERE pid=$1")
	if err != nil {
		return errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	_, err = stmt.Exec(pid)
	if err != nil {
		return errors.New(err.Error() + " in executing")
	}
	return nil
}

func (conn ConnDB) GetPhotosByUid(uid int) ([]common.Photo, error) {
	var photos = []common.Photo{}
	var photo common.Photo

	stmt, err := conn.db.Prepare("SELECT * FROM photos WHERE uid=$1")
	if err != nil {
		return photos, errors.New(err.Error() + " in preparing")
	}
	rows, err := stmt.Query(uid)
	if err != nil {
		return photos, errors.New(err.Error() + " in executing")
	}
	for rows.Next() {
		err = rows.Scan(&(photo.Pid), &(photo.Uid), &(photo.Src))
		if err != nil {
			return photos, errors.New(err.Error() + " in rows")
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

func (conn ConnDB) GetPhotoByPid(pid int) (common.Photo, error) {
	var photo common.Photo

	stmt, err := conn.db.Prepare("SELECT * FROM photos WHERE pid=$1")
	if err != nil {
		return photo, errors.New(err.Error() + " in preparing")
	}
	row, err := stmt.Query(pid)
	if err != nil {
		return photo, errors.New(err.Error() + " in executing")
	}
	if row.Next() {
		err = row.Scan(&(photo.Pid), &(photo.Uid), &(photo.Src))
		if err != nil {
			return photo, errors.New(err.Error() + " in rows")
		}
		return photo, nil
	}
	return photo, errDef.RecordNotFound
}
