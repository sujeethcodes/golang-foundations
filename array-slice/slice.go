package main

import "fmt"

// This program demonstrates how to use slices in Golang.
// Slices are dynamic arrays — they can grow or shrink in size.
// Unlike arrays, slices do not have a fixed length.

func main() {
	// Declaring a slice (no fixed size)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial slice:", slice)

	// Appending values to the slice
	// Note: append only works with slices, not arrays
	slice = append(slice, 6, 7, 8, 9)
	fmt.Println("After appending values:", slice)

	// Updating a value in the slice
	slice[0] = 100
	fmt.Println("After changing first value:", slice)

	// Accessing individual elements
	fmt.Println("First element:", slice[0])
	fmt.Println("Last element:", slice[len(slice)-1])

	// Length and Capacity
	fmt.Println("Length of slice:", len(slice))
	fmt.Println("Capacity of slice:", cap(slice))

	// Slicing a slice (getting part of a slice)
	// slice[inclusive : exclusive]
	part := slice[2:5] // from index 2 to 4
	fmt.Println("Sliced part (index 2 to 4):", part)

	// Converting an array to a slice
	arr := [3]int{1, 2, 3}
	arrayConverted := arr[:]
	fmt.Println("Array converted to slice:", arrayConverted)

	// Make concept - create slice with len=0, cap=2
	makeArr := make([]int, 0, 1)
	fmt.Println("Before append - len:", len(makeArr), "cap:", cap(makeArr))

	makeArr = append(makeArr, 1, 2, 3, 4, 5)
	fmt.Println("After append - make:", makeArr)
	fmt.Println("After append - len:", len(makeArr), "cap:", cap(makeArr))

}
