package main

import "fmt"

// 0ms Beats 100.00%
// DP state definition:
// dp[i] = true if s[:i] can be segmented into dictionary words
// Two aspects: (1) s ends with a matched word at position i, AND (2) the chain before it is continuously matched
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	// starting character "" is true, used none of the wordDict.
	dp[0] = true

	for i := 1; i < len(dp); i++ {
		for _, word := range wordDict {
			if i < len(word) {
				continue
			}
			if string(s[i-len(word):i]) == word && dp[i-len(word)] {
				dp[i] = true
			}
		}
	}
	return dp[len(s)]
}

func main() {
	type Test struct {
		inS        string
		inWordDict []string
		out        bool
	}
	tests := []Test{
		{
			"leetcode",
			[]string{"leet", "code"},
			true,
		},
	}
	for _, t := range tests {
		fmt.Printf("%+v, %+v\n", t.inS, t.inWordDict)
		res := wordBreak(t.inS, t.inWordDict)
		fmt.Printf("res = %+v\t\t=> %v\n\n", res, res == t.out)
	}
}
