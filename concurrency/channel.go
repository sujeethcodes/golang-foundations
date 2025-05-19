package main

import (
	"fmt"
	"time"
)

// buffer channel
func printBufferChnValue(numArr chan int) {
	for eachNum := range numArr {
		fmt.Println("buffered channel value", eachNum)
	}
}
func assignValueBufferdChn(n int) {
	numArr := make(chan int, n)
	go printBufferChnValue(numArr)
	for i := 0; i < n; i++ {
		numArr <- i
	}
	close(numArr)
}

// un buffered channel
func printUnBufferChnValue(numArr chan int) {
	for eachNum := range numArr {
		fmt.Println("un buffered channel vaue", eachNum)
	}
}
func assignValueUnBufferdChn(n int) {
	numArr := make(chan int)
	go printUnBufferChnValue(numArr)
	for i := 0; i < n; i++ {
		numArr <- i
	}
	close(numArr)
}

func main() {
	n := 10
	assignValueBufferdChn(n)
	assignValueUnBufferdChn(n)
	time.Sleep(time.Second)
}

// Channel has 2 types:
// -------------------------
// 1️⃣ Buffered Channel
// -------------------------
// - Syntax: channelName := make(chan Type, capacity)
// - Buffered channel has a fixed size (buffer limit)
// - It allows sending data **without immediate receiver**, until buffer is full
// - If buffer is full, sender will block (wait)
// - More efficient for performance

// -------------------------
// 2️⃣ Unbuffered Channel
// -------------------------
// - Syntax: channelName := make(chan Type)
// - No buffer → sender and receiver must be active **at the same time**
// - Sender waits until receiver receives the data
// - Slower than buffered but safer
