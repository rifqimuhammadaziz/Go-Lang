package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	man1 := Man{"Rifqi"}
	man1.Married()
	fmt.Println(man1.Name)
}