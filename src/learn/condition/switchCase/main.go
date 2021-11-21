package main

import "fmt"

func main() {
	fmt.Println(payDebt(5000, 5000))
	fmt.Println(checkRGB("orange"))

	name := "Rifqi"
	switch name {
	case "Rifqi":
		fmt.Println("Hello Rifqi")
	case "Aziz":
		fmt.Println("Hello Aziz")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// Short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Without condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")		
	}
}

func payDebt(balance int, debt int) (message string) {
	switch {
	case balance > debt:
		message = "Uang kamu kebanyakan"
	case balance == debt:
		message = "Hutang kamu lunas"
	case balance < debt:
		message = "Uang kamu kurang"
	}
	return
}

func checkRGB(color string) (message string) {
	switch color {
	case "red":
		message = "Warna termasuk RGB"
	case "green":
		message = "Warna termasuk RGB"
	case "blue":
		message = "Warna termasuk RGB"
	default:
		message = "Warna tidak termasuk RGB"
	}
	return
}