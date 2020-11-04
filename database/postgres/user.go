package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errors"
	"database/sql"
	"strconv"
	"strings"
	"time"
)

func (conn ConnDB) SetNewUser(mail string, encryptedPass string) (common.User, error) {
	var user common.User
	stmt, err := conn.db.Prepare("INSERT INTO users (mail, encryptedPass) VALUES ($1, $2) RETURNING uid, mail")
	if err != nil {
		return user, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(mail, encryptedPass).Scan(&user.Uid, &user.Mail)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
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
	if user.Fname != "" && user.Lname != "" && user.AvaID != 0 {
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
		row       *sql.Rows
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
	row, err = stmt.Query(uid)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
		err = row.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
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

func (conn *ConnDB) GetUserByMail(mail string) (common.User, error) {
	var (
		user      common.User
		err       error
		row       *sql.Rows
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
	row, err = stmt.Query(mail)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
		err = row.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
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

func (conn *ConnDB) GetUsersByQuery(query string) ([]common.SearchUser, error) {
	var (
		user      common.SearchUser
		users     []common.SearchUser
		interests string
		birth     interface{}
		intPtr    *int
		date      time.Time
		ok        bool
	)
	rows, err := conn.db.Query(query)
	if err != nil {
		return nil, errors.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
			&user.Lname, &birth, &user.Gender, &user.Orientation,
			&user.Bio, &user.AvaID, &user.Latitude, &user.Longitude, &interests,
			&user.Status, &user.Rating, &user.Avatar, &intPtr)
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
		if intPtr == nil {
			user.IsLiked = false
		} else {
			user.IsLiked = true
		}
		users = append(users, user)
	}
	return users, nil
}

func (conn *ConnDB) GetUserForAuth(mail string, encryptedPass string) (common.User, error) {
	var (
		user      common.User
		err       error
		row       *sql.Rows
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
	row, err = stmt.Query(mail, encryptedPass)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
		err = row.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
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
	row, err := stmt.Query(mail)
	if err != nil {
		return false, errors.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
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
	row, err := stmt.Query(uid)
	if err != nil {
		return false, errors.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
		return true, nil
	}
	return false, nil
}
