package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func main() {
	var person1 Person
	person1.Name = "Rifqi"
	sayHello(person1)

	animal1 := Animal {
		Name: "Cat",
	}
	sayHello(animal1)
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}