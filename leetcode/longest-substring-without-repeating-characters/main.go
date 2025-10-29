package main

import "fmt"

func isRepeated(s string, l, r int) (bool, int) {
	if len(s) < 1 {
		return false, -1
	}
	m := make(map[byte]int)
	for i := l; i < r; i++ {
		if p, ok := m[s[i]]; ok {
			return true, p
		}
		m[s[i]] = i
	}
	return false, -1
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	wsize := 1
	max := 0
	j := 0
	for wsize <= len(s) {
		i := j
		for (i + wsize) <= len(s) {
			repeated, pos := isRepeated(s, i, i+wsize)
			if !repeated {
				j = i
				max = wsize
				break
			}
			i = pos + 1
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
