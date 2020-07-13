package main

import (
	"MatchaServer/myDatabase"
	"MatchaServer/handlers"
	"fmt"
)

func main() {
	var conn myDatabase.ConnDB
	var err error

	defer func() {
		_ = recover()
	}()

	fmt.Print("Connecting to database\t")
	err = conn.Connect()
	if err != nil {
		fmt.Println("\033[31m", "- error:", err, "\033[m")
		panic(err)
	} else {
		fmt.Println("\033[32m", "- done", "\033[m")
	}

	fmt.Print("Drop users table\t")
	err = conn.DropUsersTable()
	if err != nil {
		fmt.Println("\033[31m", "- error:", err, "\033[m")
		panic(err)
	} else {
		fmt.Println("\033[32m", "- done", "\033[m")
	}

	fmt.Print("Drop ENUM types in db\t")
	err = conn.DropEnumTypes()
	if err != nil {
		fmt.Println("\033[31m", "- error:", err, "\033[m")
		panic(err)
	} else {
		fmt.Println("\033[32m", "- done", "\033[m")
	}

	fmt.Print("Create ENUM types in db\t")
	err = conn.CreateEnumTypes()
	if err != nil {
		fmt.Println("\033[31m", "- error:", err, "\033[m")
		panic(err)
	} else {
		fmt.Println("\033[32m", "- done", "\033[m")
	}

	fmt.Print("Create users table\t")
	err = conn.CreateUsersTable()
	if err != nil {
		fmt.Println("\033[31m", "- error:", err, "\033[m")
		panic(err)
	} else {
		fmt.Println("\033[32m", "- done", "\033[m")
	}

	fmt.Print("Add admin@gmail.com user")
	err = conn.SetNewUser("admin@gmail.com", handlers.PasswdHash("admin"))
	if err != nil {
		fmt.Println("\033[31m", "- error:", err, "\033[m")
		panic(err)
	} else {
		fmt.Println("\033[32m", "- done", "\033[m")
	}
}
