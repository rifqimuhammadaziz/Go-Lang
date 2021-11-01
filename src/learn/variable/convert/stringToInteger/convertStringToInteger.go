package main

import (
	"fmt"     // print to console
	"strconv" // convert data type
)

func main() {
	tipeDataString := "10000"
	convert,_ := strconv.Atoi(tipeDataString) // Convert string to integer
	hasilAkhir := convert + 5000
	fmt.Println("Hasil convert tipe data: ", hasilAkhir)
}