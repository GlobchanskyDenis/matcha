package myDatabase

import (
	"fmt"
	// "MatchaServer/config"
	// "MatchaServer/session"
	"database/sql"
	_ "github.com/lib/pq"
	// "strconv"
)

func (conn ConnDB) SearchUsersByOneFilter(filter string) ([]UserStruct, error) {
	var (
		users []UserStruct
		user  UserStruct
		err   error
		rows  *sql.Rows
	)

	rows, err = conn.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&(user.Id), &(user.Login), &(user.Passwd), &(user.Mail),
			&(user.Phone), &(user.Age), &(user.Gender), &(user.Orientation),
			&(user.Biography), &(user.AvaPhotoID), &(user.AccType), &(user.Rating))
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (conn *ConnDB) GetUserById(userId int) (UserStruct, error) {
	var (
		user UserStruct
		err error
		row *sql.Rows
	)

	stmt, err := conn.db.Prepare("SELECT * FROM users WHERE id=$1")
	if err != nil {
		return user, fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	row, err = stmt.Query(userId)
	if err != nil {
		return user, fmt.Errorf("%s in query", err)
	}
	if row.Next() {
		err = row.Scan(&(user.Id), &(user.Login), &(user.Passwd), &(user.Mail),
			&(user.Phone), &(user.Age), &(user.Gender), &(user.Orientation),
			&(user.Biography), &(user.AvaPhotoID), &(user.AccType), &(user.Rating))
		if err != nil {
			return UserStruct{}, fmt.Errorf("%s", err)
		}
	}
	return user, nil
}

func (db *ConnDB) GetUserDataForAuth(login string, passwd string) (UserStruct, error) {
	var (
		user UserStruct
		err  error
		row  *sql.Rows
	)

	stmt, err := db.db.Prepare("SELECT * FROM users WHERE (login=$1 OR mail=$1 OR phone=$1) AND passwd=$2")
	if err != nil {
		return user, fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	row, err = stmt.Query(login, passwd)
	if err != nil {
		return user, fmt.Errorf("%s in query", err)
	}
	if row.Next() {
		err = row.Scan(&(user.Id), &(user.Login), &(user.Passwd), &(user.Mail),
			&(user.Phone), &(user.Age), &(user.Gender), &(user.Orientation),
			&(user.Biography), &(user.AvaPhotoID), &(user.AccType), &(user.Rating))
		if err != nil {
			return UserStruct{}, fmt.Errorf("%s", err)
		}
	}
	return user, nil
}
