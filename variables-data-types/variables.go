package main

import "fmt"

// shorthandVariable how to declare variables using shorthand syntax inside a function.
func shorthandVariable() {
	// Single variable shorthand declaration
	name := "sujeeth"
	fmt.Println(name)

	// Multiple variable shorthand declaration in one line
	age, email, job := 26, "sujeethcode@gmail.com", "Software Developer"
	fmt.Println(age, email, job)

	// Note:
	// Shorthand variable declaration (:=) only works inside functions,
	// not at the package level
}

func main() {
	shorthandVariable()
}
