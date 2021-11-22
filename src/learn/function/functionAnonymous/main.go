package main

import "fmt"

type Blacklist func(string) bool

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("Rifqi", blacklist)

	registerUser("root", func(name string) bool {
		return name == "admin"
	})
	registerUser("Rifqi", func(name string) bool {
		return name == "Rifqi"
	})
}

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) { // if true
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}