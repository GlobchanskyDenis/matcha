package main

import (
	"MatchaServer/myDatabase"
	"MatchaServer/config"
	"net/http"
)

func main() {
	var conn myDatabase.ConnDB
	var err error

	err = conn.Connect()
	if err != nil {
		println(config.RED + "Server cannot start - " + err.Error() + config.NO_COLOR)
	} else {
		http.HandleFunc("/user/reg/", conn.HttpHandlerUserReg)
		http.HandleFunc("/user/auth/", conn.HttpHandlerUserAuth)
		http.HandleFunc("/user/update/status/", conn.HttpHandlerUserUpdateStatus)
		http.HandleFunc("/user/update/", conn.HttpHandlerUserUpdate)
		http.HandleFunc("/user/delete/", conn.HttpHandlerUserDelete)
		http.HandleFunc("/search/", conn.HttpHandlerSearch)
		http.HandleFunc("/ws/auth/", conn.WebSocketHandlerAuth)

		println(config.GREEN + "starting server at :3000" + config.NO_COLOR)
		http.ListenAndServe(":3000", nil)
	}
}
