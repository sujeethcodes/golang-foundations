package main

import "fmt"

// shorthandVariable shows how to declare variables using shorthand syntax inside a function.
func shorthandVariable() {
	// Single variable shorthand declaration
	name := "sujeeth"
	fmt.Println(name)

	// Multiple variable shorthand declaration in one line
	age, email, job := 26, "sujeethcode@gmail.com", "Software Developer"
	fmt.Println(age, email, job)

	// Note:
	// Shorthand variable declaration (:=) only works inside functions,
	// not at the package level.
}

// varDeclare shows how to declare variables using the var keyword.
func varDeclare() {
	// Single variable declaration using var
	var name = "sujeeth"
	fmt.Println(name)

	// Multiple variable declaration in one line using var
	var age, email, job = 26, "sujeethcode@gmail.com", "Software Developer"
	fmt.Println(age, email, job)

	// Note:
	// var declaration can be used both inside and outside functions,
	// and optionally with explicit type.
}

// constDeclare shows how to declare constants in Go.
func constDeclare() {
	// Constants are declared using 'const' keyword and cannot be changed.
	const name = "sujeeth"

	// name = "codes" Constants are declared using 'const' keyword and cannot be changed.

	fmt.Println(name)

	// Note:
	// Constants must be assigned a value at the time of declaration,
	// and their value cannot be modified later.
}

func main() {
	shorthandVariable()
	varDeclare()
	constDeclare()
}
