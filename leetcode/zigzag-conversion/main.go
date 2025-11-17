package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	down := true
	r := 0
	for _, s := range s {
		rows[r] = rows[r] + string(s)
		switch r {
		case numRows - 1:
			down = false
		case 0:
			down = true
		}
		if down {
			r++
		} else {
			r--
		}
	}
	res := ""
	for _, str := range rows {
		res = res + str
	}
	return res
}

func main() {
	type Input struct {
		S       string
		NumRows int
	}
	type TestCase struct {
		input    Input
		expected string
	}
	tests := []TestCase{
		{Input{S: "PAYPALISHIRING", NumRows: 3}, "PAHNAPLSIIGYIR"},
		{Input{S: "PAYPALISHIRING", NumRows: 4}, "PINALSIGYAHRPI"},
	}
	for _, test := range tests {
		fmt.Println(test.input)
		res := convert(test.input.S, test.input.NumRows)
		fmt.Println(test.expected, "->", res, test.expected == res)
		fmt.Println("- - - - -")
	}
}
