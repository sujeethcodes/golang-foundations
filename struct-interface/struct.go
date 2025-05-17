package main

import "fmt"

// Struct is used to store multiple fields of different data types in a single variable.

type Person struct {
	Name    string
	Age     int
	Mail    string
	Address interface{}
}

func main() {
	person := Person{
		Name:    "sujeeth",
		Age:     26,
		Mail:    "sujeethcodes@gmail.com",
		Address: []interface{}{"chennai", "perungudi", 12345, "tamilnadu"},
	}
	fmt.Println(person)
}
