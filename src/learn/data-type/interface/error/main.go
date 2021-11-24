package main

import (
	"errors"
	"fmt"
)

func main() {
	// var contohError error = errors.New("Application error...")
	// fmt.Println(contohError.Error())

	hasil, err := pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error", err)
	}
}

func pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh diisi 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}
