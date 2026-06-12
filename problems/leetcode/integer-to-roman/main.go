package main

import "fmt"

func intToRoman(num int) string {
	varr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	carr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	r := num
	s := ""
	for i, v := range varr {
		x := r / v
		for x > 0 {
			s = s + carr[i]
			x--
		}
		r = r % v
	}
	return s
}

func main() {
	type Test struct {
		Input    int
		Expected string
	}
	tests := []Test{
		{1, "I"},
		{2, "II"},
		{4, "IV"},
		{90, "XC"},
	}
	for _, test := range tests {
		output := intToRoman(test.Input)
		fmt.Println(test.Input)
		fmt.Println(test.Expected, ":", output, " ---> ", output == test.Expected)
	}
}
