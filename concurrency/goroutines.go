package main

import (
	"fmt"
	"time"
)

// Simple goroutine function
func printValue() {
	data := map[string]string{
		"Name": "Sujeeth",
		"Role": "Software Engineer",
	}
	fmt.Println(data)
}

func main() {
	fmt.Println("Main started")

	//  "go" keyword runs the function in a separate goroutine
	go printValue()

	//  Main function itself runs in a goroutine (main goroutine)
	// So, we add Sleep to give time for other goroutines to finish
	time.Sleep(1 * time.Second) // Give enough time for goroutine to finish

	fmt.Println("Main ended")
}

// what is concurrency, when they used and, what is usecase ?

// --------------------------
// concurrency
// concurrent means its working in same time
// its help to run mutliple task execute in same time
// --------------------------

// --------------------------
// What is concurrency ?
// --------------------------
// Background tasks (example: send email, log to DB)
// Run multiple functions at the same time
// I/O operations (API calls, file reading)
// Improve performance (if tasks are independent)

// --------------------------
// Why use concurrency ? (benifits)
// Non - blocking main code (main thread won’t wait)
// helps to better scalability
// --------------------------

// --------------------------
// Usecase example ?
// api handle request 1000 request(each request as a goroutine or we can implement limit of goroutines)
// Download manager downloading multiple files at once
// Web scraper visiting multiple sites together
// --------------------------
