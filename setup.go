package main

import (
	"MatchaServer/myDatabase"
	"MatchaServer/handlers"
	"fmt"
	// "encoding/json"
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

	fmt.Print("Add admin user\t\t")
	err = conn.SetNewUser("admin", handlers.PasswdHash("admin"), "adminMail@gmail.com", "+7(999)888-77-66")
	if err != nil {
		fmt.Println("\033[31m", "- error:", err, "\033[m")
		panic(err)
	} else {
		fmt.Println("\033[32m", "- done", "\033[m")
	}

	// fmt.Print("Get all users\t\t")
	// _, err = conn.GetUsers()
	// if err != nil {
	// 	fmt.Println("\033[31m", "- error:", err, "\033[m")
	// 	panic(err)
	// } else {
	// 	fmt.Println("\033[32m", "- done", "\033[m")
	// }

	// fmt.Print("Get user data for auth\t")
	// user, err := conn.GetUserDataForAuth("admin", handlers.PasswdHash("admin"))
	// if err != nil {
	// 	fmt.Println("\033[31m", "- error:", err, "\033[m")
	// 	panic(err)
	// } else {
	// 	user.Passwd = ""
	// 	jsonUser, err := json.Marshal(user)
	// 	if err != nil {
	// 		fmt.Println("\033[31m", "- error:", err, "\033[m")
	// 	} else {
	// 		fmt.Println("\033[32m", "- done", "\033[m")
	// 		fmt.Println(string(jsonUser))
	// 	}
	// }
}
