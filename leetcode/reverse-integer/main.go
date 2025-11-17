package main

import (
	"fmt"
)

const maxInt = 2147483647

func reverse(x int) int {
	res := 0
	positive := x >= 0
	n := x
	if x < 0 {
		n = -x
	}
	for ; n > 0; n /= 10 {
		res = res*10 + n%10
	}
	if res > maxInt {
		return 0
	}
	if positive {
		return res
	}
	return -res
}

func main() {
	type TestCase struct {
		input    int
		expected int
	}
	tests := []TestCase{
		{123, 321},
		{-123, -321},
	}
	for _, test := range tests {
		fmt.Println(test.input)
		res := reverse(test.input)
		fmt.Println(test.expected, "->", res, test.expected == res)
		fmt.Println("- - - - -")
	}
}
