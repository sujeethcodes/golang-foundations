package main

import (
	"errors"
	"fmt"
)

// --------------------------------------
// Custom Error Struct
// --------------------------------------

// Error struct represents a custom error with a code and message
// Use case: When you want to include additional error information like error code
type Error struct {
	Code    int
	Message string
}

// This method makes Error struct satisfy the built-in `error` interface
// Use case: So that it can be returned like any other error
func (e Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// --------------------------------------
// Basic Error Handling Examples
// --------------------------------------

// checkValue checks if the input is 1
// Use case: Return error with dynamic value using `fmt.Errorf`
func checkValue(n int) error {
	if n != 1 {
		return fmt.Errorf("value is not 1, got %d", n)
	}
	return nil
}

// checkProgrammingLang checks if the input is "Golang"
// Use case: Simple error string using `errors.New`
func checkProgrammingLang(lang string) error {
	if lang != "Golang" {
		return errors.New("it is not Golang")
	}
	return nil
}

// checkAge checks if the age is less than 18
// Use case: Return custom error with extra context using struct
func checkAge(n int) error {
	if n < 18 {
		return Error{
			Code:    401,
			Message: fmt.Sprintf("Age is less than 18, got %d", n),
		}
	}
	return nil
}

// --------------------------------------
// Main Function - Calling and Handling Errors
// --------------------------------------

func main() {
	// Handle error returned from checkValue
	if err := checkValue(2); err != nil {
		fmt.Println("checkValue Error:", err)
	}

	// Handle error returned from checkProgrammingLang
	if err := checkProgrammingLang("Python"); err != nil {
		fmt.Println("checkProgrammingLang Error:", err)
	}

	// Handle error returned from checkAge
	if err := checkAge(13); err != nil {
		fmt.Println("checkAge Error:", err)
	}
}
