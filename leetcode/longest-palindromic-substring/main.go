package main

import "fmt"

func checkPalindrome(s string) bool {
	// fmt.Println("arr:", s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	l := 0
	r := 0
	m := 1
	for z := 2; z <= len(s); z++ {
		// fmt.Println("z:", z)
		for i := 0; i+z-1 < len(s); i++ {
			if checkPalindrome(s[i : i+z]) {
				if z > m {
					m = z
					l = i
					r = i + z - 1
				}
				break
			}
		}
	}
	return s[l : r+1]
}

func main() {
	type TestCase struct {
		input    string
		expected string
	}
	tests := []TestCase{
		{"bb", "bb"},
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
	}
	for _, test := range tests {
		fmt.Println(test.input)
		fmt.Println(test.expected, " -> ", longestPalindrome(test.input))
		fmt.Println("- - - - -")
	}
}
