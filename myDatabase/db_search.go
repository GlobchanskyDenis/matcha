package myDatabase

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"MatchaServer/config"
)

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
		err error
		row *sql.Rows
	)

	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE uid=$1")
	if err != nil {
		return user, fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	row, err = stmt.Query(uid)
	if err != nil {
		return user, fmt.Errorf("%s in query", err)
	}
	if row.Next() {
		err = rows.Scan(&(user.Uid), &(user.Mail), &(user.Passwd), &(user.Fname),
			&(user.Lname), &(user.Age), &(user.Gender), &(user.Orientation),
			&(user.Biography), &(user.AvaPhotoID), &(user.AccType), &(user.Rating))
		if err != nil {
			return UserStruct{}, fmt.Errorf("%s", err)
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
		return user, fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	row, err = stmt.Query(mail, passwd)
	if err != nil {
		return user, fmt.Errorf("%s in query", err)
	}
	if row.Next() {
		err = rows.Scan(&(user.Uid), &(user.Mail), &(user.Passwd), &(user.Fname),
			&(user.Lname), &(user.Age), &(user.Gender), &(user.Orientation),
			&(user.Biography), &(user.AvaPhotoID), &(user.AccType), &(user.Rating))
		if err != nil {
			return UserStruct{}, fmt.Errorf("%s", err)
		}
	}
	return user, nil
}
