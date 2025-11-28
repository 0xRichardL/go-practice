package main

import (
	"fmt"
)

// 2ms Beats 16.18%
// 9.85MB Beats 28.60%
func rotate1(nums []int, k int) {
	if k == 0 {
		return
	}
	k %= len(nums)
	shifted := make([]int, k)
	for i := 0; i < k; i++ {
		shifted[i] = nums[len(nums)-k+i]
	}
	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	for i := 0; i < k; i++ {
		nums[i] = shifted[i]
	}
}

// 0ms Beats 100.00%
// 9.54MB Beats 50.22%
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	n := len(nums)
	k %= n
	k = n - k
	for i := 0; i < k/2; i++ {
		t := nums[i]
		nums[i] = nums[k-i-1]
		nums[k-i-1] = t
	}
	// fmt.Println("#1", nums)
	for i := 0; i < (n-k)/2; i++ {
		t := nums[k+i]
		nums[k+i] = nums[n-i-1]
		nums[n-i-1] = t
	}
	// fmt.Println("#2", nums)
	for i := 0; i < n/2; i++ {
		t := nums[i]
		nums[i] = nums[n-i-1]
		nums[n-i-1] = t
	}
	// fmt.Println("#3", nums)
}

func main() {
	type Input struct {
		nums []int
		k    int
	}
	type Test struct {
		input  Input
		output []int
	}
	tests := []Test{
		{
			Input{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			Input{[]int{1, 2, 3, 4, 5, 6, 7}, 0},
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			Input{[]int{-1, -100, 3, 99}, 2},
			[]int{3, 99, -1, -100},
		},
	}
	for _, test := range tests {
		fmt.Println(test.input)
		rotate(test.input.nums, test.input.k)
		fmt.Println(test.output, " : ", test.input, " ---> ", fmt.Sprint(test.output) == fmt.Sprint(test.input.nums))
	}
}
