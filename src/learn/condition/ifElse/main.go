package main

import (
	"fmt"
)

func main() {
	trafficLamp("blue")
	payDebt(300,300)

	// Short statement
	name := "Rifqi"
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}

func trafficLamp(color string) {
	if color == "red" {
		fmt.Println("STOP!")
	} else if color == "green" {
		fmt.Println("GO!")
	} else {
		fmt.Println("BE CAREFUL!")
	}
}

func payDebt(balance int, debt int) (message string) {
	if balance == debt {
		message := "Uangnya pas, hutang sudah lunas!"
		fmt.Println(message)
	} else if balance > debt {
		message := "Uang kamu kebanyakan, kamu dapat kembalian!"
		fmt.Println(message)
	} else {
		message := "Uang kamu kurang!"
		fmt.Println(message)
	}

	return
}
