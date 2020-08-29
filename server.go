package main

import (
	"MatchaServer/apiServer"
	"MatchaServer/common"
	"net/http"
)

func main() {
	server, err := apiServer.New()
	if err != nil {
		println(common.RED + "Server cannot start - " + err.Error() + common.NO_COLOR)
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

		println(common.GREEN + "starting server at :3000" + common.NO_COLOR)
		http.ListenAndServe(":3000", nil)
		println(common.RED + "Порт 3000 занят" + common.NO_COLOR)
	}
}
