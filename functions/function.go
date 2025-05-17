package main

import "fmt"

// Simple function call
func Greet() {
	fmt.Println("Hello Sujeeth")
}

// Function with parameter
func Add(a, b int) {
	fmt.Println(a + b)
}

// Function with return value
func Multiply(a, b int) int {
	return a * b
}

// Function with multiple return values
func Divide(a, b int) (int, int) {
	return a / b, a % b
}

// Named return value
func Person() (name string, age int) {
	name, age = "sujeeth", 26
	return
}

// Variadic parameter
func Variadic(val ...int) {
	sum := 0
	for _, each := range val {
		sum += each
	}
	fmt.Println(sum)

}

func main() {
	Greet()
	Add(5, 5)
	multiply := Multiply(5, 5)
	fmt.Println(multiply)
	val1, val2 := Divide(5, 5)
	fmt.Println(val1, val2)
	name, age := Person()
	fmt.Println(name, age)
	num := []int{1, 2, 3, 4, 5}
	Variadic(num...)

}

//  How to run this program:
// 1. Open the folder: functions
// 2. Run the command:  go run function.go
