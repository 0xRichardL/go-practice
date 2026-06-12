package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			break
		}
		if sum < target {
			l++
		} else {
			r--
		}
	}
	return []int{l + 1, r + 1}
}

func main() {
	type Test struct {
		inNumbers []int
		inTarget  int
		output    []int
	}
	tests := []Test{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
	}
	for _, test := range tests {
		fmt.Println(test.inNumbers, test.inTarget)
		res := twoSum(test.inNumbers, test.inTarget)
		fmt.Println(test.output, " : ", res)
	}
}
