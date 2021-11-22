package main

import "fmt"

func main() {
	name := "Rifqi"
	counter := 0

	increment := func() {
		name = "Aziz" // changed
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter) // counter value changed in increment func
	fmt.Println(name)
}