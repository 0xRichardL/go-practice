package main

import "fmt"

// 133ms Beats 5.63%
func jump_n2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if j+nums[j] >= i {
				if dp[i] == 0 {
					dp[i] = dp[j] + 1
				} else {
					if dp[i] > dp[j]+1 {
						dp[i] = dp[j] + 1
					}
				}
			}
		}
	}
	return dp[n-1] - 1
}

// At any jumping position `i`, check for the farthest expansion for candidate `j` where `farthest < j+nums[j]`.
// At the last candidate `j` = `i+nums[i]`, count 1 jump and set the next end by best candidate expansion.
// This satisfies the concept of O(n^2) above:
// Previous best candidate can jump to index `i`.
func jump(nums []int) int {
	farthest := 0
	jumps := 0
	currentEnd := 0
	for i := 0; i < len(nums)-1; i++ {
		fmt.Println(currentEnd, farthest, jumps)
		if farthest < i+nums[i] {
			farthest = i + nums[i]
		}
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}
	return jumps
}

func main() {
	type Test struct {
		in  []int
		out int
	}
	tests := []Test{
		{[]int{2, 999, 1, 1, 1, 1, 1}, 2},
	}
	for _, t := range tests {
		fmt.Println(t.in)
		res := jump(t.in)
		fmt.Printf("%v -> %v ; %v\n", t.out, res, res == t.out)
	}
}
