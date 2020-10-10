package postgres

import (
	"MatchaServer/config"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
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
	if err != nil {
		return err
	}
	conn.db = db
	return conn.db.Ping()
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
	_, err := db.Exec("DROP TABLE IF EXISTS notifs")
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
	_, err = db.Exec("DROP TABLE IF EXISTS users")
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
	_, err := db.Exec("CREATE TABLE users (uid SERIAL PRIMARY KEY, " +
		"mail VARCHAR(" + strconv.Itoa(config.MAIL_MAX_LEN) + ") UNIQUE NOT NULL DEFAULT '', " +
		"encryptedPass VARCHAR(35) NOT NULL, " +
		"fname VARCHAR(" + strconv.Itoa(config.NAME_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"lname VARCHAR(" + strconv.Itoa(config.NAME_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"birth DATE DEFAULT NULL, " +
		"gender enum_gender NOT NULL DEFAULT '', " +
		"orientation enum_orientation NOT NULL DEFAULT '', " +
		"bio VARCHAR(" + strconv.Itoa(config.BIO_MAX_LEN) + ") NOT NULL DEFAULT '', " +
		"avaID INTEGER NOT NULL DEFAULT 0, " +
		"latitude FLOAT DEFAULT 0, " +
		"longitude FLOAT DEFAULT 0, " +
		"interests VARCHAR(" + strconv.Itoa(config.INTEREST_MAX_LEN) + ")[] DEFAULT '{}'," +
		"status enum_status NOT NULL DEFAULT 'not confirmed'," +
		"rating INTEGER NOT NULL DEFAULT 0)")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE INDEX birth_idx ON users (birth)")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE INDEX sex_idx ON users (gender, orientation)")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE INDEX location_idx ON users (latitude, longitude)")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE INDEX interests_idx ON users (interests)")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE INDEX rating_idx ON users (rating)")
	if err != nil {
		return err
	}
	return nil
}

func (conn ConnDB) CreateNotifsTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE notifs (nid SERIAL PRIMARY KEY, " +
		"uidSender INT NOT NULL, " +
		"uidReceiver INT NOT NULL, " +
		"body VARCHAR(" + strconv.Itoa(config.NOTIF_MAX_LEN) + ") NOT NULL, " +
		"CONSTRAINT notifSender_fkey FOREIGN KEY (uidSender) REFERENCES users(uid) ON DELETE RESTRICT, " +
		"CONSTRAINT notifReceiver_fkey FOREIGN KEY (uidReceiver) REFERENCES users(uid) ON DELETE RESTRICT)")
	return err
}

func (conn ConnDB) CreateMessagesTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE messages (mid SERIAL PRIMARY KEY, " +
		"uidSender INT NOT NULL, " +
		"uidReceiver INT NOT NULL, " +
		"body VARCHAR(" + strconv.Itoa(config.MESSAGE_MAX_LEN) + ") NOT NULL, " +
		"CONSTRAINT message_Sender_fkey FOREIGN KEY (uidSender) REFERENCES users(uid) ON DELETE RESTRICT, " +
		"CONSTRAINT message_Receiver_fkey FOREIGN KEY (uidReceiver) REFERENCES users(uid) ON DELETE RESTRICT)")
	return err
}

func (conn ConnDB) CreatePhotosTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE photos (pid SERIAL PRIMARY KEY, " +
		"uid INT NOT NULL, " +
		"src VARCHAR(" + strconv.Itoa(config.PHOTO_MAX_LEN) + ") NOT NULL, " + ///// ЗАМЕНИТЬ В ПОСЛЕДСТВИИ НА НУЖНЫЙ ТИП ДАННЫХ !!!!!!!!!!      bytea
		"CONSTRAINT photos_fkey FOREIGN KEY (uid) REFERENCES users(uid) ON DELETE RESTRICT)")
	return err
}

func (conn ConnDB) CreateDevicesTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE devices (id SERIAL PRIMARY KEY, " +
		"uid INT NOT NULL, " +
		"device VARCHAR(" + strconv.Itoa(config.DEVICE_MAX_LEN) + ") NOT NULL, " +
		"CONSTRAINT device_fkey FOREIGN KEY (uid) REFERENCES users(uid) ON DELETE RESTRICT)")
	return err
}

func (conn ConnDB) CreateInterestsTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE interests (id SERIAL PRIMARY KEY, " +
		"name VARCHAR(" + strconv.Itoa(config.INTEREST_MAX_LEN) + ") NOT NULL)")
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE INDEX interests_table_idx ON interests (name)")
	return err
}
