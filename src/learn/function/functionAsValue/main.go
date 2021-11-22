package main

import "fmt"

func main() {
	greet := getMyName
	fmt.Println(greet("Rifqi"))
}

func getMyName(name string) string {
	return "Hello, " + name
}