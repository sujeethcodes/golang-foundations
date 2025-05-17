package main

import "fmt"

func main() {
	// -------- BREAK Example --------
	// break    → exits the entire loop when condition matches

	// When i == 5, the loop will stop completely.

	n := 10
	for i := 1; i < n; i++ {
		if i == 5 {
			break // Exit the loop when i is 5
		}
		fmt.Println("Value:", i)
	}

	// Output: Value: 1 2 3 4

	// -------- CONTINUE Example --------
	// continue → skips the current iteration and continues with the next

	// When i == 5, it will skip this iteration and move to next.

	x := 10
	for i := 1; i < x; i++ {
		if i == 5 {
			continue // Skip when i is 5
		}
		fmt.Println("Value (5 is skipped):", i)
	}

	// Output: Value (5 is skipped): 1 2 3 4 6 7 8 9
}

//  How to run this program:
// 1. Open the folder: control-structures
// 2. Run the command: go run loop_control.go
