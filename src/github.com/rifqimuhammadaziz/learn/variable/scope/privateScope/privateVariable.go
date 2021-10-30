package main

import (
	"fmt"
)

func main() {
	// Only can be declared inside function
	alamatKampus, alamatDomisili := "Jl. Imam Bonjol No. 113", "Kota Tegal"
	fmt.Println("Alamat Kampus: " + alamatKampus)
	fmt.Println("Alamat Domisili: " + alamatDomisili)

	// Data Type: Integer
	tanggalLahir := "21 Mei 1999"
	var umur int = 22
	fmt.Println("Tanggal Lahir: ", tanggalLahir)
	fmt.Println("Umur: ", umur, "tahun")

	a := 10
	b := 20
	fmt.Println("Totalnya: ", a + b)
}