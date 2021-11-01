package main

import "fmt"

func main() {
	fmt.Println(payDebt(5000, 5000))
	fmt.Println(checkRGB("orange"))
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