package main

import (
	"fmt"
	"time"
)


func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println("Tahun:", now.Year())
	fmt.Println("Bulan:", now.Month())
	fmt.Println("Hari:", now.Day())
	fmt.Println("Jam:", now.Hour())
	fmt.Println("Menit:", now.Minute())
	fmt.Println("Detik:", now.Second())
	fmt.Println("Milidetik:", now.Nanosecond())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02" // default format	
	parse, _ := time.Parse(layout, "2020-10-01")
	fmt.Println(parse)
}