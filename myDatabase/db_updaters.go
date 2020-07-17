package myDatabase

import (
	// "fmt"
	"MatchaServer/config"
	"errors"
)

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
