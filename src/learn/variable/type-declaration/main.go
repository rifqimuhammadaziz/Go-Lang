package main

import "fmt"

func main() {
	type NoKTP string // NoKTP alias to string
	type Married bool // Married alias to boolean


	var noKtpRifqi NoKTP = "3376123123231231"
	var marriedStatus Married = false
	fmt.Println(noKtpRifqi)
	fmt.Println(marriedStatus)
}