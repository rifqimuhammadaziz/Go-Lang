package main

import "fmt"

func main() {
	var a = 100
	a += 10 // a = a + 10
	fmt.Println(a)
	
	var b = 200
	b -= 20 // a = a - 20
	fmt.Println(b)

	var c = 2
	c *= 10 // c = c * 10
	fmt.Println(c)

	var d = 4
	d /= 2 // d = d / 2
	fmt.Println(d)

	var e = 10
	e %= 2 // e = e % 2
	fmt.Println(e)
}