package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	// i is the position to write the next allowed value
	i := 0
	for j := 0; j < len(nums); j++ {
		// keep at most two duplicates: allow write if we have written < 2 items
		// or current value differs from value at i-2
		if i < 2 || nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
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
		{[]int{1, 1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7},
	}
	for _, test := range tests {
		fmt.Println(test.input)
		res := removeDuplicates(test.input)
		fmt.Println(test.output, " : ", res, " ---> ", test.output == res)
	}
}
