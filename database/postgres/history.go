package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errors"
	"strconv"
)

func (conn ConnDB) SetNewHistoryReference(uid int, targetUid int) error {
	/*
	**	Transaction start
	 */
	tx, err := conn.db.Begin()
	if err != nil {
		return errors.DatabaseTransactionError.AddOriginalError(err)
	}
	defer tx.Rollback()
	/*
	**	Remove all records of users
	 */
	stmt, err := tx.Prepare("DELETE FROM history WHERE uid=$1 AND targetUid=$2")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, targetUid)
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	/*
	**	Create new record
	 */
	stmt, err = tx.Prepare("INSERT INTO history (uid, targetUid, time) VALUES ($1, $2, NOW())")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	result, err := stmt.Exec(uid, targetUid)
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	// handle results
	nbr64, err := result.RowsAffected()
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	if int(nbr64) == 0 {
		return errors.ImpossibleToExecute.WithArguments("Не получилось создать заметку о посещении",
			"Failed to make history reference")
	}
	if int(nbr64) != 1 {
		return errors.NewArg("Создано "+strconv.Itoa(int(nbr64))+" записей",
			strconv.Itoa(int(nbr64))+" times of historycal reference was created")
	}
	/*
	**	Close transaction
	 */
	err = tx.Commit()
	if err != nil {
		return errors.DatabaseTransactionError.AddOriginalError(err)
	}
	return nil
}

func (conn ConnDB) GetHistoryReferencesByUid(uid int) ([]common.HistoryReference, error) {
	var (
		reference  common.HistoryReference
		references []common.HistoryReference
	)
	stmt, err := conn.db.Prepare(`SELECT id, time, users.uid, fname, lname, rating, src FROM
	(SELECT * FROM history WHERE uid=$1 ORDER BY id DESC) AS history_reference
	LEFT JOIN users ON history_reference.targetUid=users.uid
	LEFT JOIN photos ON avaId=pid`)
	if err != nil {
		return nil, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid)
	if err != nil {
		return nil, errors.DatabaseQueryError.AddOriginalError(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&reference.Id, &reference.Time.Time, &reference.User.Uid, &reference.User.Fname,
			&reference.User.Lname, &reference.User.Rating, &reference.User.Avatar)
		if err != nil {
			return nil, errors.DatabaseScanError.AddOriginalError(err)
		}
		references = append(references, reference)
	}
	return references, nil
}

func (conn ConnDB) GetHistoryReferencesByTargetUid(targetUid int) ([]common.HistoryReference, error) {
	var (
		reference  common.HistoryReference
		references []common.HistoryReference
	)
	stmt, err := conn.db.Prepare(`SELECT id, time, users.uid, fname, lname, rating, src FROM
	(SELECT * FROM history WHERE targetUid=$1 ORDER BY id DESC) AS history_reference
	LEFT JOIN users ON history_reference.uid=users.uid
	LEFT JOIN photos ON avaId=pid`)
	if err != nil {
		return nil, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(targetUid)
	if err != nil {
		return nil, errors.DatabaseQueryError.AddOriginalError(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&reference.Id, &reference.Time.Time, &reference.User.Uid, &reference.User.Fname,
			&reference.User.Lname, &reference.User.Rating, &reference.User.Avatar)
		if err != nil {
			return nil, errors.DatabaseScanError.AddOriginalError(err)
		}
		references = append(references, reference)
	}
	return references, nil
}

func (conn ConnDB) DropUserHistory(uid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM history WHERE uid=$1 OR targetUid=$1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid)
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	return nil
}
