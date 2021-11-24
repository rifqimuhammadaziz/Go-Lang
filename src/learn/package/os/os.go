package main

import (
	"fmt"
	"os" // golang.org/pkg/os
)

func main() {
	// get args (location of file)
	args := os.Args
	fmt.Println("Argument:", args)

	// get hostname
	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", name)
	} else {
		fmt.Println("Error:", err.Error())
	}
}