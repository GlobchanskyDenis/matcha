package main

import (
	"MatchaServer/handlers"
	"MatchaServer/myDatabase"
	"MatchaServer/config"
)

func main() {
	var conn myDatabase.ConnDB
	var err error

	print("Connecting to database\t")
	err = conn.Connect()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Drop users table\t")
	err = conn.DropUsersTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Drop ENUM types in db\t")
	err = conn.DropEnumTypes()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create ENUM types in db\t")
	err = conn.CreateEnumTypes()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create users table\t")
	err = conn.CreateUsersTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Add admin@gmail.com user")
	err = conn.SetNewUser("admin@gmail.com", handlers.PasswdHash("admin"))
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Set all fields to user\t")
	err = conn.UpdateUser(config.User{1, "admin@gmail.com",
	handlers.PasswdHash("admin"),
	"admin", "superUser",
	30, "male", "getero", "", 0,
	"confirmed", 0})
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}
}
