package main

import (
	"MatchaServer/apiServer"
	"MatchaServer/common"
	"MatchaServer/handlers"
	"time"
)

func main() {
	server, err := apiServer.New("config/")
	print("Connecting to database\t\t")
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Drop all tables\t\t\t")
	err = server.Db.DropAllTables()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Drop ENUM types in db\t\t")
	err = server.Db.DropEnumTypes()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Create ENUM types in db\t\t")
	err = server.Db.CreateEnumTypes()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Create users table\t\t")
	err = server.Db.CreateUsersTable()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Create notif table\t\t")
	err = server.Db.CreateNotifsTable()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Create message table\t\t")
	err = server.Db.CreateMessagesTable()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Create photo table\t\t")
	err = server.Db.CreatePhotosTable()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Create known devices table\t")
	err = server.Db.CreateDevicesTable()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Create interests table\t\t")
	err = server.Db.CreateInterestsTable()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Create likes table\t\t")
	err = server.Db.CreateLikesTable()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Create ignores table\t\t")
	err = server.Db.CreateIgnoresTable()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Create claims table\t\t")
	err = server.Db.CreateClaimsTable()
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Add admin@gmail.com user\t")
	user, err := server.Db.SetNewUser("admin@gmail.com", handlers.PassHash("admin"))
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}

	print("Set all fields to user\t\t")
	user.EncryptedPass = handlers.PassHash("admin")
	user.Fname = "admin"
	user.Lname = "superUser"
	date, err := time.Parse("2006-01-02", "1989-10-23")
	user.Birth.Time = &date
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	}
	user.Age = 30
	user.Gender = "male"
	user.Orientation = "hetero"
	err = server.Db.UpdateUser(user)
	if err != nil {
		println(common.RED + " - error: " + err.Error() + common.NO_COLOR)
		return
	} else {
		println(common.GREEN + " - done" + common.NO_COLOR)
	}
}
