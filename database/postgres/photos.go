package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errors"
	"strconv"
)

func (conn ConnDB) SetNewPhoto(uid int, src string) (int, error) {
	var pid int
	stmt, err := conn.db.Prepare("INSERT INTO photos (uid, src) VALUES ($1, $2) RETURNING pid")
	if err != nil {
		return pid, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	row, err := stmt.Query(uid, src)
	if err != nil {
		return pid, errors.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
		err = row.Scan(&pid)
		if err != nil {
			return pid, errors.DatabaseScanError.AddOriginalError(err)
		}
	}
	return pid, nil
}

func (conn ConnDB) DeletePhoto(pid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM photos WHERE pid=$1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(pid)
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	// handle results
	nbr64, err := result.RowsAffected()
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	if int(nbr64) != 1 {
		return errors.NewArg("Неожиданное количество измененных строк - "+strconv.Itoa(int(nbr64)),
			"Unexpectable amount of changed lines - "+strconv.Itoa(int(nbr64)))
	}
	return nil
}

func (conn ConnDB) GetPhotosByUid(uid int) ([]common.Photo, error) {
	var photos = []common.Photo{}
	var photo common.Photo

	stmt, err := conn.db.Prepare("SELECT * FROM photos WHERE uid=$1")
	if err != nil {
		return nil, errors.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query(uid)
	if err != nil {
		return nil, errors.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&(photo.Pid), &(photo.Uid), &(photo.Src))
		if err != nil {
			return photos, errors.DatabaseScanError.AddOriginalError(err)
		}
		photos = append(photos, photo)
	}
	return photos, nil
}

func (conn ConnDB) GetPhotoByPid(pid int) (common.Photo, error) {
	var photo common.Photo

	stmt, err := conn.db.Prepare("SELECT * FROM photos WHERE pid=$1")
	if err != nil {
		return photo, errors.DatabasePreparingError.AddOriginalError(err)
	}
	row, err := stmt.Query(pid)
	if err != nil {
		return photo, errors.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
		err = row.Scan(&(photo.Pid), &(photo.Uid), &(photo.Src))
		if err != nil {
			return photo, errors.DatabaseScanError.AddOriginalError(err)
		}
		return photo, nil
	}
	return photo, errors.RecordNotFound
}
