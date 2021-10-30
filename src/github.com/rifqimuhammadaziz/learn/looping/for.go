package main

import (
	"fmt"
)

func main() {
	hitungDeret()
	loopNumber()
}

func hitungDeret() {
	total := 0
	for i := 1; i <= 10; i++ {
		total = total + i
	}
	
	fmt.Println("Hasil jumlah deret angka sampai n:", total)
}

// Loop function
func loopNumber() {
	// Looping 0 until <4 (5x)
	for i := 0; i < 5; i++ {
		fmt.Println("Loop angka ke - ", i)
	}
	// OR
	// i := 0
	// for i < 5 {
	// 	fmt.Println("Loop angka ke - ", i)
	// 	i++
	// }
}