package main

import (
	"MatchaServer/config"
	"MatchaServer/handlers"
	"MatchaServer/httpHandlers"
)

func main() {
	var conn = httpHandlers.ConnAll{}
	var err error

	print("Connecting to database\t\t")
	err = conn.ConnectAll()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Drop all tables\t\t\t")
	err = conn.Db.DropAllTables()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Drop ENUM types in db\t\t")
	err = conn.Db.DropEnumTypes()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create ENUM types in db\t\t")
	err = conn.Db.CreateEnumTypes()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create users table\t\t")
	err = conn.Db.CreateUsersTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create notif table\t\t")
	err = conn.Db.CreateNotifTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create message table\t\t")
	err = conn.Db.CreateMessageTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create photo table\t\t")
	err = conn.Db.CreatePhotoTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create known devices table\t")
	err = conn.Db.CreateDevicesTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Add admin@gmail.com user\t")
	err = conn.Db.SetNewUser("admin@gmail.com", handlers.PasswdHash("admin"))
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Set all fields to user\t\t")
	err = conn.Db.UpdateUser(config.User{1, "admin@gmail.com",
		handlers.PasswdHash("admin"),
		"admin", "superUser",
		30, "male", "hetero", "", 0,
		"confirmed", 0})
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}
}
