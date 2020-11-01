package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errors"
	"strconv"
	"strings"
	"time"
)

func (conn ConnDB) SetNewIgnore(uidSender int, uidReceiver int) error {
	/*
	**	Transaction start
	 */
	tx, err := conn.db.Begin()
	if err != nil {
		return errors.DatabaseTransactionError.AddOriginalError(err)
	}
	defer tx.Rollback()
	/*
	**	Set ignore to its table
	 */
	stmt, err := tx.Prepare("INSERT INTO ignore (uidSender, uidReceiver) VALUES ($1, $2)")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(uidSender, uidReceiver)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "ignore_pkey"`) {
			return errors.ImpossibleToExecute
		}
		if strings.Contains(err.Error(), `ignoreSender_fkey`) || strings.Contains(err.Error(), `ignoreReceiver_fkey`) {
			return errors.UserNotExist
		}
		return errors.DatabaseQueryError.AddOriginalError(err)
	}
	// handle results
	nbr64, err := result.RowsAffected()
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	if int(nbr64) == 0 {
		return errors.ImpossibleToExecute
	}
	if int(nbr64) != 1 {
		return errors.NewArg("Добавлено "+strconv.Itoa(int(nbr64))+" игнорирований",
			strconv.Itoa(int(nbr64))+" ignores was added")
	}
	/*
	**	Decrement users rating
	 */
	stmt, err = tx.Prepare("UPDATE users SET rating=rating-3 WHERE uid=$1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	result, err = stmt.Exec(uidReceiver)
	if err != nil {
		return errors.DatabaseQueryError.AddOriginalError(err)
	}
	// handle results
	nbr64, err = result.RowsAffected()
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	if int(nbr64) == 0 {
		return errors.ImpossibleToExecute
	}
	if int(nbr64) != 1 {
		return errors.NewArg("Увеличено "+strconv.Itoa(int(nbr64))+" раз рейтинга",
			strconv.Itoa(int(nbr64))+" times of user rating was decreased")
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

func (conn ConnDB) UnsetIgnore(uidSender int, uidReceiver int) error {
	/*
	**	Transaction start
	 */
	tx, err := conn.db.Begin()
	if err != nil {
		return errors.DatabaseTransactionError.AddOriginalError(err)
	}
	defer tx.Rollback()

	/*
	**	Unset ignore from its table
	 */
	stmt, err := tx.Prepare("DELETE FROM ignore WHERE uidSender=$1 AND uidReceiver=$2")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(uidSender, uidReceiver)
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	// handle results
	nbr64, err := result.RowsAffected()
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	if int(nbr64) == 0 {
		return errors.ImpossibleToExecute
	}
	if int(nbr64) != 1 {
		return errors.NewArg("Удалено "+strconv.Itoa(int(nbr64))+" игнорирований",
			strconv.Itoa(int(nbr64))+" ignores was deleted")
	}
	/*
	**	Increment users rating
	 */
	stmt, err = tx.Prepare("UPDATE users SET rating=rating+3 WHERE uid=$1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	result, err = stmt.Exec(uidReceiver)
	if err != nil {
		return errors.DatabaseQueryError.AddOriginalError(err)
	}
	// handle results
	nbr64, err = result.RowsAffected()
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	if int(nbr64) == 0 {
		return errors.ImpossibleToExecute
	}
	if int(nbr64) != 1 {
		return errors.NewArg("Добавлено "+strconv.Itoa(int(nbr64))+" раз рейтинга",
			strconv.Itoa(int(nbr64))+" times of user rating was increased")
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

func (conn ConnDB) DropUserIgnores(uid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM ignore WHERE uidSender=$1 OR uidReceiver=$1")
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

func (conn ConnDB) GetIgnoredUsers(uidSender int) ([]common.User, error) {
	var (
		user      common.User
		users     []common.User
		interests string
		birth     interface{}
		date      time.Time
		ok        bool
	)

	query := `SELECT ignored_users.uid, mail, encryptedpass, fname, lname, birth, gender, orientation, bio, avaid,
		latitude, longitude, interests, status, rating, src FROM
	(SELECT * FROM users WHERE uid IN (SELECT uidReceiver FROM ignore WHERE uidSender = $1)) AS ignored_users
	LEFT JOIN photos ON avaId=pid`
	rows, err := conn.db.Query(query, uidSender)
	if err != nil {
		return nil, errors.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
			&user.Lname, &birth, &user.Gender, &user.Orientation,
			&user.Bio, &user.AvaID, &user.Latitude, &user.Longitude, &interests,
			&user.Status, &user.Rating, &user.Avatar)
		if err != nil {
			return nil, errors.DatabaseScanError.AddOriginalError(err)
		}
		// handle user Interests
		if len(interests) > 2 {
			strArr := strings.Split(string(interests[1:len(interests)-1]), ",")
			for _, strItem := range strArr {
				user.Interests = append(user.Interests, strItem)
			}
		}
		// handle user birth and age
		if birth != nil {
			date, ok = birth.(time.Time)
			if ok {
				user.Birth.Time = &date
				user.Age = int(time.Since(*user.Birth.Time).Hours() / 24 / 365.27)
			} else {
				return nil, errors.NewArg("не верный тип данных birth", "wrong type of birth")
			}
		} else {
			user.Birth.Time = nil
		}
		users = append(users, user)
	}
	return users, nil
}
