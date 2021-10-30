package main

import (
	"fmt"
)

// Data Type: String (Global)
var namaDepan = "Rifqi "
var namaBelakang = "Muhammad Aziz"
var namaLengkap = namaDepan + namaBelakang
var universitas, fakultas = "Universitas Dian Nuswantoro", "Fakultas Ilmu Komputer"

func main() {
	// call variable
	fmt.Println("Nama Depan: " + namaDepan)
	fmt.Println("Nama Belakang: " + namaBelakang)
	fmt.Println("Nama Lengkap: " + namaLengkap)

	fmt.Println("Asal Kampus: " + universitas)
	fmt.Println("Fakultas: " + fakultas)
}