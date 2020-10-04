package postgres

import (
	"MatchaServer/common"
	"MatchaServer/errors"
	"database/sql"
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
	_, err = stmt.Exec(uid)
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
	}
	return nil
}

func (conn *ConnDB) UpdateUser(user common.User) error {
	var interests = strings.Join(user.Interests, ",")
	stmt, err := conn.db.Prepare("UPDATE users SET " +
		"mail=$2, encryptedPass=$3, fname=$4, lname=$5, birth=$6, gender=$7, " +
		"orientation=$8, bio=$9, avaID=$10, latitude=$11, longitude=$12, " +
		"interests='{" + interests + "}', status=$13, rating=$14 WHERE uid=$1")
	if err != nil {
		return errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Uid, user.Mail, user.EncryptedPass, user.Fname,
		user.Lname, user.Birth.Time, user.Gender, user.Orientation,
		user.Bio, user.AvaID, user.Latitude, user.Longitude, user.Status, user.Rating)
	if err != nil {
		return errors.DatabaseExecutingError.AddOriginalError(err)
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

	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE uid=$1")
	if err != nil {
		return user, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	row, err = stmt.Query(uid)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
		err = row.Scan(&(user.Uid), &(user.Mail), &(user.EncryptedPass), &(user.Fname),
			&(user.Lname), &birth, &(user.Gender), &(user.Orientation),
			&(user.Bio), &(user.AvaID), &user.Latitude, &user.Longitude, &interests,
			&(user.Status), &(user.Rating))
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
	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE mail=$1")
	if err != nil {
		return user, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	row, err = stmt.Query(mail)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
		err = row.Scan(&(user.Uid), &(user.Mail), &(user.EncryptedPass), &(user.Fname),
			&(user.Lname), &birth, &(user.Gender), &(user.Orientation),
			&(user.Bio), &(user.AvaID), &user.Latitude, &user.Longitude, &interests,
			&(user.Status), &(user.Rating))
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

func (conn *ConnDB) GetUsersByQuery(query string) ([]common.User, error) {
	var (
		user      common.User
		users     []common.User
		interests string
		birth     interface{}
		date      time.Time
		ok        bool
	)
	rows, err := conn.db.Query(query)
	if err != nil {
		return nil, errors.DatabaseQueryError.AddOriginalError(err)
	}
	for rows.Next() {
		err = rows.Scan(&(user.Uid), &(user.Mail), &(user.EncryptedPass), &(user.Fname),
			&(user.Lname), &birth, &(user.Gender), &(user.Orientation),
			&(user.Bio), &(user.AvaID), &user.Latitude, &user.Longitude, &interests,
			&(user.Status), &(user.Rating))
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

	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE mail=$1 AND encryptedPass=$2")
	if err != nil {
		return user, errors.DatabasePreparingError.AddOriginalError(err)
	}
	defer stmt.Close()
	row, err = stmt.Query(mail, encryptedPass)
	if err != nil {
		return user, errors.DatabaseQueryError.AddOriginalError(err)
	}
	if row.Next() {
		err = row.Scan(&(user.Uid), &(user.Mail), &(user.EncryptedPass), &(user.Fname),
			&(user.Lname), &birth, &(user.Gender), &(user.Orientation),
			&(user.Bio), &(user.AvaID), &user.Latitude, &user.Longitude, &interests,
			&(user.Status), &(user.Rating))
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
