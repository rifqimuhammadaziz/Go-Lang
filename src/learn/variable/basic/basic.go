package main

import "fmt"

func main() {
	var name string
	name = "Rifqi"
	fmt.Println(name)

	var (
		firstName = "Rifqi"
		lastName  = "Muhammad Aziz"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

	fullName := "Rifqi Muhammad Aziz"
	fmt.Println(fullName)

	var age int
	age = 22
	fmt.Println(age)
}