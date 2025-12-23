package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}
	slice1 := original[1:4]
	slice2 := original[2:]

	fmt.Printf("original: %v\n", original)
	fmt.Printf("slice1 [1:4]: %v\n", slice1)
	fmt.Printf("slice2 [2:]: %v\n", slice2)

	// Modifying slice1 affects original and slice2
	slice1[1] = 99
	fmt.Printf("\nAfter slice1[1] = 99:\n")
	fmt.Printf("original: %v\n", original)
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)

	// Appending might cause reallocation
	fmt.Println("\nAppending to slice1:")
	slice1 = append(slice1, 100)
	slice1[0] = 77
	fmt.Printf("After append and modify:\n")
	fmt.Printf("original: %v (might be affected)\n", original)
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
}
