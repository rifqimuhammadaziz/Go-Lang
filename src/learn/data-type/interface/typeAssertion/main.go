package main

import "fmt"

func main() {
	// general type assertion
	var result interface{} = random()
	// var resultString string = result.(string) // make sure the 'result' (interface) is string
	// fmt.Println(resultString)

	// switch type assertion
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type interface")
	}
}

func random() interface{} {
	return true
}