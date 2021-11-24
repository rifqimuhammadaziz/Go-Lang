package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	alamat := Address{
		City: "Tegal",
		Province: "Jawa Tengah",
		Country: "",
	}
	fmt.Println(alamat) // origin variable not changed
	ChangeAddressToIndonesia(&alamat) // change value
	fmt.Println(alamat) // origin variable changed
}

func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}