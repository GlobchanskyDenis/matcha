package main

import (
	// "time"
	"fmt"
	// "MatchaServer/config"
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

func (conn *ConnDB) CreateFloatTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE test1 (id SERIAL, latitude FLOAT NOT NULL, " +
		"longitude FLOAT NOT NULL)")
	return err
}

func (conn *ConnDB) DropFloatTable() error {
	_, err := conn.db.Exec("DROP TABLE IF EXISTS test1")
	return err
}

func (conn *ConnDB) InsertFloatVal(latitude float32, longitude float32) (int, error) {
	var id int
	stmt, err := conn.db.Prepare("INSERT INTO test1 (latitude, longitude) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return 0, errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	err = stmt.QueryRow(latitude, longitude).Scan(&id)
	if err != nil {
		return 0, errors.New(err.Error() + " in executing")
	}
	return id, nil
}

func (conn *ConnDB) GetFloatByID(id int) (float32, float32, error) {
	var latitude, longitude float32
	stmt, err := conn.db.Prepare("SELECT latitude, longitude FROM test1 WHERE id=$1")
	if err != nil {
		return 0, 0, errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	row, err := stmt.Query(id)
	if err != nil {
		return 0, 0, errors.New(err.Error() + " in query")
	}
	if row.Next() {
		err = row.Scan(&latitude, &longitude)
		if err != nil {
			return 0, 0, err
		}
	}
	return latitude, longitude, nil
}


func main() {
	var (
		latitude, longitude float32
	)
	db := New()
	err := db.Connect()
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}
	err = db.DropFloatTable()
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}
	err = db.CreateFloatTable()
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}
	latitude = 32.6
	longitude = 3.1415
	fmt.Printf("latitude=%f, longitude=%f\n", latitude, longitude)
	id, err := db.InsertFloatVal(latitude, longitude)
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}
	latitude, longitude, err = db.GetFloatByID(id)
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}

	fmt.Printf("latitude=%f, longitude=%f\n", latitude, longitude)
}