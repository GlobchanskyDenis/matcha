package main

import (
	"MatchaServer/myDatabase"
	"fmt"
	"net/http"
)

func main() {
	var conn myDatabase.ConnDB
	var err error

	err = conn.Connect()

	if err != nil {
		fmt.Println("\033[31mServer cannot start -", err, "\033[m")
	} else {
		http.HandleFunc("/user/", conn.HttpHandlerUser)
		http.HandleFunc("/users/", conn.HttpHandlerUsers)
		http.HandleFunc("/auth/", conn.HttpHandlerAuth)
		http.HandleFunc("/ws/", conn.WebSocketHandlerAuth)
		fmt.Println("\033[32m" + "starting server at :3000" + "\033[m")
		http.ListenAndServe(":3000", nil)
	}
}
