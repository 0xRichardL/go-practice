package main

import "fmt"

// 25ms Beats 51.99%
// Classic DP, finding the LIS of previous subsequences.
// Transition: dp[x] = (max of dp[:x]) + 1
func lengthOfLIS_dp(nums []int) int {
	dp := make([]int, len(nums))
	max := 1
	for i := range nums {
		for j := range dp[:i] {
			if nums[j] < nums[i] && dp[i] < dp[j] {
				dp[i] = dp[j]
			}
		}
		dp[i]++
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

// 0ms Beats 100.00%
// Greedy + Binary Search
// tails[i] = the smallest tail of all increasing subsequences with length i+1
// each nums[x] either extends tails or updates tails
// Update rule: find the first element in tails which is >= nums[x] and replace it.
func lengthOfLIS(nums []int) int {
	tails := make([]int, 0)

	bSearch := func(target int) int {
		l, r := 0, len(tails)-1
		for l < r {
			mid := l + (r-l)/2
			if tails[mid] >= target {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return l
	}

	for _, num := range nums {
		if len(tails) == 0 {
			tails = append(tails, num)
			continue
		}
		if tails[len(tails)-1] < num {
			tails = append(tails, num)
		} else {
			tails[bSearch(num)] = num
		}
	}
	return len(tails)
}

func main() {
	type Test struct {
		in  []int
		out int
	}
	tests := []Test{
		{
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
		{
			[]int{0, 1, 0, 3, 2, 3},
			4,
		},
		{
			[]int{7, 7, 7, 7, 7, 7, 7},
			1,
		},
	}
	for _, t := range tests {
		fmt.Printf("%+v\n", t.in)
		res := lengthOfLIS(t.in)
		fmt.Printf("res = %+v\t\t=> %v\n\n", res, res == t.out)
	}
}
