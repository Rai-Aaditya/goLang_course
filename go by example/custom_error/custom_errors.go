package main

import (
	"errors"
	"fmt"
)

// It’s possible to define custom error types by implementing the Error() method on them. Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.

// 1. Define the custom error struct
// A custom error type usually has the suffix “Error”.
type argError struct {
	arg     int
	message string
}

// 2. This method implicitly satisfies Go's built-in `error` interface
// Adding this Error method makes argError implement the error interface.
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// f returns an `error` interface
func f(arg int) (int, error) {
	if arg == 42 {
		// We are returning a concrete *argError, but Go treats it as an `error` interface here
		return -1, &argError{arg, "can't work with it"}
	}
	// Return our custom error.
	return arg + 3, nil
}

// EXTENSION: A new function that explicitly requires the `error` interface as an argument
func logError(err error) {
	if err != nil {
		// We are explicitly calling the Error() method defined by the interface
		fmt.Println("--- Logging Error Interface ---")
		fmt.Printf("Error encountered: %s\n", err.Error())
		fmt.Println("-------------------------------")
	}
}

func main() {
	// err is of type `error` (the interface)
	_, err := f(42)

	// Explicitly using the interface by passing it to another function
	logError(err)

	// To get the custom fields (arg, message) back out of the generic `error` interface,
	// we use standard Go's `errors.As` function.
	if ae, ok := errors.AsType[*argError](err); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
	// Example of a normal error not matching argError
	_, normalErr := f(10) // Returns nil error
	if normalErr == nil {
		fmt.Println("\nf(10) ran successfully without using the error interface.")
	}
}
