package main

import "fmt"

// Requirement:
// Given two strings `target` and `source``, return true if `target` can be constructed by using the letters from `source` and false otherwise.
// Example:
// target: andpad - source: "andpizzadone" -> true
// target: andpad - soruce: "anhpadz" -> false
// target: andpad - soruce: "anhpaddz" -> true

// constraint:
// target and source contains alphabet lowercase only

func canConstruct(target string, source string) bool {
	m := make(map[rune]int)
	for _, r := range source {
		m[r]++
	}
	for _, r := range target {
		m[r]--
		if m[r] < 0 {
			return false
		}
	}
	return true
}

func main() {
	type Test struct {
		Input    []string
		Expected bool
	}
	tests := []Test{
		{[]string{"andpad", "andpizzadone"}, true},
		{[]string{"andpad", "anhpadz"}, false},
	}
	for _, test := range tests {
		output := canConstruct(test.Input[0], test.Input[1])
		fmt.Println(test.Input)
		fmt.Println(test.Expected, " ---> ", output)
	}
}
