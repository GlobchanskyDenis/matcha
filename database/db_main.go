package database

import (
	"MatchaServer/config"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"strconv"
)

type ConnDB struct {
	db *sql.DB
}

func (conn *ConnDB) Connect() error {
	var dsn string

	dsn = "user=" + config.DB_USER + " password=" + config.DB_PASSWD + " dbname=" + config.DB_NAME + " host=" + config.DB_HOST + " sslmode=disable"
	db, err := sql.Open(config.DB_TYPE, dsn)
	conn.db = db
	return err
}

///////////// SETUP FUNCTIONS //////////////////

func (Conn ConnDB) TruncateAllTables() error {
	db := Conn.db
	_, err := db.Exec("TRUNCATE TABLE users") //IF EXISTS
	if err != nil {
		return err
	}
	_, err = db.Exec("TRUNCATE TABLE notif") //IF EXISTS
	if err != nil {
		return err
	}
	_, err = db.Exec("TRUNCATE TABLE message") //IF EXISTS
	if err != nil {
		return err
	}
	_, err = db.Exec("TRUNCATE TABLE photo") //IF EXISTS
	if err != nil {
		return err
	}
	return nil
}

func (Conn ConnDB) DropAllTables() error {
	db := Conn.db
	_, err := db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS notif")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS message")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS photo")
	if err != nil {
		return err
	}
	return nil
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
	_, err := db.Exec("CREATE TABLE users (uid SERIAL NOT NULL, " +
		"mail VARCHAR(" + strconv.Itoa(config.MAIL_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"PRIMARY KEY (mail), " +
		"passwd VARCHAR(35) NOT NULL, " +
		"fname VARCHAR(" + strconv.Itoa(config.NAME_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"lname VARCHAR(" + strconv.Itoa(config.NAME_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"age INTEGER NOT NULL DEFAULT 0, " +
		"gender enum_gender NOT NULL DEFAULT '', " +
		"orientation enum_orientation NOT NULL DEFAULT '', " +
		"biography VARCHAR(" + strconv.Itoa(config.BIOGRAPHY_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"avaPhotoID INTEGER NOT NULL DEFAULT 0," +
		"accType acc_type NOT NULL DEFAULT 'not confirmed'," +
		"rating INTEGER NOT NULL DEFAULT 0)")
	return err
}

func (conn ConnDB) CreateNotifTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE notif (nid SERIAL NOT NULL, " +
		"PRIMARY KEY (nid), " +
		"uid INT NOT NULL, " +
		"body VARCHAR(" + strconv.Itoa(config.NOTIF_MAX_LEN) + ") NOT NULL DEFAULT '')")
	return err
}

func (conn ConnDB) CreateMessageTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE message (mid SERIAL NOT NULL, " +
		"PRIMARY KEY (mid), " +
		"uidSender INT NOT NULL, " +
		"uidReceiver INT NOT NULL, " +
		"message VARCHAR(" + strconv.Itoa(config.MESSAGE_MAX_LEN) + ") NOT NULL DEFAULT '')")
	return err
}

func (conn ConnDB) CreatePhotoTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE photo (pid SERIAL NOT NULL, " +
		"PRIMARY KEY (pid), " +
		"uid INT NOT NULL, " +
		"photo BIT(" + strconv.Itoa(config.PHOTO_MAX_LEN) + ") NOT NULL DEFAULT '')") ///// ЗАМЕНИТЬ В ПОСЛЕДСТВИИ НА НУЖНЫЙ ТИП ДАННЫХ !!!!!!!!!!
	return err
}

/////////////// MOST NEEDED FUNCTIONS ////////////////////////////

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
