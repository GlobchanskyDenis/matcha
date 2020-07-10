package myDatabase

import (
	"fmt"
)

func (conn *ConnDB) UpdateUser(user UserStruct) error {
	stmt, err := conn.db.Prepare("UPDATE users SET login=$2, passwd=$3, mail=$4 WHERE id=$1")
	if err != nil {
		return fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Login, user.Passwd, user.Mail)
	if err != nil {
		return fmt.Errorf("%s in executing", err)
	}
	return nil
}