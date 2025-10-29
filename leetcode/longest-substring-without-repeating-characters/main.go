package main

import "fmt"

func isRepeated(s string, l, r int) bool {
	if len(s) < 1 {
		return false
	}
	m := make(map[byte]bool)
	for i := l; i < r; i++ {
		if m[s[i]] {
			return true
		}
		m[s[i]] = true
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	wsize := 1
	max := 0
	for wsize <= len(s) {
		i := 0
		for (i + wsize) <= len(s) {
			if !isRepeated(s, i, i+wsize) {
				max = wsize
				break
			}
			i++
		}
		wsize++
	}
	return max
}

func main() {
	test := []string{
		"aab",
		"asjrgapa",
	}

	for _, str := range test {
		fmt.Printf("\"%s\" => %d\n", str, lengthOfLongestSubstring(str))
	}
}
