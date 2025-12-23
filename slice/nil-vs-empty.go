package main

import "fmt"

func main() {
	var nilSlice []int
	emptySlice := []int{}
	makeSlice := make([]int, 0)

	fmt.Printf("nilSlice: %v, len: %d, cap: %d, nil: %t\n",
		nilSlice, len(nilSlice), cap(nilSlice), nilSlice == nil)
	fmt.Printf("emptySlice: %v, len: %d, cap: %d, nil: %t\n",
		emptySlice, len(emptySlice), cap(emptySlice), emptySlice == nil)
	fmt.Printf("makeSlice: %v, len: %d, cap: %d, nil: %t\n",
		makeSlice, len(makeSlice), cap(makeSlice), makeSlice == nil)

	// All can be appended to
	nilSlice = append(nilSlice, 1)
	emptySlice = append(emptySlice, 1)
	makeSlice = append(makeSlice, 1)
	fmt.Printf("After append(1): nil=%v, empty=%v, make=%v\n",
		nilSlice, emptySlice, makeSlice)
}
