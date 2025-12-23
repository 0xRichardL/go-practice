package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("Original: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// Various slicing operations
	s1 := s[2:5]
	fmt.Printf("s[2:5]: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))

	s2 := s[:3]
	fmt.Printf("s[:3]: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))

	s3 := s[3:]
	fmt.Printf("s[3:]: %v, len=%d, cap=%d\n", s3, len(s3), cap(s3))

	s4 := s[:]
	fmt.Printf("s[:]: %v, len=%d, cap=%d\n", s4, len(s4), cap(s4))

	// Empty slice from same position
	s5 := s[3:3]
	fmt.Printf("s[3:3]: %v, len=%d, cap=%d\n", s5, len(s5), cap(s5))
}
