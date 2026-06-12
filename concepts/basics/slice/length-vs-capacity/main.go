package main

import "fmt"

func main() {
	s := make([]int, 3, 5)
	fmt.Printf("make([]int, 3, 5): %v, len: %d, cap: %d\n", s, len(s), cap(s))

	// Length determines accessible elements
	s[0] = 1
	s[1] = 2
	s[2] = 3
	// s[3] = 4 // This would panic: index out of range

	// Append uses capacity when available
	s = append(s, 4)
	fmt.Printf("After append(4): %v, len: %d, cap: %d\n", s, len(s), cap(s))

	s = append(s, 5)
	fmt.Printf("After append(5): %v, len: %d, cap: %d\n", s, len(s), cap(s))

	// Exceeding capacity causes reallocation
	s = append(s, 6)
	fmt.Printf("After append(6): %v, len: %d, cap: %d (reallocated)\n", s, len(s), cap(s))
}
