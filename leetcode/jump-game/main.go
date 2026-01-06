package main

import "fmt"

// 0ms Beats 100.00%
// 9.39 MB Beats 33.03%
func canJump(nums []int) bool {
	max := 0
	last := len(nums) - 1
	for i := 0; i <= max; i++ {
		if i+nums[i] > max {
			max = i + nums[i]
		}
		if max >= last {
			return true
		}
	}
	return max >= last
}

// 0ms Beats 100.00%
// 8.93MB Beats 64.41%
func canJump_no_last(nums []int) bool {
	max := 0
	for i := 0; i <= max; i++ {
		if i+nums[i] > max {
			max = i + nums[i]
		}
		if max >= len(nums)-1 {
			return true
		}
	}
	return max >= len(nums)-1
}

func main() {
	type Test struct {
		in  []int
		out bool
	}
	tests := []Test{
		{[]int{2, 2, 0, 2, 0, 2, 0, 0, 2, 0}, false},
	}
	for _, t := range tests {
		fmt.Println(t.in)
		res := canJump(t.in)
		fmt.Printf("%v -> %v ; %v\n", t.out, res, res == t.out)
	}
}
