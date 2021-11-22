package main

import "fmt"

func main() {
	funcLoop := factorialLoop(5)
	fmt.Println(funcLoop)
	fmt.Println(5 * 4 * 3 * 2 * 1) // proof

	funcRecursive := factorialRecursive(5)
	fmt.Println(funcRecursive)
}

// basic function
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

// recursive function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value - 1)
	}
}