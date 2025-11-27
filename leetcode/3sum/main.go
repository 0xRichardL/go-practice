package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	n := len(nums)
	res := make([][]int, 0)
	unique := make(map[[3]int]bool)
	addRes := func(n1, n2, n3 int) {
		arr := []int{n1, n2, n3}
		sort.Ints(arr)
		if _, ok := unique[[3]int{n1, n2, n3}]; !ok {
			res = append(res, arr)
			unique[[3]int{n1, n2, n3}] = true
		}
	}
	for i := 0; i < n-2; i++ {
		l, r := i+1, n-1
		for l < r {
			fmt.Println(nums[i], nums[l], nums[r])
			fmt.Println(nums[i]+nums[l], nums[r])
			if nums[i]+nums[l]+nums[r] == 0 {
				addRes(nums[i], nums[l], nums[r])
			}
			abs := nums[i] + nums[l]
			if abs < 0 {
				abs = -abs
			}
			if nums[r] > abs {
				r--
			} else {
				l++
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
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			[]int{-100, -70, -60, 110, 120, 130, 160},
			[][]int{{-100, -60, 160}, {-70, -60, 130}},
		},
	}
	for _, test := range tests {
		output := threeSum(test.Input)
		fmt.Println(test.Input)
		fmt.Println(test.Expected, " ---> ", output)
	}
}
