package main

import (
	"fmt"
	"sync"
)

// --------------------------
// sync.WaitGroup is used to control and wait for goroutines to finish
// --------------------------

// PrintValues reads values from the channel and prints them
func PrintValues(lanChn chan string, wg *sync.WaitGroup) {
	defer wg.Done() // defer ensures Done() is called after function completes
	for eachLan := range lanChn {
		fmt.Println(eachLan)
	}
}

func main() {
	var wg sync.WaitGroup // WaitGroup initialized

	language := []string{"golang", "javascript", "rust", "php"}

	// Create a buffered channel with size equal to number of languages
	lanChn := make(chan string, len(language))

	wg.Add(1) // Add one goroutine to WaitGroup

	// Start the goroutine
	go PrintValues(lanChn, &wg)

	// Send each language to the channel
	for _, eachlanguage := range language {
		lanChn <- eachlanguage
	}

	close(lanChn) // Close the channel so goroutine stops reading

	wg.Wait() // Wait for all goroutines to finish
}
