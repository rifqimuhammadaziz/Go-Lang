package main

import "fmt"

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32) // convert int32 to int64
	var nilai8 int8 = int8(nilai32) // Integer overflow

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Rifqi"
	var e byte = name[0]
	var eString string = string(e)
	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}