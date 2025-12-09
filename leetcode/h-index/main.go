package main

import (
	"fmt"
	"sort"
)

// 0ms Beats 100%
// 4.2MB Beats 81.44%
func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})
	i := 0
	for ; i < len(citations); i++ {
		if citations[i] < i+1 {
			break
		}
	}
	return i
}

func main() {
	type Test struct {
		input  []int
		output int
	}
	tests := []Test{
		{[]int{3, 0, 6, 1, 5}, 3},
		{[]int{1, 3, 1}, 1},
		{[]int{0, 0, 0}, 0},
		{[]int{10, 8, 5, 4, 3}, 4},
		{[]int{25, 8, 5, 3, 3}, 3},
	}
	for _, t := range tests {
		fmt.Println(t.input)
		result := hIndex(t.input)
		fmt.Println(t.output, " : ", result, " ---> ", t.output == result)
	}
}
