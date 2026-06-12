package main

// 0ms Beats 100.00%
// 3.89MB Beats 99.57%
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	max := nums[0]
	if max < nums[1] {
		max = nums[1]
	}
	for i := 2; i < n; i++ {
		prepMax := nums[0]
		for j := 1; j < i-1; j++ {
			if prepMax < nums[j] {
				prepMax = nums[j]
			}
		}
		nums[i] += prepMax
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}
