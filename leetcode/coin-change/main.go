package main

import "math"

// 10ms Beats 81.12%
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range amount + 1 {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i >= coin && dp[i-coin] != math.MaxInt {
				if dp[i] > dp[i-coin]+1 {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
