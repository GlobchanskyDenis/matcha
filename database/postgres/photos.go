package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errDef"
)

func (conn ConnDB) SetNewPhoto(uid int, src string) (int, error) {
	var pid int
	stmt, err := conn.db.Prepare("INSERT INTO photos (uid, src) VALUES ($1, $2) RETURNING pid")
	if err != nil {
		return pid, errDef.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	row, err := stmt.Query(uid, src)
	if err != nil {
		return pid, errDef.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
		err = row.Scan(&pid)
		if err != nil {
			return pid, errDef.DatabaseScanError.AddOriginalError(err)
		}
	}
	return pid, nil
}

func (conn ConnDB) DeletePhoto(pid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM photos WHERE pid=$1")
	if err != nil {
		return errDef.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(pid)
	if err != nil {
		return errDef.DatabaseExecutingError.AddOriginalError(err)
	}
	return nil
}

func (conn ConnDB) GetPhotosByUid(uid int) ([]common.Photo, error) {
	var photos = []common.Photo{}
	var photo common.Photo

	stmt, err := conn.db.Prepare("SELECT * FROM photos WHERE uid=$1")
	if err != nil {
		return nil, errDef.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query(uid)
	if err != nil {
		return nil, errDef.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&(photo.Pid), &(photo.Uid), &(photo.Src))
		if err != nil {
			return photos, errDef.DatabaseScanError.AddOriginalError(err)
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

func (conn ConnDB) GetPhotoByPid(pid int) (common.Photo, error) {
	var photo common.Photo

	stmt, err := conn.db.Prepare("SELECT * FROM photos WHERE pid=$1")
	if err != nil {
		return photo, errDef.DatabasePreparingError.AddOriginalError(err)
	}
	row, err := stmt.Query(pid)
	if err != nil {
		return photo, errDef.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
		err = row.Scan(&(photo.Pid), &(photo.Uid), &(photo.Src))
		if err != nil {
			return photo, errDef.DatabaseScanError.AddOriginalError(err)
		}
		return photo, nil
	}
	return photo, errDef.RecordNotFound
}
