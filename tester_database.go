package main

import (
	"MatchaServer/myDatabase"
	"MatchaServer/handlers"
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	var conn myDatabase.ConnDB
	var wg = &sync.WaitGroup{}
	var mu = &sync.Mutex{}
	var getChan = make(chan interface{})
	var users []myDatabase.UserStruct
	var mainErr error

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("ERROR: ", err)
		}
	}()

	mainErr = conn.Connect()

	if mainErr != nil {
		panic(mainErr)
	}

	mainErr = conn.TruncateUsersTable()
	if mainErr != nil {
		panic(mainErr)
	}

	// return

	wg.Add(1)
	go func(conn myDatabase.ConnDB, mu *sync.Mutex, wg *sync.WaitGroup) {
		err := conn.SetNewUser("bsabre", handlers.PasswdHash("Den23@"))
		if err != nil {
			mu.Lock()
			mainErr = err
			mu.Unlock()
		}
		wg.Done()
	}(conn, mu, wg)

	wg.Add(1)
	go func(conn myDatabase.ConnDB, mu *sync.Mutex, wg *sync.WaitGroup) {
		err := conn.SetNewUser("admin", handlers.PasswdHash("admin"))
		if err != nil {
			mu.Lock()
			mainErr = err
			mu.Unlock()
		}
		wg.Done()
	}(conn, mu, wg)

	wg.Add(1)
	go func(conn myDatabase.ConnDB, mu *sync.Mutex, wg *sync.WaitGroup) {
		err := conn.SetNewUser("bsabre-c", handlers.PasswdHash("Den23@"))
		if err != nil {
			mu.Lock()
			mainErr = err
			mu.Unlock()
		}
		wg.Done()
	}(conn, mu, wg)

	wg.Wait()

	if mainErr != nil {
		panic(mainErr)
	}
	// wg.Add(1)
	go func(ch chan interface{}, conn myDatabase.ConnDB) {
		users, err := conn.GetUsers()
		if err != nil {
			mainErr = err
			ch <- []myDatabase.UserStruct{}
			return
		}
		ch <- users
	}(getChan, conn)

	users = (<-getChan).([]myDatabase.UserStruct)

	// wg.Wait()

	if mainErr != nil {
		panic(mainErr)
	}
	// for _, user := range users {
	// 	fmt.Println(user)
	// }

	var (
		answer string
	)

	jsonUsers, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	answer = string(jsonUsers)

	fmt.Println(answer)
}
