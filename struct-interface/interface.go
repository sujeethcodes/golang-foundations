package main

import "fmt"

// interface{} allows storing values of any data type in array (slice) format
// It's useful when the type is not known at compile time (dynamic data)
func createUser() interface{} {
	interfaceVal := []interface{}{
		"sujeeth",      // string
		26,             // int
		"sujeethcodes", // string
		"chennai",      // string
		"perungudi",    // string
		12345,          // int
		"tamilnadu",    // string
	}
	return interfaceVal
}

func main() {
	res := createUser()
	fmt.Println(res) // Print the interface{} value which is a slice of mixed data types
}
