package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errors"
	"strconv"
	"strings"
	"time"
)

func (conn ConnDB) SetNewLike(uidSender int, uidReceiver int) error {
	/*
	**	Transaction start
	 */
	tx, err := conn.db.Begin()
	if err != nil {
		return errors.DatabaseTransactionError.AddOriginalError(err)
	}
	defer tx.Rollback()
	/*
	**	Set like into likes table
	 */
	stmt, err := tx.Prepare("INSERT INTO likes (uidSender, uidReceiver) VALUES ($1, $2)")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(uidSender, uidReceiver)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "likes_pkey"`) {
			return errors.ImpossibleToExecute //.AddOriginalError(err)
		}
		if strings.Contains(err.Error(), `likeSender_fkey`) || strings.Contains(err.Error(), `likeReceiver_fkey`) {
			return errors.UserNotExist //.AddOriginalError(err)
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
		return errors.NewArg("Добавлено "+strconv.Itoa(int(nbr64))+" лайков",
			strconv.Itoa(int(nbr64))+" likes was added")
	}
	/*
	**	Increment users rating
	 */
	stmt, err = tx.Prepare("UPDATE users SET rating=rating+1 WHERE uid=$1")
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
		return errors.NewArg("Уменьшено "+strconv.Itoa(int(nbr64))+" рейтинга",
			strconv.Itoa(int(nbr64))+" of user rating was decreased")
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

func (conn ConnDB) UnsetLike(uidSender int, uidReceiver int) error {
	/*
	**	Transaction start
	 */
	tx, err := conn.db.Begin()
	if err != nil {
		return errors.DatabaseTransactionError.AddOriginalError(err)
	}
	defer tx.Rollback()

	/*
	**	Unset like into likes table
	 */
	stmt, err := tx.Prepare("DELETE FROM likes WHERE uidSender=$1 AND uidReceiver=$2")
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
		return errors.NewArg("Удалено "+strconv.Itoa(int(nbr64))+" лайков",
			strconv.Itoa(int(nbr64))+" likes was deleted")
	}
	/*
	**	Decrement users rating
	 */
	stmt, err = tx.Prepare("UPDATE users SET rating=rating-1 WHERE uid=$1")
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
		return errors.NewArg("Убрано "+strconv.Itoa(int(nbr64))+" рейтинга",
			strconv.Itoa(int(nbr64))+" of user rating was decreased")
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

func (conn ConnDB) DropUsersLikes(uid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM likes WHERE uidSender=$1 OR uidReceiver=$1")
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

func (conn ConnDB) GetFriendUsers(myUid int) ([]common.FriendUser, error) {
	var (
		user      common.FriendUser
		users     []common.FriendUser
		interests string
		birth     interface{}
		date      time.Time
		ok        bool
	)

	// Стэк запроса: добавление последнего сообщения к пользователю, добавление фотографии к пользователю,
	//  поиск пользователей удовлетворяющих условию - пользователи поставили друг другу лайк.

	query := `SELECT uid, mail, encryptedpass, fname, lname, birth, gender, orientation,
		bio, avaid, latitude, longitude, interests, status, rating, src, uidSender, uidReceiver, body FROM
	(SELECT permitted_users.uid, mail, encryptedpass, fname, lname, birth, gender, orientation,
		bio, avaid, latitude, longitude, interests, status, rating, src FROM
	(SELECT uid, mail, encryptedpass, fname, lname, birth, gender, orientation,
		bio, avaid, latitude, longitude, interests, status, rating FROM
	users INNER JOIN
	(SELECT uidSender FROM
	(SELECT uidSender FROM likes WHERE uidReceiver = $1) AS T1 INNER JOIN
	(SELECT uidReceiver FROM likes WHERE uidSender = $1) AS T2
	 ON T1.uidSender = T2.uidReceiver)
	AS can_talk ON users.uid = can_talk.uidSender)
	AS permitted_users LEFT JOIN photos ON avaId = pid WHERE permitted_users.uid != $1) AS T3 LEFT JOIN
	(SELECT * FROM messages WHERE uidSender = $1 or uidReceiver = $1 ORDER BY mid DESC LIMIT 1) AS
	T4 ON uid = uidSender OR uid = uidReceiver`

	// Неполный запрос - нет присоединенного тела сообщения
	// query := `SELECT permitted_users.uid, mail, encryptedpass, fname, lname, birth, gender, orientation,
	//  							bio, avaid, latitude, longitude, interests, status, rating, src FROM
	// (SELECT uid, mail, encryptedpass, fname, lname, birth, gender, orientation,
	//  						bio, avaid, latitude, longitude, interests, status, rating FROM
	// users INNER JOIN
	// 	(SELECT uidSender FROM
	// 	(SELECT uidSender FROM likes WHERE uidReceiver = $1) AS T1 INNER JOIN
	// 	(SELECT uidReceiver FROM likes WHERE uidSender = $1) AS T2
	// 	ON T1.uidSender = T2.uidReceiver)
	// AS can_talk ON users.uid = can_talk.uidSender)
	// AS permitted_users LEFT JOIN photos ON avaId = pid WHERE permitted_users.uid != $1`

	// Старый интересный запрос. Сохранить на будущее. По идее он менее эффективен чем новый
	// query := "SELECT * FROM users WHERE uid IN (SELECT uidReceiver FROM likes " +
	// 	"WHERE uidSender=$1 AND uidReceiver IN (SELECT uidSender FROM likes WHERE uidReceiver=$1))"

	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return nil, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(myUid)
	if err != nil {
		return nil, errors.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
			&user.Lname, &birth, &user.Gender, &user.Orientation,
			&user.Bio, &user.AvaID, &user.Latitude, &user.Longitude, &interests,
			&user.Status, &user.Rating, &user.Avatar, &user.UidSender,
			&user.UidReceiver, &user.LastMessageBody)
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
		// if stringPtr == nil {
		// 	user.WasTalk = false
		// } else {
		// 	user.WasTalk = true
		// }
		users = append(users, user)
	}
	return users, nil
}

func (conn ConnDB) IsICanSpeakWithUser(myUid, otherUid int) (bool, error) {
	query := "SELECT * FROM likes WHERE uidSender=$1 AND " +
		"uidReceiver IN (SELECT uidSender FROM likes WHERE uidReceiver=$1 AND uidSender=$2)"
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return false, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(myUid, otherUid)
	if err != nil {
		return false, errors.DatabaseExecutingError.AddOriginalError(err)
	}
	// handle results
	nbr64, err := result.RowsAffected()
	if err != nil {
		return false, errors.DatabaseExecutingError.AddOriginalError(err)
	}
	if int(nbr64) == 0 {
		// I cant speak with user
		return false, nil
	}
	if int(nbr64) > 1 {
		return false, errors.NewArg("Неожиданное количество измененных строк - "+strconv.Itoa(int(nbr64)),
			"Unexpectable amount of changed lines - "+strconv.Itoa(int(nbr64)))
	}
	return true, nil
}
