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

	// PUT
	mux.Handle("/user/create/", server.PutMethodMiddleWare(
		http.HandlerFunc(server.UserCreate)))
	mux.Handle("/photo/upload/", server.PutMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.PhotoUpload))))
	mux.Handle("/like/set/", server.PutMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.LikeSet))))
	mux.Handle("/ignore/set/", server.PutMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.IgnoreSet))))
	mux.Handle("/claim/set/", server.PutMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.ClaimSet))))

	// POST
	mux.Handle("/user/auth/", server.PostMethodMiddleWare(
		http.HandlerFunc(server.UserAuth)))
	mux.Handle("/photo/download/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.PhotoDownload))))
	mux.Handle("/user/get/friends/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.UserGetFriends))))
	mux.Handle("/user/get/ignored/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.UserGetIgnored))))
	mux.Handle("/user/get/claimed/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.UserGetClaimed))))
	mux.Handle("/user/get/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.UserGet))))
	mux.Handle("/search/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.Search))))
	mux.Handle("/message/get/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.MessageGet))))
	mux.Handle("/notification/get/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.NotificationGet))))
	mux.Handle("/history/scansOfMe/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.HistoryScans))))
	mux.Handle("/history/myViews/", server.PostMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.HistoryViews))))

	// PATCH
	mux.Handle("/user/update/status/", server.PatchMethodMiddleWare(
		http.HandlerFunc(server.UserUpdateStatus)))
	mux.Handle("/user/update/", server.PatchMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.UserUpdate))))

	// DELETE
	mux.Handle("/user/delete/", server.DeleteMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.UserDelete))))
	mux.Handle("/message/delete/", server.DeleteMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.MessageDelete))))
	mux.Handle("/notification/delete/", server.DeleteMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.NotificationDelete))))
	mux.Handle("/like/unset/", server.DeleteMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.LikeUnset))))
	mux.Handle("/ignore/unset/", server.DeleteMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.IgnoreUnset))))
	mux.Handle("/claim/unset/", server.DeleteMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.ClaimUnset))))
	mux.Handle("/photo/delete/", server.DeleteMethodMiddleWare(
		server.CheckAuthMiddleWare(http.HandlerFunc(server.PhotoDelete))))

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
