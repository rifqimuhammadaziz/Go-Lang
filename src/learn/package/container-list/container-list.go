package main

import (
	"container/list" // golang.org/pkg/container/list
	"fmt"
)


func main() {
	data := list.New()

	// add to end of data
	data.PushBack("Rifqi")
	data.PushBack("Muhammad")
	data.PushBack("Aziz")

	// add to first of data
	data.PushFront("Testing")

	fmt.Println(data.Front().Value) // get value first data
	fmt.Println(data.Back().Value) // get value end data

	fmt.Println("========================")

	// print all data by order (first to end / front to back)
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("========================")

	// print all data by order (end to first / back to front)
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

	fmt.Println("========================")

	// next to first data
	fmt.Println(data.Front().Next())
}