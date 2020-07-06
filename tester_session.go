package main

import (
	"MicroProjects/session"
	"fmt"
)

func main() {
	var cook = session.CreateSession()

	fmt.Println(cook, "\n")

	item1 := cook.AddUserToSession("bsabre", 23)

	fmt.Println(cook, "\n")

	_ = cook.AddUserToSession("admin", 1)

	fmt.Println(cook, "\n")

	// fmt.Println(cook.Field[item1].Login)
	// fmt.Println(cook.Field[item2].Login)

	// fmt.Println(cook[item1].Login)
	// fmt.Println(cook[item2].Login)

	user1, err := cook.FindUserByToken(item1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user1)
	}
}
