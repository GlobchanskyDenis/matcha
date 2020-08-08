package postgres

import (
	"MatchaServer/config"
	"MatchaServer/errDef"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"
)

func (conn ConnDB) SetNewUser(mail string, encryptedPass string) (config.User, error) {
	var user config.User
	stmt, err := conn.db.Prepare("INSERT INTO users (mail, encryptedPass) VALUES ($1, $2) RETURNING uid, mail")
	if err != nil {
		return user, errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	err = stmt.QueryRow(mail, encryptedPass).Scan(&user.Uid, &user.Mail)
	if err != nil {
		return user, errors.New(err.Error() + " in executing")
	}
	return user, nil
}

func (conn *ConnDB) DeleteUser(uid int) error {
	stmt, err := conn.db.Prepare("DELETE FROM users WHERE uid=$1")
	if err != nil {
		return errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid)
	if err != nil {
		return errors.New(err.Error() + " in executing")
	}
	return nil
}

func (conn *ConnDB) UpdateUser(user config.User) error {
	var interests string
	for _, item := range user.Interests {
		interests += item + ", "
	}
	if len(interests) > 2 {
		interests = string(interests[:len(interests) - 2])
	}
	stmt, err := conn.db.Prepare("UPDATE users SET " +
		"mail=$2, encryptedPass=$3, fname=$4, lname=$5, birth=$6, gender=$7, " +
		"orientation=$8, bio=$9, avaID=$10, latitude=$11, longitude=$12, " +
		"interests='{" + interests + "}', status=$13, rating=$14 WHERE uid=$1")
	if err != nil {
		return errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Uid, user.Mail, user.EncryptedPass, user.Fname,
		user.Lname, user.Birth.Format("2006-01-02"), user.Gender, user.Orientation,
		user.Bio, user.AvaID, user.Latitude, user.Longitude, user.Status, user.Rating)
	if err != nil {
		return errors.New(err.Error() + " in executing")
	}
	return nil
}

// Тут никак не реализован фильтр !!!!!!!!!!!!!!!
func (conn ConnDB) SearchUsersByOneFilter(filter string) ([]config.User, error) {
	var (
		users []config.User
		user  config.User
		err   error
		rows  *sql.Rows
		birth string
		interests string
	)

	rows, err = conn.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&user.Uid, &user.Mail, &user.EncryptedPass, &user.Fname,
			&user.Lname, &birth, &user.Gender, &user.Orientation,
			&user.Bio, &user.AvaID, &user.Latitude, &user.Longitude, &interests,
			&user.Status, &user.Rating)
		if err != nil {
			return nil, err
		}
		// handle user Interests
		if len(interests) > 2 {
			strArr := strings.Split(string(interests[1:len(interests)-1]), ",")
			for _, strItem := range strArr {
				user.Interests = append(user.Interests, strItem)
			}
		}
		// handle user birth and age
		if len(birth) > 10 {
			birth = string(birth[:10])
			user.Birth, err = time.Parse("2006-01-02", birth)
			if err != nil {
				return nil, err
			}
			user.Age = int(time.Since(user.Birth).Hours() / 24 / 365.27)
		}
		users = append(users, user)
	}
	return users, err
}

func (conn *ConnDB) GetUserByUid(uid int) (config.User, error) {
	var (
		user config.User
		err  error
		row  *sql.Rows
		birth string
		interests string
	)

	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE uid=$1")
	if err != nil {
		return user, errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	row, err = stmt.Query(uid)
	if err != nil {
		return user, errors.New(err.Error() + " in query")
	}
	if row.Next() {
		err = row.Scan(&(user.Uid), &(user.Mail), &(user.EncryptedPass), &(user.Fname),
			&(user.Lname), &birth, &(user.Gender), &(user.Orientation),
			&(user.Bio), &(user.AvaID), &user.Latitude, &user.Longitude, &interests,
			&(user.Status), &(user.Rating))
		if err != nil {
			return user, err
		}
	} else {
		return user, errDef.RecordNotFound
	}
	// handle user Interests
	if len(interests) > 2 {
		strArr := strings.Split(string(interests[1:len(interests)-1]), ",")
		for _, strItem := range strArr {
			user.Interests = append(user.Interests, strItem)
		}
	}
	// handle user birth and age
	if len(birth) > 10 {
		birth = string(birth[:10])
		user.Birth, err = time.Parse("2006-01-02", birth)
		if err != nil {
			return user, err
		}
		user.Age = int(time.Since(user.Birth).Hours() / 24 / 365.27)
	}
	return user, nil
}

func (conn *ConnDB) GetUserByMail(mail string) (config.User, error) {
	var (
		user config.User
		err  error
		row  *sql.Rows
		birth string
		interests string
	)
	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE mail=$1")
	if err != nil {
		return user, errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	row, err = stmt.Query(mail)
	if err != nil {
		return user, errors.New(err.Error() + " in query")
	}
	if row.Next() {
		err = row.Scan(&(user.Uid), &(user.Mail), &(user.EncryptedPass), &(user.Fname),
			&(user.Lname), &birth, &(user.Gender), &(user.Orientation),
			&(user.Bio), &(user.AvaID), &user.Latitude, &user.Longitude, &interests,
			&(user.Status), &(user.Rating))
		if err != nil {
			return user, err
		}
	} else {
		return user, errDef.RecordNotFound
	}
	// handle user Interests
	if len(interests) > 2 {
		strArr := strings.Split(string(interests[1:len(interests)-1]), ",")
		for _, strItem := range strArr {
			user.Interests = append(user.Interests, strItem)
		}
	}
	// handle user birth and age
	if len(birth) > 10 {
		birth = string(birth[:10])
		user.Birth, err = time.Parse("2006-01-02", birth)
		if err != nil {
			return user, err
		}
		user.Age = int(time.Since(user.Birth).Hours() / 24 / 365.27)
	}
	return user, nil
}

func (conn *ConnDB) GetUserForAuth(mail string, encryptedPass string) (config.User, error) {
	var (
		user config.User
		err  error
		row  *sql.Rows
		birth string
		interests string
	)

	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE mail=$1 AND encryptedPass=$2")
	if err != nil {
		return user, errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	row, err = stmt.Query(mail, encryptedPass)
	if err != nil {
		return user, errors.New(err.Error() + " in query")
	}
	if row.Next() {
		err = row.Scan(&(user.Uid), &(user.Mail), &(user.EncryptedPass), &(user.Fname),
			&(user.Lname), &birth, &(user.Gender), &(user.Orientation),
			&(user.Bio), &(user.AvaID), &user.Latitude, &user.Longitude, &interests,
			&(user.Status), &(user.Rating))
		if err != nil {
			return user, err
		}
	} else {
		return user, errDef.RecordNotFound
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
	if len(birth) > 10 {
		birth = string(birth[:10])
		user.Birth, err = time.Parse("2006-01-02", birth)
		if err != nil {
			return user, err
		}
		user.Age = int(time.Since(user.Birth).Hours() / 24 / 365.27)
	}
	return user, nil
}

func (conn *ConnDB) GetLoggedUsers(uid []int) ([]config.User, error) {
	var users = []config.User{}
	var user config.User
	var birth string
	var interests string

	if len(uid) == 0 {
		return users, nil
	}

	query := "SELECT * FROM users WHERE uid IN ("
	length := len(uid)
	for i := 1; i <= length; i++ {
		query += "$" + strconv.Itoa(i) + ", "
	}
	tmp := []byte(query)
	tmp = tmp[:len(tmp) - 2]
	query = string(tmp) + ")"

	stmt, err := conn.db.Prepare(query)
	if err != nil {
		return users, errors.New(err.Error() + " in preparing")
	}

	interfaceSlice := make([]interface{}, len(uid))
	for i, val := range uid {
		interfaceSlice[i] = val
	}

	rows, err := stmt.Query(interfaceSlice...)
	for rows.Next() {
		err = rows.Scan(&(user.Uid), &(user.Mail), &(user.EncryptedPass), &(user.Fname),
			&(user.Lname), &birth, &(user.Gender), &(user.Orientation),
			&(user.Bio), &(user.AvaID), &user.Latitude, &user.Longitude, &interests,
			&(user.Status), &(user.Rating))
		if err != nil {
			return nil, err
		}
		// handle user Interests
		if len(interests) > 2 {
			strArr := strings.Split(string(interests[1:len(interests)-1]), ",")
			for _, strItem := range strArr {
				user.Interests = append(user.Interests, strItem)
			}
		}
		// handle user birth and age
		if len(birth) > 10 {
			birth = string(birth[:10])
			user.Birth, err = time.Parse("2006-01-02", birth)
			if err != nil {
				return nil, err
			}
			user.Age = int(time.Since(user.Birth).Hours() / 24 / 365.27)
		}
		users = append(users, user)
	}
	return users, nil
}

func (conn ConnDB) IsUserExistsByMail(mail string) (bool, error) {
	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE mail=$1")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	row, err := stmt.Query(mail)
	if err != nil {
		return false, err
	}
	if row.Next() {
		return true, nil
	}
	return false, nil
}

func (conn ConnDB) IsUserExistsByUid(uid int) (bool, error) {
	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE uid=$1")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	row, err := stmt.Query(uid)
	if err != nil {
		return false, err
	}
	if row.Next() {
		return true, nil
	}
	return false, nil
}
