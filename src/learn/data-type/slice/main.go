package main

import "fmt"

func main() {
	// Undefined array length
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7] // Index ke 4 sampai 6
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // Length of slice
	fmt.Println(cap(slice1)) // Capacity of slice (first of slice until end of array)

	months[5] = "Diubah" // Array changed
	fmt.Println(slice1)

	slice1[0] = "MeiDiedit" // Slice changed
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Edit")
	fmt.Println(slice3)
	slice3[1] = "DesemberEdit"
	fmt.Println(slice3)
	fmt.Println(slice2)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Rifqi"
	newSlice[1] = "Muhammad"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice) // Copy slice from newSlice to copySlice
	fmt.Println(copySlice)

	iniArray := [5]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}