package main

import (
	"fmt"
)

// Custom error type
type ValidationError struct {
	Field string
	Msg   string
}

// Implement the error interface
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error on %s: %s", e.Field, e.Msg)
}

func sendErr() error {
	return &ValidationError{Field: "Input", Msg: "cannot be empty"}
}

func main() {
	err := sendErr()
	if err != nil {
		if ve, ok := err.(*ValidationError); ok { // Type assertion works because err is of type error
			fmt.Printf("Field: %s, Error: %s\n", ve.Field, ve.Msg)
		} else {
			fmt.Println("General error:", err)
		}
	}
}
