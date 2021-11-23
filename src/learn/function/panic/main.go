package main

import "fmt"

// PANIC : execute the function even has error
func main() {
	runApp(true)
}

func endApp() {
	fmt.Println("Application closed.")
}

func runApp(error bool) {
	defer endApp() // executable function even has error
	if error {
		panic("Application Error!")
	}
	fmt.Println("Application Running...")
}