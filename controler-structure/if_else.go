package main

import "fmt"

// This program uses if-else control structure to evaluate a student's grade
// based on the score. It checks conditions in order and prints the matching grade.

func main() {
	score := 95

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 35 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}
}

//  How to run this program:
// 1. Open the folder: control-structures
// 2. Run the command: go run if_else.go
