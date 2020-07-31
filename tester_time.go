package main

import (
	"time"
	"fmt"
	"MatchaServer/config"
	"errors"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	DB_HOST = "localhost"
	DB_NAME = "matcha_db"
	DB_USER = "bsabre"
	DB_PASS = "23"
	DB_TYPE = "postgres"
)

type ConnDB struct {
	db *sql.DB
}

func New() *ConnDB {
	return &(ConnDB{})
}

func (conn *ConnDB) Connect() error {
	var dsn string

	dsn = "user=" + DB_USER + " password=" + DB_PASS + " dbname=" + DB_NAME + " host=" + DB_HOST + " sslmode=disable"
	db, err := sql.Open(DB_TYPE, dsn)
	conn.db = db
	return err
}

func (db *ConnDB) GetUserForAuth(mail string, encryptedPass string) (config.User, error) {
	var (
		user config.User
		err  error
		row  *sql.Rows
		birth string
	)

	stmt, err := db.db.Prepare("SELECT * FROM users WHERE mail=$1 AND encryptedPass=$2")
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
			&(user.Bio), &(user.AvaID), &(user.Status), &(user.Rating))
		if err != nil {
			return user, err
		}
	}
	birth = string(birth[:10])
	user.Birth, err = time.Parse("2006-01-02", birth)
	print("birth ")
	println(birth)
	if err != nil {
		return user, errors.New(err.Error() + " qwerty")
	}
	user.Age = int(time.Since(user.Birth).Hours() / 24 / 365.27)
	return user, nil
}

func main() {
	t1, err := time.Parse("2006-01-02", "1989-08-01")
	// t1, err := time.Parse("2006-01-02", "1989-07-31")
	if err != nil {
		println("Error in parse: " + err.Error())
		return
	}
	println(t1.Format("2006-01-02"))
	// dur := time.Since(t1)
	// fmt.Println(dur)
	years := int(time.Since(t1).Hours() / 24 / 365.27)
	fmt.Println(years)

	db := New()
	err = db.Connect()
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}
	user, err := db.GetUserForAuth("admin@gmail.com", "54gg9d6")
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}
	fmt.Println(user)

	fmt.Println(time.Time{})
}