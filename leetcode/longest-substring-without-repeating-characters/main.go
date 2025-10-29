package main

import "fmt"

// Pass 987/988 test cases
// func isRepeated(s string, l, r int) (bool, int) {
// 	if len(s) < 1 {
// 		return false, -1
// 	}
// 	m := make(map[byte]int)
// 	for i := l; i < r; i++ {
// 		if p, ok := m[s[i]]; ok {
// 			return true, p
// 		}
// 		m[s[i]] = i
// 	}
// 	return false, -1
// }

// func lengthOfLongestSubstring(s string) int {
// 	if len(s) == 0 {
// 		return 0
// 	}
// 	wsize := 1
// 	max := 0
// 	j := 0
// 	for wsize <= len(s) {
// 		i := j
// 		for (i + wsize) <= len(s) {
// 			repeated, pos := isRepeated(s, i, i+wsize)
// 			if !repeated {
// 				j = i
// 				max = wsize
// 				break
// 			}
// 			i = pos + 1
// 		}
// 		wsize++
// 	}
// 	return max
// }

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	l := 0
	max := 0
	for r := 0; r < len(s); r++ {
		if pos, ok := m[s[r]]; ok && pos >= l {
			l = pos + 1
		} else {
			sl := r - l + 1
			if sl > max {
				max = sl
			}
		}
		m[s[r]] = r
	}
	return max
}

func main() {
	test := []string{
		"",
		"aab",
		"abba",
		"asjrgapa",
	}

	for _, str := range test {
		fmt.Printf("\"%s\" => %d\n", str, lengthOfLongestSubstring(str))
	}
}
