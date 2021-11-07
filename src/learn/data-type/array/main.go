package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Rifqi"
	names[1] = "Muhammad"
	names[2] = "Aziz"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values)
	fmt.Println(len(values)) // Length of Array
}