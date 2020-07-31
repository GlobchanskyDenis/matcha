package postgres

import (
	"MatchaServer/config"
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
)

type ConnDB struct {
	db *sql.DB
}

func New() *ConnDB {
	return &(ConnDB{})
}

func (conn *ConnDB) Connect() error {
	var dsn string

	dsn = "user=" + config.DB_USER + " password=" + config.DB_PASS + " dbname=" + config.DB_NAME + " host=" + config.DB_HOST + " sslmode=disable"
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
	_, err = db.Exec("DROP TABLE IF EXISTS devices")
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
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TYPE IF EXISTS acc_status")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TYPE IF EXISTS enum_status")
	return err
}

func (conn ConnDB) CreateEnumTypes() error {
	db := conn.db
	_, err := db.Exec("CREATE TYPE enum_gender AS ENUM ('male', 'female', '')")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE TYPE enum_orientation AS ENUM ('hetero', 'bi', 'homo', '')")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TYPE enum_status AS ENUM ('confirmed', 'not confirmed', '')")
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TYPE acc_status AS ENUM ('confirmed', 'not confirmed', '')")
	return err
}

// не учтены location, лайки пользователей хранятся в отдельной таблице, картинки в отдельной таблице, tags в отдельной таблице, уведомления в отдельной таблице

func (conn ConnDB) CreateUsersTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE users (uid SERIAL NOT NULL, " +
		"mail VARCHAR(" + strconv.Itoa(config.MAIL_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"PRIMARY KEY (mail), " +
		"encryptedPass VARCHAR(35) NOT NULL, " +
		"fname VARCHAR(" + strconv.Itoa(config.NAME_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"lname VARCHAR(" + strconv.Itoa(config.NAME_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"birth DATE DEFAULT '1955-02-03', " +
		"gender enum_gender NOT NULL DEFAULT '', " +
		"orientation enum_orientation NOT NULL DEFAULT '', " +
		"bio VARCHAR(" + strconv.Itoa(config.BIO_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"avaID INTEGER NOT NULL DEFAULT 0," +
		"status enum_status NOT NULL DEFAULT 'not confirmed'," +
		"rating INTEGER NOT NULL DEFAULT 0)")
	return err
}

func (conn ConnDB) CreateNotifTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE notif (nid SERIAL NOT NULL, " +
		"PRIMARY KEY (nid), " +
		"uidSender INT NOT NULL, " +
		"uidReceiver INT NOT NULL, " +
		"body VARCHAR(" + strconv.Itoa(config.NOTIF_MAX_LEN) + ") NOT NULL)")
	return err
}

func (conn ConnDB) CreateMessageTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE message (mid SERIAL NOT NULL, " +
		"PRIMARY KEY (mid), " +
		"uidSender INT NOT NULL, " +
		"uidReceiver INT NOT NULL, " +
		"body VARCHAR(" + strconv.Itoa(config.MESSAGE_MAX_LEN) + ") NOT NULL)")
	return err
}

func (conn ConnDB) CreatePhotoTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE photo (pid SERIAL NOT NULL, " +
		"PRIMARY KEY (pid), " +
		"uid INT NOT NULL, " +
		"body BIT(" + strconv.Itoa(config.PHOTO_MAX_LEN) + ") NOT NULL)") ///// ЗАМЕНИТЬ В ПОСЛЕДСТВИИ НА НУЖНЫЙ ТИП ДАННЫХ !!!!!!!!!!
	return err
}

func (conn ConnDB) CreateDevicesTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE devices (id SERIAL NOT NULL, " +
		"PRIMARY KEY (id), " +
		"uid INT NOT NULL, " +
		"device VARCHAR(" + strconv.Itoa(config.DEVICE_MAX_LEN) + ") NOT NULL)")
	return err
}
