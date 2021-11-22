package main

import (
	"fmt" // print to console
)

func main() {
	namaAwal := "Rifqi"
	namaAkhir := "Muhammad Aziz"
	sayHelloTo(namaAwal, namaAkhir)
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}