package database

import (
	"MatchaServer/config"
	"database/sql"
	"errors"
	"strconv"
)

func (conn ConnDB) SetNewUser(mail string, passwd string) error {
	stmt, err := conn.db.Prepare("INSERT INTO users (mail, passwd) VALUES ($1, $2)")
	if err != nil {
		return errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	_, err = stmt.Exec(mail, passwd)
	if err != nil {
		return errors.New(err.Error() + " in executing")
	}
	return nil
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
	stmt, err := conn.db.Prepare("UPDATE users SET " +
		"mail=$2, passwd=$3, fname=$4, lname=$5, age=$6, gender=$7, " +
		"orientation=$8, biography=$9, avaPhotoID=$10, accType=$11, rating=$12  " +
		"WHERE uid=$1")
	if err != nil {
		return errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Uid, user.Mail, user.Passwd, user.Fname,
		user.Lname, user.Age, user.Gender, user.Orientation,
		user.Biography, user.AvaPhotoID, user.AccType, user.Rating)
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
	)

	rows, err = conn.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&(user.Uid), &(user.Mail), &(user.Passwd), &(user.Fname),
			&(user.Lname), &(user.Age), &(user.Gender), &(user.Orientation),
			&(user.Biography), &(user.AvaPhotoID), &(user.AccType), &(user.Rating))
		if err != nil {
			return nil, err
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
		err = row.Scan(&(user.Uid), &(user.Mail), &(user.Passwd), &(user.Fname),
			&(user.Lname), &(user.Age), &(user.Gender), &(user.Orientation),
			&(user.Biography), &(user.AvaPhotoID), &(user.AccType), &(user.Rating))
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (conn *ConnDB) GetUserByMail(mail string) (config.User, error) {
	var (
		user config.User
		err  error
		row  *sql.Rows
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
		err = row.Scan(&(user.Uid), &(user.Mail), &(user.Passwd), &(user.Fname),
			&(user.Lname), &(user.Age), &(user.Gender), &(user.Orientation),
			&(user.Biography), &(user.AvaPhotoID), &(user.AccType), &(user.Rating))
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (db *ConnDB) GetUserDataForAuth(mail string, passwd string) (config.User, error) {
	var (
		user config.User
		err  error
		row  *sql.Rows
	)

	stmt, err := db.db.Prepare("SELECT * FROM users WHERE mail=$1 AND passwd=$2")
	if err != nil {
		return user, errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	row, err = stmt.Query(mail, passwd)
	if err != nil {
		return user, errors.New(err.Error() + " in query")
	}
	if row.Next() {
		err = row.Scan(&(user.Uid), &(user.Mail), &(user.Passwd), &(user.Fname),
			&(user.Lname), &(user.Age), &(user.Gender), &(user.Orientation),
			&(user.Biography), &(user.AvaPhotoID), &(user.AccType), &(user.Rating))
		if err != nil {
			return user, err
		}
	}
	return user, nil
}

func (conn *ConnDB) GetLoggedUsers(uid []int) ([]config.User, error) {
	var users = []config.User{}
	var user config.User

	if len(uid) == 0 {
		return users, nil
	}

	query := "SELECT * FROM users WHERE uid IN ("
	length := len(uid)
	for i := 1; i <= length; i++ {
		query += "$" + strconv.Itoa(i) + ", "
	}
	tmp := []byte(query)
	tmp = tmp[:(len(tmp) - 2)]
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
		err = rows.Scan(&(user.Uid), &(user.Mail), &(user.Passwd), &(user.Fname),
			&(user.Lname), &(user.Age), &(user.Gender), &(user.Orientation),
			&(user.Biography), &(user.AvaPhotoID), &(user.AccType), &(user.Rating))
		if err != nil {
			return nil, err
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
