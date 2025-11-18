package main

import "fmt"

const (
	maxInt = 2147483647
	minInt = -2147483648
)

func myAtoi(s string) int {
	res := 0
	pos := true
	i := 0
	// Skipping white space.
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	// Empty case.
	if i == len(s) {
		return 0
	}
	// start reading, sign reading.
	switch s[i] {
	case '-':
		pos = false
		i++
	case '+':
		pos = true
		i++
	}
	// Digits reading.
	for ; i < len(s); i++ {
		r := s[i]
		// Beware of result limit of int type, should be escaped when just out of it.
		if r < '0' || r > '9' || res >= maxInt {
			break
		}
		res = res*10 + int(r-'0')
	}
	if !pos {
		res = -res
		if res < minInt {
			return minInt
		}
		return res
	}
	if res > maxInt {
		return maxInt
	}
	return res
}

func main() {
	type TestCase struct {
		input    string
		expected int
	}
	tests := []TestCase{
		{"2147483647", 2147483647},
		{"9223372036854775808", 2147483647},
		{"   -42  ", -42},
		{"   abcd-42abcd  ", 0},
		{"0-1", 0},
		{"+-1", 0},
		{"+1", 1},
	}
	for _, test := range tests {
		fmt.Println(test.input)
		res := myAtoi(test.input)
		fmt.Println(test.expected, "->", res, test.expected == res)
		fmt.Println("- - - - -")
	}
}
