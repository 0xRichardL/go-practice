package main

import (
	"fmt"
)

// 212ms | Beats 15.50%
func dfs(s, p string, i, j int) bool {
	// fmt.Printf("%d, %d -> ", i, j)
	if i >= len(s) && j >= len(p) {
		return true
	}
	if j >= len(p) {
		return false
	}
	match := i < len(s) && (p[j] == '.' || p[j] == s[i])
	if j < len(p)-1 && p[j+1] == '*' {
		return match && dfs(s, p, i+1, j) || dfs(s, p, i, j+2)
	}
	return match && dfs(s, p, i+1, j+1)
}

// 0ms | Beats 100.0%
func dpf(s, p string) bool {
	n := len(s)
	m := len(p)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for j := 2; j <= m; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if p[j-1] == '.' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
			if p[j-1] == '*' {
				zeroM := dp[i][j-2]
				prepM := s[i-1] == p[j-2] || p[j-2] == '.'
				moreM := dp[i-1][j]
				dp[i][j] = zeroM || (prepM && moreM)
			}
		}
	}
	return dp[n][m]
}

func isMatch(s string, p string) bool {
	// return dfs(s, p, 0, 0)
	return dpf(s, p)
}

func main() {
	type TestCase struct {
		input    []string
		expected bool
	}
	tests := []TestCase{
		{[]string{"aaaaaaa", "aaaaaaa"}, true},
		{[]string{"aaaaaaa", "a*"}, true},
		{[]string{"aaaaaaa", "a.aaaaa"}, true},
		{[]string{"aaaaaaa", "a.aaaa"}, false},
		{[]string{"aab", "c*a*b"}, true},
		{[]string{"ab", ".*c"}, false},
		{[]string{"aa", "aaa"}, false},
		{[]string{"aaa", "ab*ac*a"}, true},
		{[]string{"aaa", "ab*ac*c"}, false},
		{[]string{"mississippi", "mis*is*p*."}, false},
		{[]string{"ssippp", "s*p*"}, false},
	}
	for _, test := range tests {
		fmt.Println(test.input)
		res := isMatch(test.input[0], test.input[1])
		fmt.Println(test.expected, ":", res, " --> ", test.expected == res)
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - -")
	}
}
