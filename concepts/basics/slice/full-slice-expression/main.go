package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// Normal slice: s[low:high] inherits large capacity
	s1 := s[2:5]
	fmt.Printf("\ns[2:5]: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))

	// Full slice expression: s[low:high:max] limits capacity
	s2 := s[2:5:5]
	fmt.Printf("s[2:5:5]: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	// This prevents append from modifying original slice
	s1 = append(s1, 100) // Might overwrite s[5]
	s2 = append(s2, 200) // Forces new allocation

	fmt.Printf("\nAfter appends:\n")
	fmt.Printf("Original s: %v\n", s)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}
