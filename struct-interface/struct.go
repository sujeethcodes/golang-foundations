package main

import "fmt"

// Struct is used to group related values together (like name, age, mail, etc).
// It can store multiple fields with different data types under one type.

type Person struct {
	Name    string        // string data type
	Age     int           // int data type
	Mail    string        // string data type
	Address []interface{} // slice of any type (interface{}) – can store mixed types
}

// Method to assign values to the Person struct using pointer receiver (*Person)
// Pointer is used so that the original data is updated

// What is the use case of a method bound to a struct?
// - Helps reuse code
// - Keeps code organized
// - Saves memory (due to pointer usage)

func (u *Person) CreateUser() {
	u.Name = "sujeeth"
	u.Age = 26
	u.Mail = "sujeethcodes@gmail.com"
	u.Address = []interface{}{"chennai", "perungudi", 12345, "tamilnadu"}

	// Loop through the Address slice and print each value
	for _, each := range u.Address {
		fmt.Println("each -----", each)
	}
}

func main() {
	// Creating a Person instance with initial values
	person := Person{
		Name:    "sujeeth",
		Age:     26,
		Mail:    "sujeethcodes@gmail.com",
		Address: []interface{}{"chennai", "perungudi", 12345, "tamilnadu"},
	}

	// Print the entire struct
	fmt.Println(person)

	// Call the method to assign values again and print address
	person.CreateUser()
}
