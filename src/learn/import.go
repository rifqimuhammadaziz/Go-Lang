package main

import (
	"fmt" // print to console
	"learn/database"
	"learn/helper"
)

func main() {
	printSayHello := helper.SayHello("Rifqi")
	fmt.Println(printSayHello)
	fmt.Println(helper.Application)

	// ERROR (cannot access private var/func)
	// fmt.Println(helper.version)
	// printSayGoodbye := helper.sayGoodbye("Rifqi")
	// fmt.Println(printSayGoodbye)

	// package initialization (filename.Function())
	connect := database.GetDatabase()
	fmt.Println(connect)
}