package main

import (
	"fmt"
	"math" // golang.org/pkg/math
) 

func main() {
	// automatic closest round float64
	fmt.Println(math.Round(1.8))
	fmt.Println(math.Round(1.2))

	fmt.Println(math.Floor(1.2)) // round down
	fmt.Println(math.Ceil(1.2)) // round up

	fmt.Println(math.Max(120, 170)) // max value
	fmt.Println(math.Min(120, 170)) // min value
}