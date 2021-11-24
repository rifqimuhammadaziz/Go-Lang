package main

import "fmt"

func main() {
	var data interface{} = Ups(3)
	fmt.Println(data)
}

// can be used multiple data type return
func Ups(i int) interface{} {
	if i == 1 {
		return 1 // integer
	} else if i == 2 {
		return true // bool
	} else {
		return "Ups" // string
	}
}