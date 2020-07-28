package main

import (
	"MatchaServer/config"
	"MatchaServer/httpHandlers"
	"net/http"
)

func main() {
	var server = httpHandlers.Server{}

	err := server.ConnectAll()
	if err != nil {
		println(config.RED + "Server cannot start - " + err.Error() + config.NO_COLOR)
	} else {
		http.HandleFunc("/user/reg/", server.HttpHandlerUserReg)
		http.HandleFunc("/user/auth/", server.HttpHandlerUserAuth)
		http.HandleFunc("/user/update/status/", server.HttpHandlerUserUpdateStatus)
		http.HandleFunc("/user/update/", server.HttpHandlerUserUpdate)
		http.HandleFunc("/user/delete/", server.HttpHandlerUserDelete)
		http.HandleFunc("/search/", server.HttpHandlerSearch)
		http.HandleFunc("/ws/auth/", server.WebSocketHandlerAuth)

		println(config.GREEN + "starting server at :3000" + config.NO_COLOR)
		http.ListenAndServe(":3000", nil)
	}
}
