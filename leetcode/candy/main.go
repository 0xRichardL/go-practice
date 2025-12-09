package main

import (
	"fmt"
)

func candy(ratings []int) int {
	arr := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			arr[i] = arr[i-1] + 1
		} else {
			arr[i] = 1
		}
	}
	fmt.Println(arr)
	sum := arr[len(arr)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if arr[i] <= arr[i+1] {
				arr[i] = arr[i+1] + 1
			}
		}
		sum = sum + arr[i]
	}
	fmt.Println(arr)
	return sum
}

func main() {
	type Test struct {
		input  []int
		output int
	}
	tests := []Test{
		{[]int{1, 0, 2}, 5},
		{[]int{1, 2, 2}, 4},
	}
	for _, t := range tests {
		fmt.Println(t.input)
		result := candy(t.input)
		fmt.Println(t.output, " : ", result, " ---> ", t.output == result)
	}
}
