package postgres

import (
	// "MatchaServer/common"
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

func (conn *ConnDB) Connect(conf *config.Sql) error {
	var dsn string

	dsn = "user=" + conf.User + " password=" + conf.Pass + " dbname=" + conf.DBName + " host=" + conf.Host + " sslmode=disable"
	db, err := sql.Open(conf.DBType, dsn)
	conn.db = db
	return err
}

func (conn *ConnDB) Close() {
	conn.db.Close()
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
	_, err = db.Exec("DROP TABLE IF EXISTS notifs")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS messages")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS photos")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS devices")
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP TABLE IF EXISTS interests")
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
		"avaID INTEGER NOT NULL DEFAULT 0, " +
		"latitude FLOAT DEFAULT 0, " +
		"longitude FLOAT DEFAULT 0, " +
		"interests VARCHAR(" + strconv.Itoa(config.INTEREST_MAX_LEN) + ")[] DEFAULT '{}'," +
		"status enum_status NOT NULL DEFAULT 'not confirmed'," +
		"rating INTEGER NOT NULL DEFAULT 0)")
	return err
}

func (conn ConnDB) CreateNotifsTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE notifs (nid SERIAL NOT NULL, " +
		"PRIMARY KEY (nid), " +
		"uidSender INT NOT NULL, " +
		"uidReceiver INT NOT NULL, " +
		"body VARCHAR(" + strconv.Itoa(config.NOTIF_MAX_LEN) + ") NOT NULL)")
	return err
}

func (conn ConnDB) CreateMessagesTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE messages (mid SERIAL NOT NULL, " +
		"PRIMARY KEY (mid), " +
		"uidSender INT NOT NULL, " +
		"uidReceiver INT NOT NULL, " +
		"body VARCHAR(" + strconv.Itoa(config.MESSAGE_MAX_LEN) + ") NOT NULL)")
	return err
}

func (conn ConnDB) CreatePhotosTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE photos (pid SERIAL NOT NULL, " +
		"PRIMARY KEY (pid), " +
		"uid INT NOT NULL, " +
		"src VARCHAR(" + strconv.Itoa(config.PHOTO_MAX_LEN) + ") NOT NULL)") ///// ЗАМЕНИТЬ В ПОСЛЕДСТВИИ НА НУЖНЫЙ ТИП ДАННЫХ !!!!!!!!!!      bytea
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

func (conn ConnDB) CreateInterestsTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE interests (id SERIAL NOT NULL, " +
		"PRIMARY KEY (id), " +
		"name VARCHAR(" + strconv.Itoa(config.INTEREST_MAX_LEN) + ") NOT NULL)")
	return err
}
