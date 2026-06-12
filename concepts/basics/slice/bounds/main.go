package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Printf("Original: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// Valid edge cases
	empty1 := s[0:0]
	empty2 := s[5:5]
	fmt.Printf("\ns[0:0]: %v\n", empty1)
	fmt.Printf("s[5:5]: %v\n", empty2)

	// Slicing beyond length but within capacity
	sub := s[2:4]
	fmt.Printf("\nsub = s[2:4]: %v, len=%d, cap=%d\n", sub, len(sub), cap(sub))

	// Can slice beyond sub's length but within original capacity
	extended := sub[:3]
	fmt.Printf("sub[:3]: %v (accessing hidden elements)\n", extended)

	// Invalid operations (would panic):
	fmt.Println("\nInvalid operations that would panic:")
	fmt.Println("// s[6:6]  // panic: slice bounds out of range")
	fmt.Println("// s[-1:3] // panic: slice bounds out of range")
	fmt.Println("// s[3:2]  // panic: slice bounds out of range (start > end)")
}
