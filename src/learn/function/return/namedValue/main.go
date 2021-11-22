package main

import "fmt"

func main() {
	fmt.Println(getFullName())

	namaAwal, namaTengah, namaAkhir := getFullName()
	fmt.Println(namaAwal)
	fmt.Println(namaTengah)
	fmt.Println(namaAkhir)
}

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Rifqi"
	middleName = "Muhammad"
	lastName = "Aziz"
	return
}