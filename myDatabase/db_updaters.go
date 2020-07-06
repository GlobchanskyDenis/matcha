package myDatabase

import (
	"fmt"
)

func (conn *ConnDB) UpdateLogin(userID int, login string) error {
	stmt, err := conn.db.Prepare("UPDATE users SET login=$2 WHERE id=$1")
	if err != nil {
		return fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, login)
	return err
}

func (conn *ConnDB) UpdatePasswd(userID int, passwd string) error {
	stmt, err := conn.db.Prepare("UPDATE users SET passwd=$2 WHERE id=$1")
	if err != nil {
		return fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, passwd)
	return err
}

func (conn *ConnDB) UpdateMail(userID int, mail string) error {
	stmt, err := conn.db.Prepare("UPDATE users SET mail=$2 WHERE id=$1")
	if err != nil {
		return fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, mail)
	return err
}

func (conn *ConnDB) UpdatePhone(userID int, phone string) error {
	stmt, err := conn.db.Prepare("UPDATE users SET phone=$2 WHERE id=$1")
	if err != nil {
		return fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(userID, phone)
	return err
}

func (conn *ConnDB) UpdateUser(user UserStruct) error {
	stmt, err := conn.db.Prepare("UPDATE users SET login=$2, passwd=$3, mail=$4, phone=$5 WHERE id=$1")
	if err != nil {
		return fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Login, user.Passwd, user.Mail, user.Phone)
	if err != nil {
		return fmt.Errorf("%s in executing", err)
	}
	return nil
}