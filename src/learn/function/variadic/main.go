package main

import "fmt"

func main() {
	result := sumAll(10, 10, 10, 10, 10)
	fmt.Println(result)
	result = sumAll()
	fmt.Println(result)

	slice := []int{2,2,2,2,2}
	fmt.Println(sumAll(slice...))
}

// single variadic func
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// multiple variadic func
