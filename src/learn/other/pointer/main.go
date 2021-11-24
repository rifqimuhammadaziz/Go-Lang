package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Tegal", "Jawa Tengah", "Indonesia"}
	address2 := address1

	address2.City = "Slawi"

	fmt.Println(address1) // address1 not replaced
	fmt.Println(address2) // address1 copied to address 2

	var address3 Address = Address{"Semarang", "Jawa Tengah", "Indonesia"}
	var address4 *Address = &address3 // address 4 referenced to address 3

	address4.City = "Jakarta" // address3 replaced by address4

	fmt.Println(address3) 
	fmt.Println(address4)

	myAddress := Address{"Bandung", "Jawa Barat", "Indonesia"}
	myAddress2 := &myAddress
	myAddress3 := &myAddress
	*myAddress2 = Address{"Tegal", "Jawa Tengah", "Indonesia"} // myAddress replaced by myAddress2
	fmt.Println(myAddress)
	fmt.Println(myAddress2)
	fmt.Println(myAddress3)

	// function new (empty address)
	var newAddress *Address = new(Address)
	fmt.Println(newAddress) // empty
	newAddress.City = "Jakarta" // add data
	fmt.Println(newAddress)
}