package main

import "fmt"

const (
	Exact = iota
	More
)

func isMatch(s string, p string) bool {
	// Build regex groups.
	type RegexGroup struct {
		Type  int
		Value string
	}
	rgs := make([]RegexGroup, 0)
	j := 0
	for i, r := range p {
		if r == '*' {
			if j < i-1 {
				rgs = append(rgs, RegexGroup{
					Type:  Exact,
					Value: p[j : i-1],
				})
			}
			rgs = append(rgs, RegexGroup{
				Type:  More,
				Value: p[i-1 : i+1],
			})
			j = i + 1
		}
	}
	if j < len(p) {
		rgs = append(rgs, RegexGroup{
			Type:  Exact,
			Value: p[j:],
		})
	}
	// fmt.Println(rgs)
	// Construct dynamic programming 2D table.
	d := make([][]bool, len(rgs))
	for i := range d {
		d[i] = make([]bool, len(s))
	}
	for i, rg := range rgs {
		if rg.Type == Exact {
			if len(rg.Value) > len(s) {
				return false
			}
			j := 0
			end := len(s) - len(rg.Value) + 1
			if i == 0 {
				end = 1
			}
			if i == len(rgs)-1 {
				j = len(s) - len(rg.Value)
			}
			for ; j < end; j++ {
				rgm := true
				for k, r := range rg.Value {
					if r == '.' {
						continue
					}
					if r != rune(s[j+k]) {
						rgm = false
						break
					}
				}
				if rgm {
					for k := j; k < j+len(rg.Value); k++ {
						d[i][k] = true
					}
				}
			}
			continue
		}
		if rg.Value[0] == '.' {
			for j := range s {
				d[i][j] = true
			}
			continue
		}
		for j, r := range s {
			if rune(rg.Value[0]) == r {
				d[i][j] = true
			}
		}
	}
	// fmt.Println(d)
	i := len(rgs) - 1
	j = len(s) - 1
	if d[i][j] == false {
		return false
	}
	for i >= 0 {
		// fmt.Println(i, ":", j)
		if j > 0 && i > 0 && d[i-1][j-1] {
			j--
			i--
			continue
		}
		if j > 0 && d[i][j-1] {
			j--
		} else {
			break
		}
	}
	return j == 0
}

func main() {
	type TestCase struct {
		input    []string
		expected bool
	}
	tests := []TestCase{
		// {[]string{"aaaaaaa", "aaaaaaa"}, true},
		// {[]string{"aaaaaaa", "a*"}, true},
		// {[]string{"aaaaaaa", "a.aaaaa"}, true},
		// {[]string{"aaaaaaa", "a.aaaa"}, false},
		// {[]string{"aab", "c*a*b"}, true},
		// {[]string{"ab", ".*c"}, false},
		{[]string{"aa", "aaa"}, false},
		{[]string{"aaa", "ab*ac*a"}, false},
	}
	for _, test := range tests {
		fmt.Println(test.input)
		res := isMatch(test.input[0], test.input[1])
		fmt.Println(test.expected, ":", res, " --> ", test.expected == res)
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - -")
	}
}
