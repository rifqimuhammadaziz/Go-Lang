package main

import "fmt"

// method1
type Customer struct {
	Name, Address string
	Age int
}

// struct function
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	// create struct data (method1)
	var cust1 Customer
	cust1.Name = "Rifqi Muhammad Aziz"
	cust1.Address = "Tegal"
	cust1.Age = 21
	fmt.Println(cust1)
	fmt.Println(cust1.Name)
	fmt.Println(cust1.Address)
	fmt.Println(cust1.Age)

	// create struct data literals (method2) 
	cust2 := Customer {
		Name: "Rifqi",
		Address: "Indonesia",
		Age: 21,
	}
	fmt.Println(cust2)

	// create struct data literals (method3) *error if declaration of struct changed
	cust3 := Customer{"Aziz", "Indonesia", 22}
	fmt.Println(cust3)

	// create struct function data
	cust3.sayHello("Bagas")
}