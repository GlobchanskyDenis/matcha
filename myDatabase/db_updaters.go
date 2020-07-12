package myDatabase

import (
	"fmt"
)

func (conn *ConnDB) UpdateUser(user UserStruct) error {
	stmt, err := conn.db.Prepare("UPDATE users SET " +
				"mail=$2, passwd=$3, fname=$4, lname=$5, age=$6, gender=$7, " +
				"orientation=$8, biography=$9, avaPhotoID=$10, accType=$11, rating=$12  " +
				"WHERE id=$1")
	if err != nil {
		return fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Mail, user.Passwd, user.Fname,
					user.Lname, user.Age, user.Gender, user.Orientation,
					user.Biography, user.AvaPhotoID, user.AccType, user.Rating)
	if err != nil {
		return fmt.Errorf("%s in executing", err)
	}
	return nil
}