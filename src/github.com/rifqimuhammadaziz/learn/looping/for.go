package main

import (
	"fmt"
)

func main() {
	hitungDeret()
}

func hitungDeret() {
	total := 0
	for i := 1; i <= 10; i++ {
		total = total + i
	}
	
	fmt.Println("Hasil jumlah deret angka sampai n:", total)
}