package main

import "fmt"

func main() {
	nums := []string{"1", "2", "3", "4", "5"}
	// Traditional for loop
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	// Range-based for loop
	for i, v := range nums {
		fmt.Println(i, v)
	}
	// While-like loop
	i := 0
	for i < len(nums) {
		fmt.Println(nums[i])
		i++
	}
}
