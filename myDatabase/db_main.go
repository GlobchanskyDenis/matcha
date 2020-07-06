package myDatabase

import (
	"fmt"
	"MatchaServer/config"
	"MatchaServer/session"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
)

type UserStruct struct {
	Id          int    `json:"id,"`
	Login       string `json:"login,"`
	Passwd      string `json:"-"`
	Mail        string `json:"mail,,omitempty"`
	Phone       string `json:"phone,,omitempty"`
	Age         int    `json:"age,,omitempty"`
	Gender      string `json:"gender,,omitempty"`
	Orientation string `json:"orientation,,omitempty"`
	Biography   string `json:"orientation,,omitempty"`
	AvaPhotoID  int    `json:"avaPhotoID,,omitempty"`
	AccType		string `json:"-"`
	Rating      int    `json:"rating,"`
}

type ConnDB struct {
	db      *sql.DB
	session session.Session
}

func (db *ConnDB) Connect() error {
	var dsn string

	dsn = "user= " + config.DB_USER + " password= " + config.DB_PASSWD + " dbname= " + config.DB_NAME + " host= " + config.DB_HOST
	conn, err := sql.Open(config.DB_TYPE, dsn)
	db.db = conn
	db.session = session.CreateSession()
	return err
}

func (Conn ConnDB) TruncateUsersTable() error {
	db := Conn.db
	_, err := db.Exec("TRUNCATE TABLE users")//IF EXISTS 
	return err
}

func (Conn ConnDB) DropUsersTable() error {
	db := Conn.db
	_, err := db.Exec("DROP TABLE IF EXISTS users")
	return err
}

func (conn ConnDB) DropEnumTypes() error {
	db := conn.db
	_, err := db.Exec("DROP TYPE IF EXISTS enum_gender")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TYPE IF EXISTS enum_orientation")
	if err != nil {
		return err
	}

	_, err = db.Exec("DROP TYPE IF EXISTS acc_type")
	return err
}

func (conn ConnDB) CreateEnumTypes() error {
	db := conn.db
	_, err := db.Exec("CREATE TYPE enum_gender AS ENUM ('male', 'female', '')")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TYPE enum_orientation AS ENUM ('getero', 'bi', 'gay', '')")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TYPE acc_type AS ENUM ('confirmed', 'not confirmed', '')")
	return err
}

// не учтены location, лайки пользователей хранятся в отдельной таблице, картинки в отдельной таблице, tags в отдельной таблице, уведомления в отдельной таблице

func (conn ConnDB) CreateUsersTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE users (id SERIAL NOT NULL, " +
		"login VARCHAR(" + strconv.Itoa(config.LOGIN_MAX_LEN) + ") NOT NULL, " +
		"passwd VARCHAR(35) NOT NULL, " +
		"mail VARCHAR(" + strconv.Itoa(config.MAIL_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"PRIMARY KEY (login), " +
		"phone VARCHAR(17) NOT NULL DEFAULT '', " +
		"age INTEGER NOT NULL DEFAULT 0, " +
		"gender enum_gender NOT NULL DEFAULT '', " +
		"orientation enum_orientation NOT NULL DEFAULT '', " +
		"biography VARCHAR(300) NOT NULL DEFAULT '', " +
		"avaPhotoID INTEGER NOT NULL DEFAULT 0," +
		"account_type acc_type NOT NULL DEFAULT 'not confirmed'," +
		"rating INTEGER NOT NULL DEFAULT 0)")
	return err
}

func (db ConnDB) SetNewUser(login string, passwd string, mail string, phone string) error {
	stmt, err := db.db.Prepare("INSERT INTO users (login, passwd, mail, phone) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return fmt.Errorf("%s in preparing", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(login, passwd, mail, phone)
	if err != nil {
		return fmt.Errorf("%s in executing", err)
	}
	return nil
}

func (db ConnDB) IsUserExists(login string) (bool, error) {
	stmt, err := db.db.Prepare("SELECT id, login FROM users WHERE login=$1")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	row, err := stmt.Query(login)
	if err != nil {
		return false, err
	}
	if row.Next() {
		return true, nil
	}
	return false, nil
}

func (db ConnDB) GetUsers() ([]UserStruct, error) {
	var (
		users []UserStruct
		user  UserStruct
		err   error
		rows  *sql.Rows
	)

	rows, err = db.db.Query("SELECT id, login, passwd, phone, gender FROM users")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&(user.Id), &(user.Login), &(user.Passwd), &(user.Phone), &(user.Gender))
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (conn *ConnDB) GetUser(userId int) (UserStruct, error) {
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
