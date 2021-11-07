package main

import (
	"fmt"
)

func main() {
	// Constant variable cannot be replaced
	const namaOrangTua = "Soedarso"
	const age = 42 // unused const not error
	fmt.Println(namaOrangTua)

	// Multiple constant
	const (
		firstName = "Rifqi"
		lastName = "Muhammad Aziz"
	)
}