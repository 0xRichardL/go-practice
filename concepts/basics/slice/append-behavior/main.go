package main

import "fmt"

func main() {
	var s []int
	fmt.Printf("Initial: len=%d, cap=%d\n", len(s), cap(s))

	for i := 0; i < 30; i++ {
		s = append(s, i)
		fmt.Printf("After append(%d): len=%d, cap=%d\n", i, len(s), cap(s))
	}
}
