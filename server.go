package main

import (
	"MatchaServer/apiServer"
	"MatchaServer/common"
	"net/http"
	"strconv"
)

func router(server *apiServer.Server) http.Handler {
	println(common.GREEN + "tracing router" + common.NO_COLOR)

	mux := http.NewServeMux()

	// GET
	mux.Handle("/interests/get/", server.GetMethodMiddleWare(
		http.HandlerFunc(server.InterestsGet)))
	mux.Handle("/ws/auth/", server.GetMethodMiddleWare(
		http.HandlerFunc(server.WebSocketAuth)))

	// POST
	mux.Handle("/user/auth/", server.PostMethodMiddleWare(
		http.HandlerFunc(server.UserAuth)))
	mux.Handle("/user/create/", server.PostMethodMiddleWare(
		http.HandlerFunc(server.UserCreate)))
	mux.Handle("/photo/download/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.PhotoDownload))))
	mux.Handle("/photo/upload/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.PhotoUpload))))
	mux.Handle("/user/get/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.UserGet))))
	mux.Handle("/search/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.Search))))

	// PATCH
	mux.Handle("/user/update/status/", server.PatchMethodMiddleWare(
		http.HandlerFunc(server.UserUpdateStatus)))
	mux.Handle("/user/update/", server.PatchMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.UserUpdate))))

	// DELETE
	mux.Handle("/user/delete/", server.DeleteMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.UserDelete))))

	serveMux := server.PanicMiddleWare(mux)

	return serveMux
}

func main() {
	server, err := apiServer.New("config/")
	if err != nil {
		println(common.RED + "Server cannot start - " + err.Error() + common.NO_COLOR)
	} else {
		mux := router(server)
		println(common.GREEN + "starting server at :" + strconv.Itoa(server.Port) + common.NO_COLOR)
		http.ListenAndServe(":"+strconv.Itoa(server.Port), mux)
		println(common.RED + "Порт " + strconv.Itoa(server.Port) + " занят" + common.NO_COLOR)
	}
}
