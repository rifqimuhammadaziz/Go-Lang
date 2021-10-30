package main

import (
	"fmt" // print to console
)

var namaOrtu string // variable kosong
var text = "This is global variable"

func main() {
	// Call function without return value
	getText()
}

// Create function without return value
func getText() {
	fmt.Println(text)
}