package main

import (
	"fmt"
	"strconv" // golang.org/pkg/strconv
)


func main() {
	boolean, err := strconv.ParseBool("true") // convert bool to string
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}


	// BASEINT {decimal: 10, binnary: 2, octal: 8, hexa: 16}
	number, err := strconv.ParseInt("1000000", 10, 64) // parse string to decimal(10) int64
	if err == nil {
		fmt.Println(number)
		fmt.Println("number + 5:", number + 5)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10) // parse integer to string decimal
	fmt.Println(value) // value as string

	// string to integer
	valueInt, _ := strconv.Atoi("150000") // _ ignore error
	fmt.Println(valueInt+2)
}