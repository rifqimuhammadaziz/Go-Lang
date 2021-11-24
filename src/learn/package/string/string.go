package main

import (
	"fmt"
	"strings"
)


func main() {
	// searching string in word
	fmt.Println(strings.Contains("Kota Tegal", "Tegal")) 
	fmt.Println(strings.Contains("Kota Tegal", "City")) 

	// split
	fmt.Println(strings.Split("Rifqi Muhammad Aziz", " "))

	fmt.Println(strings.ToLower("Xenosty Theordy")) // lowercase string
	fmt.Println(strings.ToUpper("Xenosty Theordy")) // uppercase string
	fmt.Println(strings.ToTitle("xenosty theordy")) // title string

	fmt.Println(strings.Trim("   Xenosty     ", " ")) // delete string at start and end (ex. delete space)

	fmt.Println(strings.Replace("Halo Halo Halo", "Halo", "Hey", 2)) // replace 2 'Halo' to 'Hey
	fmt.Println(strings.ReplaceAll("Halo Halo Halo Halo", "Halo", "Hey")) // replace all 'Halo' to 'Hey
}