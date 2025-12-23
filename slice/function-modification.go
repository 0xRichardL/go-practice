package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	fmt.Printf("Before modifyElements: %v\n", s)
	modifyElements(s)
	fmt.Printf("After modifyElements: %v\n", s)

	fmt.Printf("\nBefore appendElement: %v, len=%d, cap=%d\n", s, len(s), cap(s))
	appendElement(s)
	fmt.Printf("After appendElement: %v, len=%d, cap=%d (unchanged)\n", s, len(s), cap(s))

	fmt.Printf("\nBefore appendElementCorrect: %v\n", s)
	s = appendElementCorrect(s)
	fmt.Printf("After appendElementCorrect: %v\n", s)
}

func modifyElements(s []int) {
	if len(s) > 0 {
		s[0] = 99 // Modifies original slice
	}
}

func appendElement(s []int) {
	s = append(s, 4) // Creates new slice, doesn't affect caller
	s[0] = 88
}

func appendElementCorrect(s []int) []int {
	return append(s, 4) // Returns new slice
}
