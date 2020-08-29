package main

import (
	"MatchaServer/apiServer"
	"MatchaServer/common"
	"net/http"
	"strconv"
)

func main() {
	server, err := apiServer.New("config/")
	if err != nil {
		println(common.RED + "Server cannot start - " + err.Error() + common.NO_COLOR)
	} else {
		http.HandleFunc("/interests/get/", server.HandlerInterestsGet)
		http.HandleFunc("/photo/download/", server.HandlerPhotoDownload)
		http.HandleFunc("/photo/upload/", server.HandlerPhotoUpload)
		http.HandleFunc("/user/auth/", server.HandlerUserAuth)
		http.HandleFunc("/user/create/", server.HandlerUserCreate)
		http.HandleFunc("/user/get/", server.HandlerUserGet)
		http.HandleFunc("/user/update/status/", server.HandlerUserUpdateStatus)
		http.HandleFunc("/user/update/", server.HandlerUserUpdate)
		http.HandleFunc("/user/delete/", server.HandlerUserDelete)
		http.HandleFunc("/search/", server.HandlerSearch)
		http.HandleFunc("/ws/auth/", server.WebSocketHandlerAuth)

		println(common.GREEN + "starting server at :" + strconv.Itoa(server.Port) + common.NO_COLOR)
		http.ListenAndServe(":"+strconv.Itoa(server.Port), nil)
		println(common.RED + "Порт " + strconv.Itoa(server.Port) + " занят" + common.NO_COLOR)
	}
}
