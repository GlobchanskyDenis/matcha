package main

import (
	"encoding/json"
	"fmt"
)

type reg struct {
	Id     int
	Login  string
	Passwd string
}

func main() {
	jsonString1 := "{\"id\":1,\"login\":\"bsabre\",\"passwd\":\"myPasswd\"}"
	jsonString2 := "{\"login\":\"admin\"}"

	fmt.Println(jsonString1)
	fmt.Println(jsonString2)

	data := []byte(jsonString1)
	var user reg

	json.Unmarshal(data, &user)
	fmt.Println(user)

	data = []byte(jsonString2)
	// var user2 reg

	// resp, err := http.Post("localhost", "application/json", bytes.NewBuffer(bytesRepresentation))
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	var result map[string]interface{}
	// json.NewDecoder(resp.body).Decode(&result)

	json.Unmarshal(data, &result)
	fmt.Println(result)

	key := "login"
	value, isExist := result[key]
	fmt.Println(key, value, isExist)
	key = "id"
	value, isExist = result[key]
	fmt.Println(key, value, isExist)
	key = "passwd"
	value, isExist = result[key]
	fmt.Println(key, value, isExist)
	// fmt.Println(result)
}
