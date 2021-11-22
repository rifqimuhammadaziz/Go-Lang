package main

import "fmt"

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
	fmt.Println(getFullName())

	// ignore return value with _
	namaPanggilan, _, _ := getFullName()
	fmt.Println(namaPanggilan)
}

func getFullName() (string, string, string) {
	return "Rifqi", "Muhammad", "Aziz"
}