package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errors"
	"database/sql"
	"math"
	"strconv"
	"strings"
	"time"
)

func (conn ConnDB) SetNewUser(mail string, encryptedPass string) (common.User, error) {
	var user common.User
	/*
	**	Transaction start
	 */
	tx, err := conn.db.Begin()
	if err != nil {
		return user, errors.DatabaseTransactionError.AddOriginalError(err)
	}
	defer tx.Rollback()
	/*
	**	Set new user
	 */
	stmt, err := tx.Prepare(`INSERT INTO users (mail, encryptedPass, status) VALUES ($1, $2, 'not confirmed')
		RETURNING uid, mail, status`)
	if err != nil {
		return user, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(mail, encryptedPass).Scan(&user.Uid, &user.Mail, &user.Status)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	/*
	**	Set user should ignore himself in search queries
	 */
	stmt, err = tx.Prepare("INSERT INTO ignores (uidSender, uidReceiver) VALUES ($1, $1)")
	if err != nil {
		return user, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.Uid)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "ignores_pkey"`) {
			return user, errors.ImpossibleToExecute.WithArguments("Вы уже игнорируете этого пользователя",
				"You are already ignoring this user")
		}
		if strings.Contains(err.Error(), `ignores_sender_fkey`) || strings.Contains(err.Error(), `ignores_receiver_fkey`) {
			return user, errors.UserNotExist
		}
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	// handle results
	nbr64, err := result.RowsAffected()
	if err != nil {
		return user, errors.DatabaseExecutingError.AddOriginalError(err)
	}
	if int(nbr64) == 0 {
		return user, errors.ImpossibleToExecute.WithArguments("Вы уже игнорируете этого пользователя",
			"You are already ignoring this user")
	}
	if int(nbr64) != 1 {
		return user, errors.NewArg("Добавлено "+strconv.Itoa(int(nbr64))+" игнорирований",
			strconv.Itoa(int(nbr64))+" ignores was added")
	}
	/*
	**	Close transaction
	 */
	err = tx.Commit()
	if err != nil {
		return user, errors.DatabaseTransactionError.AddOriginalError(err)
	}
	return user, nil
}

func (conn *ConnDB) DeleteUser(uid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM users WHERE uid=$1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(uid)
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

func (conn *ConnDB) UpdateUser(user common.User) error {
	var interests = strings.Join(user.Interests, ",")
	var searchVisibility bool
	if user.Fname != "" && user.Lname != "" && user.AvaID != nil {
		stmt, err := conn.db.Prepare("SELECT uidSender FROM claims WHERE uidReceiver = $1")
		if err != nil {
			stmt.Close()
			return errors.DatabasePreparingError.AddOriginalError(err)
		}
		rows, err := stmt.Query(user.Uid)
		if err != nil {
			stmt.Close()
			return errors.DatabaseQueryError.AddOriginalError(err)
		}
		defer rows.Close()
		if !rows.Next() {
			searchVisibility = true
		}
	}
	stmt, err := conn.db.Prepare("UPDATE users SET " +
		"mail=$2, encryptedPass=$3, fname=$4, lname=$5, birth=$6, gender=$7, " +
		"orientation=$8, bio=$9, avaID=$10, latitude=$11, longitude=$12, " +
		"interests='{" + interests + "}', status=$13, rating=$14, search_visibility=$15 WHERE uid=$1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(user.Uid, user.Mail, user.EncryptedPass, user.Fname,
		user.Lname, user.Birth.Time, user.Gender, user.Orientation,
		user.Bio, user.AvaID, user.Latitude, user.Longitude, user.Status, user.Rating, searchVisibility)
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

func (conn *ConnDB) GetUserByUid(uid int) (common.User, error) {
	var (
		user      common.User
		err       error
		rows      *sql.Rows
		birth     interface{}
		date      time.Time
		ok        bool
		interests string
	)
	query := `SELECT users.uid, mail, encryptedpass, fname, lname, birth, gender, orientation, bio, avaid,
		latitude, longitude, interests, status, rating, src FROM
		users LEFT JOIN photos ON avaId = pid WHERE users.uid=$1`
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return user, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	rows, err = stmt.Query(uid)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
			&user.Lname, &birth, &user.Gender, &user.Orientation,
			&user.Bio, &user.AvaID, &user.Latitude, &user.Longitude, &interests,
			&user.Status, &user.Rating, &user.Avatar)
		if err != nil {
			return user, errors.DatabaseScanError.AddOriginalError(err)
		}
	} else {
		return user, errors.RecordNotFound
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
			return user, errors.NewArg("не верный тип данных birth", "wrong type of birth")
		}
	} else {
		user.Birth.Time = nil
	}
	return user, nil
}

func (conn *ConnDB) GetTargetUserByUid(myUid int, targetUid int) (common.TargetUser, error) {
	var (
		user             common.TargetUser
		err              error
		rows             *sql.Rows
		birth            interface{}
		date             time.Time
		ok               bool
		interests        string
		ptr1, ptr2, ptr3 *int
	)
	query := `SELECT users.uid, mail, encryptedpass, fname, lname, birth, gender, orientation, bio, avaid, latitude,
	longitude, interests, status, rating, src, is_ignored.uidReceiver, is_claimed.uidReceiver, is_liked.uidReceiver FROM
	users LEFT JOIN photos ON avaId = pid
	LEFT JOIN (SELECT uidReceiver FROM ignores WHERE uidSender=$2)
		AS is_ignored ON users.uid=is_ignored.uidReceiver
	LEFT JOIN (SELECT uidReceiver FROM claims WHERE uidSender=$2)
		AS is_claimed ON users.uid=is_claimed.uidReceiver
	LEFT JOIN (SELECT uidReceiver FROM likes WHERE uidSender=$2)
		AS is_liked ON users.uid=is_liked.uidReceiver
	WHERE users.uid=$1`
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return user, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	rows, err = stmt.Query(targetUid, myUid)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
			&user.Lname, &birth, &user.Gender, &user.Orientation,
			&user.Bio, &user.AvaID, &user.Latitude, &user.Longitude, &interests,
			&user.Status, &user.Rating, &user.Avatar, &ptr1, &ptr2, &ptr3)
		if err != nil {
			return user, errors.DatabaseScanError.AddOriginalError(err)
		}
	} else {
		return user, errors.RecordNotFound
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
			return user, errors.NewArg("не верный тип данных birth", "wrong type of birth")
		}
	} else {
		user.Birth.Time = nil
	}
	// handle isIgnored param
	if ptr1 != nil {
		user.IsIgnored = true
	} else {
		user.IsIgnored = false
	}
	// handle isClaimed param
	if ptr2 != nil {
		user.IsClaimed = true
	} else {
		user.IsClaimed = false
	}
	// handle isLiked param
	if ptr3 != nil {
		user.IsLiked = true
	} else {
		user.IsLiked = false
	}
	return user, nil
}

func (conn *ConnDB) GetUserWithLikeInfo(targetUid int, myUid int) (common.SearchUser, error) {
	var (
		user      common.SearchUser
		err       error
		rows      *sql.Rows
		birth     interface{}
		date      time.Time
		ok        bool
		interests string
		intPtr    *int
	)
	query := `SELECT users.uid, mail, encryptedpass, fname, lname, birth, gender, orientation, bio, avaid,
		latitude, longitude, interests, status, rating, src, uidReceiver FROM
		users LEFT JOIN photos ON avaId = pid
		LEFT JOIN (SELECT uidReceiver FROM likes WHERE uidSender=$2) AS T1 ON users.uid=uidReceiver WHERE users.uid=$1`
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return user, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	rows, err = stmt.Query(targetUid, myUid)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
			&user.Lname, &birth, &user.Gender, &user.Orientation,
			&user.Bio, &user.AvaID, &user.Latitude, &user.Longitude, &interests,
			&user.Status, &user.Rating, &user.Avatar, &intPtr)
		if err != nil {
			return user, errors.DatabaseScanError.AddOriginalError(err)
		}
	} else {
		return user, errors.RecordNotFound
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
			return user, errors.NewArg("не верный тип данных birth", "wrong type of birth")
		}
	} else {
		user.Birth.Time = nil
	}
	if intPtr != nil {
		user.IsLiked = true
	} else {
		user.IsLiked = false
	}
	return user, nil
}

func (conn *ConnDB) GetUserByMail(mail string) (common.User, error) {
	var (
		user      common.User
		err       error
		rows      *sql.Rows
		birth     interface{}
		date      time.Time
		ok        bool
		interests string
	)
	query := `SELECT users.uid, mail, encryptedpass, fname, lname, birth, gender, orientation, bio, avaid,
		latitude, longitude, interests, status, rating, src FROM
		users LEFT JOIN photos ON avaId = pid WHERE users.mail=$1`
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return user, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	rows, err = stmt.Query(mail)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
			&user.Lname, &birth, &user.Gender, &user.Orientation,
			&user.Bio, &user.AvaID, &user.Latitude, &user.Longitude, &interests,
			&user.Status, &user.Rating, &user.Avatar)
		if err != nil {
			return user, errors.DatabaseScanError.AddOriginalError(err)
		}
	} else {
		return user, errors.RecordNotFound
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
			return user, errors.NewArg("не верный тип данных birth", "wrong type of birth")
		}
	} else {
		user.Birth.Time = nil
	}
	return user, nil
}

func (conn *ConnDB) GetUsersByQuery(query string, sourceUser common.User) ([]common.SearchUser, error) {
	var (
		user             common.SearchUser
		users            []common.SearchUser
		interests        string
		birth            interface{}
		intPtr1, intPtr2 *int
		date             time.Time
		ok               bool
	)
	rows, err := conn.db.Query(query, sourceUser.Uid)
	if err != nil {
		return nil, errors.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		user.Interests = nil
		user.Birth.Time = nil

		err = rows.Scan(&user.Uid, &user.Fname, &user.Lname, &birth, &user.Gender,
			&user.Orientation, &user.AvaID, &user.Latitude, &user.Longitude,
			&interests, &user.Rating, &user.Avatar, &intPtr1, &intPtr2)
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
		// handle flag that user was liked by user who searches
		if intPtr1 == nil {
			user.IsMatch = false
		} else {
			user.IsMatch = true
		}
		if intPtr2 == nil {
			user.IsLiked = false
		} else {
			user.IsLiked = true
		}
		if user.Latitude != nil && user.Longitude != nil &&
			sourceUser.Latitude != nil && sourceUser.Longitude != nil {
			deltaLat := *user.Latitude - *sourceUser.Latitude
			deltaLong := *user.Longitude - *sourceUser.Longitude
			Range := math.Sqrt(deltaLat*deltaLat+deltaLong*deltaLong) * 111
			user.Range = &Range
		}
		users = append(users, user)
	}
	return users, nil
}

func (conn *ConnDB) GetUserForAuth(mail string, encryptedPass string) (common.User, error) {
	var (
		user      common.User
		err       error
		rows      *sql.Rows
		birth     interface{}
		date      time.Time
		interests string
		ok        bool
	)

	query := `SELECT users.uid, mail, encryptedpass, fname, lname, birth, gender, orientation, bio, avaid,
		latitude, longitude, interests, status, rating, src FROM
		users LEFT JOIN photos ON avaId = pid WHERE mail=$1 AND encryptedPass=$2`
	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return user, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	rows, err = stmt.Query(mail, encryptedPass)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	defer rows.Close()
	if rows.Next() {
		err = rows.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
			&user.Lname, &birth, &user.Gender, &user.Orientation,
			&user.Bio, &user.AvaID, &user.Latitude, &user.Longitude, &interests,
			&user.Status, &user.Rating, &user.Avatar)
		if err != nil {
			return user, errors.DatabaseScanError.AddOriginalError(err)
		}
	} else {
		return user, errors.RecordNotFound
	}
	// handle user Interests
	// parse string of interests into []string
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
			return user, errors.NewArg("не верный тип данных birth", "wrong type of birth")
		}
	} else {
		user.Birth.Time = nil
	}
	return user, nil
}

func (conn ConnDB) IsUserExistsByMail(mail string) (bool, error) {
	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE mail=$1")
	if err != nil {
		return false, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(mail)
	if err != nil {
		return false, errors.DatabaseQueryError.AddOriginalError(err)
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	}
	return false, nil
}

func (conn ConnDB) IsUserExistsByUid(uid int) (bool, error) {
	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE uid=$1")
	if err != nil {
		return false, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(uid)
	if err != nil {
		return false, errors.DatabaseQueryError.AddOriginalError(err)
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	}
	return false, nil
}
