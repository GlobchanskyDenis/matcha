package main

import (
	"fmt"
	"errors"
	"database/sql"
	_ "github.com/lib/pq"
	"strconv"
	"strings"
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

func (conn *ConnDB) CreateArrTable() error {
	db := conn.db
	_, err := db.Exec("CREATE TABLE test1 (id SERIAL, arr INTEGER[] NOT NULL)")
	return err
}

func (conn *ConnDB) DropArrTable() error {
	_, err := conn.db.Exec("DROP TABLE IF EXISTS test1")
	return err
}

func (conn *ConnDB) InsertArrVal(arr []int) (int, error) {
	var id int
	var values string
	for _, val := range arr {
		values += strconv.Itoa(val) + ", "
	}
	if len(values) > 2 {
		values = string(values[:(len(values) - 2)])
	}
	print("values: ")
	println(values)
	stmt, err := conn.db.Prepare("INSERT INTO test1 (arr) VALUES ('{" + values + "}') RETURNING id")
	if err != nil {
		return 0, errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	err = stmt.QueryRow().Scan(&id)
	if err != nil {
		return 0, errors.New(err.Error() + " in executing")
	}
	return id, nil
}

func (conn *ConnDB) GetArrByID(id int) (string, error) {
	var arr string
	stmt, err := conn.db.Prepare("SELECT arr FROM test1 WHERE id=$1")
	if err != nil {
		return "", errors.New(err.Error() + " in preparing")
	}
	defer stmt.Close()
	row, err := stmt.Query(id)
	if err != nil {
		return "", errors.New(err.Error() + " in query")
	}
	if row.Next() {
		err = row.Scan(&arr)
		if err != nil {
			return "", err
		}
	}
	return arr, nil
}


func main() {
	var (
		arr []int
		ret string
		strArr []string
		intArr []int
	)
	db := New()
	err := db.Connect()
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}
	err = db.DropArrTable()
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}
	err = db.CreateArrTable()
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}
	arr = append(arr, 3, 7, 18, 1, -146)
	fmt.Println("arr=", arr)
	id, err := db.InsertArrVal(arr)
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}
	ret, err = db.GetArrByID(id)
	if err != nil {
		print("Error: ")
		println(err.Error())
		return
	}
	// fmt.Printf("ret=%s\n", ret)
	ret = string(ret[1:len(ret)-1])
	// fmt.Printf("ret=%s\n", ret)
	strArr = strings.Split(ret, ",")
	// fmt.Printf("ret=%q\n", strArr)
	for _, strItem := range strArr {
		intItem, err := strconv.Atoi(strItem)
		if err != nil {
			print("Error: ")
			println(err.Error())
			return
		}
		intArr = append(intArr, intItem)
	}
	// fmt.Printf("ret=%V\n", intArr)
}