package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Msg string
}

// Implement the error interface
func (e *CustomError) Error() string {
	return e.Msg
}

func main() {
	var err error = &CustomError{Msg: "This is a custom error"}
	fmt.Println("Custom error:", err.Error())

	// Wrapping an error
	wrappedErr := fmt.Errorf("An error occurred: %w", err)
	fmt.Println("Wrapped error:", wrappedErr.Error())

	// Unwrapping the error
	unwrappedErr := errors.Unwrap(wrappedErr)
	fmt.Println(unwrappedErr.Error())

	// Check if the error is of type CustomError
	fmt.Println("Check error match:", errors.Is(wrappedErr, err)) // true

	secondWrappedErr := fmt.Errorf("Another layer: %w", wrappedErr)

	refErr := &CustomError{}
	errors.As(secondWrappedErr, &refErr)
	fmt.Println("Extracted error message:", refErr.Error())
}
