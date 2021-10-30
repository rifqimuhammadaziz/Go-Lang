package main

import (
	"fmt"     // print to console
	"strconv" // convert data type
)

func main() {
	gaji := 10000
	fmt.Println("Gaji saat ini " + strconv.Itoa(gaji)) // Convert integer to string

	baseDataType := 5000
	convertDataType := strconv.Itoa(baseDataType) // Convert integer to string
	fmt.Println("Hasil convert tipe data: ", convertDataType)
}