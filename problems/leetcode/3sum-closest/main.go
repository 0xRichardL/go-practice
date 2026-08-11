package sumclosest

import (
	"math"
	"sort"
)

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func twoSum(nums []int, target int) int {
	left, right := 0, len(nums)-1
	minDelta := math.MaxInt
	minSum := math.MaxInt
	sum := 0
	for right > left {
		sum = nums[left] + nums[right]
		if sum == target {
			return sum
		}
		delta := absInt(sum - target)
		if minDelta > delta {
			minDelta = delta
			minSum = sum
		}
		if sum < target {
			left++
		} else {
			right--
		}
	}
	return minSum
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	delta := math.MaxInt
	sum := 0
	for i := 0; i < len(nums)-2; i++ {
		newSum := nums[i] + twoSum(nums[i+1:], target-nums[i])
		newDelta := absInt(newSum - target)
		if delta > newDelta {
			sum = newSum
			delta = newDelta
		}
	}
	return sum
}
