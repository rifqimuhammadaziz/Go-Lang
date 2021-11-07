package main

import "fmt"

func main() {
	// KEY: VALUE
	person := map[string]string{
		"name":    "Rifqi Muhammad Aziz",
		"address": "Tegal",
	}

	// Add data to map
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])
	fmt.Println(len(person))

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar GoLang"
	book["author"] = "Programmer Zaman Now"
	book["wrongKey"] = "Delete"
	fmt.Println(book)
	delete(book, "wrongKey")
	fmt.Println(book)
}