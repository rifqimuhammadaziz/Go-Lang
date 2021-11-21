package main

import (
	"fmt"
)

func main() {
	hitungDeret()
	loopNumber()

	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	slice := []string{"Rifqi", "Muhammad", "Aziz"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Rifqi"
	person["title"] = "Programmer"
	for key, value := range person {
		fmt.Println(key, "=", value)
	}

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