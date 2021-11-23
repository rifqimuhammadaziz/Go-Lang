package main

import "fmt"

// DEFER : continue to execute the defer function inside the common function if any error occured
func main() {
	runApplication(0)
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() // executable function even has error
	fmt.Println("Run Application!")
	result := 10 / value
	fmt.Println("Result", result)
}