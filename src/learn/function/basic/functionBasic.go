package main

import (
	"fmt" // print to console
)


var namaOrtu string // variable kosong
var text = "This is global variable"

func main() {
	// Call function
	fmt.Println("Memanggil fungsi printNumber: ", printNumber())
}

func printNumber() int {
	return 15
}