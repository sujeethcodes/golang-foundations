package main

import "fmt"

// This program demonstrates how to declare and use fixed-size arrays in Golang.
// Arrays in Go have a fixed size defined at the time of declaration.
// If you try to access an index beyond the declared size, the program will panic with an "index out of bounds" error.

func main() {
	// Declaring an array of 5 integers
	var arr = [5]int{1, 2, 3, 4, 5}

	// Printing the entire array
	fmt.Println("Array values:", arr)

	// Accessing individual elements
	fmt.Println("First element:", arr[0])
	fmt.Println("Last element:", arr[4])

	// Uncommenting the below line will cause a runtime error:
	// fmt.Println(arr[5]) // --> panic: runtime error: index out of range [5] with length 5

	// Updating array value
	arr[0] = 100
	fmt.Println("After changing the first element:", arr)

	// Note:
	// If the number of values is less than the array size,
	// remaining values will be set to zero by default.

	var paddedArr = [6]int{1, 2, 3, 4, 5}
	fmt.Println("Array with unused index:", paddedArr) // Output: [1 2 3 4 5 0]
}
