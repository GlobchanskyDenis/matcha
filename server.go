package main

import (
	"MatchaServer/apiServer"
	"MatchaServer/common"
	"net/http"
	"strconv"
)

func router(server *apiServer.Server) {
	println(common.GREEN + "tracing router" + common.NO_COLOR)

	// GET
	http.HandleFunc("/interests/get/", server.HandlerInterestsGet)
	http.HandleFunc("/ws/auth/", server.WebSocketHandlerAuth)
		
	// POST
	http.HandleFunc("/photo/download/", server.HandlerPhotoDownload)
	http.HandleFunc("/photo/upload/", server.HandlerPhotoUpload)
	http.HandleFunc("/user/auth/", server.HandlerUserAuth)
	http.HandleFunc("/user/create/", server.HandlerUserCreate)
	http.HandleFunc("/user/get/", server.HandlerUserGet)
	http.HandleFunc("/search/", server.HandlerSearch)

	// PATCH
	http.HandleFunc("/user/update/status/", server.HandlerUserUpdateStatus)
	http.HandleFunc("/user/update/", server.HandlerUserUpdate)

	// DELETE
	http.HandleFunc("/user/delete/", server.HandlerUserDelete)
}

func main() {
	server, err := apiServer.New("config/")
	if err != nil {
		println(common.RED + "Server cannot start - " + err.Error() + common.NO_COLOR)
	} else {
		router(server)
		println(common.GREEN + "starting server at :" + strconv.Itoa(server.Port) + common.NO_COLOR)
		http.ListenAndServe(":"+strconv.Itoa(server.Port), nil)
		println(common.RED + "Порт " + strconv.Itoa(server.Port) + " занят" + common.NO_COLOR)
	}
}
