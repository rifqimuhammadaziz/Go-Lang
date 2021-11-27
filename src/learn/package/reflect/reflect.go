package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}
type Person struct {
	Name   string `required:"true"`
	Age    int
	Gender string
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" { // if tag 'required' of data struct is 'true'
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Rifqi"}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())                   // get total field
	fmt.Println(sampleType.Field(0).Name)                // get field name
	fmt.Println(sampleType.Field(0).Tag.Get("required")) // get tag value
	fmt.Println(sampleType.Field(0).Tag.Get("max"))      // get tag value

	fmt.Println(IsValid(sample)) // true (sample has name "Rifqi" - required field)

	person1 := Person{"", 22, "Male"}
	fmt.Println(IsValid(person1))
}
