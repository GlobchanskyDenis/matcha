package main

import (
	"fmt"
	// "sync"
	"MatchaServer/myDatabase"
	"net/http"
)

func main() {
	var conn myDatabase.ConnDB
	var err error

	err = conn.Connect()

	if err != nil {
		fmt.Println(err)
	} else {
		http.HandleFunc("/user/", conn.HttpHandlerUser)
		http.HandleFunc("/users/", conn.HttpHandlerUsers)
		http.HandleFunc("/auth/", conn.HttpHandlerAuth)
		fmt.Println("\033[32m" + "starting server at :3000" + "\033[m")
		http.ListenAndServe(":3000", nil)
	}
}
