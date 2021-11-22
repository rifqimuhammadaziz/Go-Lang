package main

import (
	"fmt" // print to console
)

// global variable
var namaOrtu string
var text = "This is global variable"

func main() {
	// call function
	sayHello()
	fmt.Println("Memanggil fungsi printNumber: ", printNumber())
}

func sayHello() {
	fmt.Println("Hello World")
}

func printNumber() int {
	return 15
}