package main

import (
	"container/ring" // golang.org/pkg/container/ring
	"fmt"
	"strconv"
)


func main() {
	var data *ring.Ring = ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}
	// print memory (pointer)
	fmt.Println(*data)

	// print data value (default function of ring for)
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}