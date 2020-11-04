package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errors"
	"strconv"
	"strings"
	"time"
)

func (conn ConnDB) SetNewClaim(uidSender int, uidReceiver int) error {
	/*
	**	Transaction start
	 */
	tx, err := conn.db.Begin()
	if err != nil {
		return errors.DatabaseTransactionError.AddOriginalError(err)
	}
	defer tx.Rollback()
	/*
	**	Set claim to its table
	 */
	stmt, err := tx.Prepare("INSERT INTO claims (uidSender, uidReceiver) VALUES ($1, $2)")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(uidSender, uidReceiver)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "claims_pkey"`) {
			return errors.ImpossibleToExecute.WithArguments("Вы уже жаловались на этого пользователя",
				"You have already reported this user")
		}
		if strings.Contains(err.Error(), `claims_sender_fkey`) || strings.Contains(err.Error(), `claims_receiver_fkey`) {
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
		return errors.ImpossibleToExecute.WithArguments("Вы уже жаловались на этого пользователя",
			"You have already reported this user")
	}
	if int(nbr64) != 1 {
		return errors.NewArg("Добавлено "+strconv.Itoa(int(nbr64))+" жалоб",
			strconv.Itoa(int(nbr64))+" claims was added")
	}
	/*
	**	Change users search visibility
	 */
	stmt, err = tx.Prepare("UPDATE users SET search_visibility=false WHERE uid=$1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
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
		return errors.ImpossibleToExecute.WithArguments("Не получилось изменить видимость пользователя",
			"Failed to change user visibility")
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

func (conn ConnDB) UnsetClaim(uidSender int, uidReceiver int) error {
	/*
	**	Transaction start
	 */
	tx, err := conn.db.Begin()
	if err != nil {
		return errors.DatabaseTransactionError.AddOriginalError(err)
	}
	defer tx.Rollback()
	/*
	**	Set claim to its table
	 */
	stmt, err := tx.Prepare("DELETE FROM claims WHERE uidSender=$1 AND uidReceiver=$2")
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
		return errors.ImpossibleToExecute.WithArguments("Вы не жаловались на этого пользователя",
			"You have not reported this user")
	}
	if int(nbr64) != 1 {
		return errors.NewArg("Удалено "+strconv.Itoa(int(nbr64))+" жалоб",
			strconv.Itoa(int(nbr64))+" claims was deleted")
	}
	/*
	**	Check is user still in black list from another user
	*/
	stmt, err = tx.Prepare("SELECT uidSender FROM claims WHERE uidReceiver = $1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	rows, err := stmt.Query(uidReceiver)
	if err != nil {
		return errors.DatabaseQueryError.AddOriginalError(err)
	}
	if rows.Next() {
		tx.Commit()
		return nil
	}
	/*
	**	Check is user fills required fields
	*/
	stmt, err = tx.Prepare("SELECT fname, lname, avaID FROM users WHERE uid = $1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	// defer stmt.Close()
	rows, err = stmt.Query(uidReceiver)
	if err != nil {
		return errors.DatabaseQueryError.AddOriginalError(err)
	}
	if !rows.Next() {
		return errors.UserNotExist
	}
	var user common.User
	err = rows.Scan(&user.Fname, &user.Lname, &user.AvaID)
	if err != nil {
		return errors.DatabaseScanError.AddOriginalError(err)
	}
	if user.Fname == "" || user.Lname == "" || user.AvaID == 0 {
		tx.Commit()
		return nil
	}
	/*
	**	Change users search visibility
	 */
	stmt, err = tx.Prepare("UPDATE users SET search_visibility=true WHERE uid=$1")
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
		return errors.ImpossibleToExecute.WithArguments("Не получилось изменить видимость пользователя",
			"Failed to change user visibility")
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

func (conn ConnDB) DropUserClaims(uid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM claims WHERE uidSender=$1 OR uidReceiver=$1")
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

func (conn ConnDB) GetClaimedUsers(uidSender int) ([]common.User, error) {
	var (
		user      common.User
		users     []common.User
		interests string
		birth     interface{}
		date      time.Time
		ok        bool
	)

	query := `SELECT claimed_users.uid, mail, encryptedpass, fname, lname, birth, gender, orientation, bio, avaid,
		latitude, longitude, interests, status, rating, src FROM
	(SELECT * FROM users WHERE uid IN (SELECT uidReceiver FROM claims WHERE uidSender = $1)) AS claimed_users
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
