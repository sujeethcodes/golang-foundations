package main

import "fmt"

func main() {
	fmt.Println("concurrency start")
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
