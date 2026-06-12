package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("Zero value slice: %v, len=%d, cap=%d, nil=%t\n", s, len(s), cap(s), s == nil)

	// Zero value slice is safe to use
	s = append(s, 1, 2, 3)
	fmt.Printf("After append: %v, len=%d, cap=%d, nil=%t\n", s, len(s), cap(s), s == nil)

	// len() and cap() work on nil slices
	var nilSlice []string
	fmt.Printf("\nNil slice: len=%d, cap=%d\n", len(nilSlice), cap(nilSlice))

	// Range over nil slice works (0 iterations)
	count := 0
	for range nilSlice {
		count++
	}
	fmt.Printf("Range over nil slice: %d iterations\n", count)
}
