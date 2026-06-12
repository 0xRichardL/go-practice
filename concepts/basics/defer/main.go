package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Defer 1")
	}()
	defer func() {
		fmt.Println("Defer 2")
	}()
	defer func() {
		fmt.Println("Defer 3")
	}()
	// Output:
	// Defer 3
	// Defer 2
	// Defer 1
}
