package main

import (

)

func main() {
	var repo = map[int]string{}

	repo[1] = "Denis"
	repo[2] = "Vasiliy"

	for key, val := range repo {
		print(key)
		print(" ")
		println(val)
		delete(repo, key)
	}

	for key, val := range repo {
		print(key)
		print(" ")
		println(val)
		// delete(repo, key)
	}
}