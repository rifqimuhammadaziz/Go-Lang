package main

import "fmt"

/** RECOVER : catch the message of error function
and continue the next command
*/
func main() {
	runApp(true)
	fmt.Println("This command still running...")
}

func endApp() {
	message := recover() // catch panic data
	if message != nil {
		fmt.Println("Error with message:", message)
	}
	fmt.Println("Application closed.")
}

func runApp(error bool) {
	defer endApp() // executable function even has error
	if error {
		panic("Application Error!")
	}
	fmt.Println("Application Running...")
}