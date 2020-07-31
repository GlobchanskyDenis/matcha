package main

import (
	"MatchaServer/apiServer"
	"MatchaServer/config"
	"MatchaServer/database/postgres"
	"MatchaServer/handlers"
	// "MatchaServer/database/fakeSql"
	"time"
)

func main() {
	print("Connecting to database\t\t")
	// server, err := apiServer.New(fakeSql.New())
	server, err := apiServer.New(postgres.New())
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Drop all tables\t\t\t")
	err = server.Db.DropAllTables()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Drop ENUM types in db\t\t")
	err = server.Db.DropEnumTypes()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create ENUM types in db\t\t")
	err = server.Db.CreateEnumTypes()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create users table\t\t")
	err = server.Db.CreateUsersTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create notif table\t\t")
	err = server.Db.CreateNotifTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create message table\t\t")
	err = server.Db.CreateMessageTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create photo table\t\t")
	err = server.Db.CreatePhotoTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Create known devices table\t")
	err = server.Db.CreateDevicesTable()
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Add admin@gmail.com user\t")
	user, err := server.Db.SetNewUser("admin@gmail.com", handlers.PassHash("admin"))
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}

	print("Set all fields to user\t\t")
	user.EncryptedPass = handlers.PassHash("admin")
	user.Fname = "admin"
	user.Lname = "superUser"
	user.Birth = time.Now()
	user.Age = 30
	user.Gender = "male"
	user.Orientation = "hetero"
	err = server.Db.UpdateUser(user)
	if err != nil {
		println(config.RED + " - error: " + err.Error() + config.NO_COLOR)
		return
	} else {
		println(config.GREEN + " - done" + config.NO_COLOR)
	}
}
