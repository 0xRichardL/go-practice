package main

import "fmt"

// 0ms Beats 100.00%
// The nums is not monotonic but the sums of them do.
func minSubArrayLen(target int, nums []int) int {
	sum, l, res := 0, 0, len(nums)+1
	for r, n := range nums {
		sum += n
		for sum >= target {
			if res > r-l+1 {
				res = r - l + 1
			}
			sum -= nums[l]
			l++
		}
	}
	if res > len(nums) {
		return 0
	}
	return res
}

func main() {
	type Test struct {
		inTarget int
		inNums   []int
		out      int
	}
	tests := []Test{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{15, []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}, 2},
		{213, []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 25, 12}, 8},
	}
	for _, tc := range tests {
		fmt.Println(tc.inTarget, tc.inNums)
		res := minSubArrayLen(tc.inTarget, tc.inNums)
		fmt.Println(res, tc.out, " ---> ", res == tc.out)
	}
}
