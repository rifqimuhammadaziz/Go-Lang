package main

import (
	"fmt"
	"sort"
)


type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	// create user slice
	users := []User{
		{"Rifqi", 20},
		{"Bagas", 22},
		{"Kurnia", 21},
		{"Ilham", 23},
	}

	// before sorting
	fmt.Println(users)

	// after sorting
	sort.Sort(UserSlice(users))
	fmt.Println(users)
}