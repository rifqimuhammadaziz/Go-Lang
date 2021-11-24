package main

import "fmt"

func main() {
	var person map[string]string = newMap("Xenosty")
	if person == nil {
		fmt.Println("Data is empty.")
	} else {
		fmt.Println(person)
	}
}

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name": name,
		}
	}
}