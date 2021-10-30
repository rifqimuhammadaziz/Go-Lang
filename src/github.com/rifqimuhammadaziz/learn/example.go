package main

import (
	"fmt" // print to console
)

var namaOrtu string // variable kosong
var text = "This is global variable"

func main() {
	// Call function without return value
	getText()
	loop1()
}

// Create function without return value
func getText() {
	fmt.Println(text)
}

// Loop function
func loop1() {
	// Looping 0 until <4 (5x)
	for i := 0; i < 5; i++ {
		fmt.Println("Loop angka ke - ", i)
	}
	// OR
	// i := 0
	// for i < 5 {
	// 	fmt.Println("Loop angka ke - ", i)
	// 	i++
	// }
}