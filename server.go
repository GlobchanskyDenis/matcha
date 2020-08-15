package main

import (
	"MatchaServer/apiServer"
	"MatchaServer/config"
	"MatchaServer/database/postgres"
	"net/http"
)

func main() {
	server, err := apiServer.New(postgres.New())
	if err != nil {
		println(config.RED + "Server cannot start - " + err.Error() + config.NO_COLOR)
	} else {
		http.HandleFunc("/interests/get/", server.HandlerInterestsGet)
		http.HandleFunc("/photo/download/", server.HandlerPhotoDownload)
		http.HandleFunc("/photo/upload/", server.HandlerPhotoUpload)
		http.HandleFunc("/user/create/", server.HandlerUserCreate)
		http.HandleFunc("/user/auth/", server.HandlerUserAuth)
		http.HandleFunc("/user/update/status/", server.HandlerUserUpdateStatus)
		http.HandleFunc("/user/update/", server.HandlerUserUpdate)
		http.HandleFunc("/user/delete/", server.HandlerUserDelete)
		http.HandleFunc("/search/", server.HandlerSearch)
		http.HandleFunc("/ws/auth/", server.WebSocketHandlerAuth)

		println(config.GREEN + "starting server at :3000" + config.NO_COLOR)
		http.ListenAndServe(":3000", nil)
		println(config.RED + "Порт 3000 занят" + config.NO_COLOR)
	}
}
