package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original: %v\n", original)

	// Shallow copy (shares memory)
	shallow := original
	shallow[0] = 99
	fmt.Printf("\nAfter shallow[0] = 99:\n")
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Shallow: %v\n", shallow)

	// Deep copy (independent memory)
	original = []int{1, 2, 3, 4, 5}
	deep := make([]int, len(original))
	copy(deep, original)
	deep[0] = 88
	fmt.Printf("\nAfter deep copy and deep[0] = 88:\n")
	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Deep: %v\n", deep)

	// copy() with different sizes
	src := []int{1, 2, 3, 4, 5}
	dst1 := make([]int, 3)
	dst2 := make([]int, 7)

	n1 := copy(dst1, src)
	n2 := copy(dst2, src)
	fmt.Printf("\ncopy to smaller dst: copied %d elements, dst=%v\n", n1, dst1)
	fmt.Printf("copy to larger dst: copied %d elements, dst=%v\n", n2, dst2)
}
