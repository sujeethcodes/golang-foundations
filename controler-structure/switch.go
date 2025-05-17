package main

import "fmt"

// This program uses switch to check the day of the week
// and prints a message based on the value. Switch is used
// as a cleaner alternative to multiple if-else blocks.

func main() {
	day := "Tuesday"
	switch {
	case day == "Monday":
		fmt.Println("Week start ")
	case day == "Thursday":
		fmt.Println("Week mid dar")
	case day == "Friday":
		fmt.Println("Tomorrow week end")
	case day == "Saturday":
		fmt.Println("week end start")
	case day == "Sunday":
		fmt.Println("Happy sunday")
	default:
		fmt.Println("Week days")
	}

}

//  How to run this program:
// 1. Open the folder: control-structures
// 2. Run the command: go run switch.go
