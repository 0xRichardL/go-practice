package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	n := len(nums)
	ahead := []map[int]bool{make(map[int]bool), make(map[int]bool)}
	for i := 1; i < n; i++ {
		ahead = append(ahead, make(map[int]bool))
		for j := i + 1; j < n; j++ {
			if _, ok := ahead[i][nums[j]]; ok {
				ahead[i][nums[j]] = false
			}
			ahead[i][nums[j]] = true
		}
	}
	d := make(map[string]bool)
	res := [][]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			s := nums[i] + nums[j]
			if ahead[j][-s] {
				arr := []int{nums[i], nums[j], -s}
				sort.Ints(arr)
				key := fmt.Sprint(arr)
				if !d[key] {
					res = append(res, arr)
				}
				d[key] = true
			}
		}
	}
	return res
}

func main() {
	type Test struct {
		Input    []int
		Expected [][]int
	}
	tests := []Test{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	}
	for _, test := range tests {
		output := threeSum(test.Input)
		fmt.Println(test.Input)
		fmt.Println(test.Expected, " ---> ", output)
	}
}
