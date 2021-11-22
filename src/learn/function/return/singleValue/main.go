package main

import (
	"fmt"
	"strconv"
)

func main() {
	result := getHello("")
	fmt.Println(result)

	fmt.Println("Hasil perkalian: ", multiply(2, 3))

	fmt.Println(getBiography(22, "Rifqi", "Engineer"))

	basicInfo, ageInfo := getBiography3(22, "Rifqi", "Engineer")
	fmt.Println(basicInfo)
	fmt.Println("Umurnya 10 tahun lagi adalah", ageInfo)
}

// return value is string
func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func multiply(number1 int, number2 int) int {
	return number1 * number2
}

// Create function with string return
func getBiography(age int, name string, jobs string) string {
	convertAge := strconv.Itoa(age) // Convert integer age to string return
	return name + " adalah seorang " + jobs + " saat ini berumur " + convertAge + " tahun"
}

// Create function with multiple return (string, string)
func getBiography2(age int, name string, jobs string) (string, string) {
	convertAge := strconv.Itoa(age)                                                       // Convert integer age to string return
	return name + " adalah seorang " + jobs, " saat ini berumur " + convertAge + " tahun" // multiple return has left and right valued by string (separate using coma)
}

// Create function with multiple return (string, int)
func getBiography3(age int, name string, jobs string) (bio string, ageIn10 int) {
	ageIn10 = age + 10
	bio = name + " adalah seorang " + jobs
	return
}