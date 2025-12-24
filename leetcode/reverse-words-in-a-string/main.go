package main

import "fmt"

// 5ms Beats 19.18%
func reverseWords(s string) string {
	res := ""
	l, r := 0, 0
	for ; r < len(s); r++ {
		// L set cond
		if s[r] != ' ' && s[l] == ' ' {
			l = r
		}
		if s[r] == ' ' && s[l] != ' ' {
			if len(res) == 0 {
				res = s[l:r]
			} else {
				res = s[l:r] + " " + res
			}
			l = r
		}
	}
	if s[l] != ' ' {
		if len(res) == 0 {
			res = s[l:r]
		} else {
			res = s[l:r] + " " + res
		}
	}
	return res
}

func main() {
	type Test struct {
		s   string
		ans string
	}
	tests := []Test{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}
	for _, test := range tests {
		res := reverseWords(test.s)
		fmt.Println(test.s)
		fmt.Println(" - - - > ", res, res == test.ans)
	}
}
