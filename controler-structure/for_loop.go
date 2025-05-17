package main

import "fmt"

// This program demonstrates how to use the `for` loop in Golang.
// The loop runs from 1 to 10 and checks whether each number is even or odd.
// The format of a basic for loop in Go:
// for initialization; condition; increment {}

func main() {
	// check the even number
	n := 10
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("odd")
		}
	}
}
