package main

import "fmt"

type BlacklistCallback func(name string) bool

func registerUser(name string, blacklist BlacklistCallback) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)

	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("meong", blacklist)

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
