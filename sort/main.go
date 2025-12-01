package main

import (
	"fmt"
	"sort"
	"time"
)

func sortSlice() {
	type Item struct {
		Date time.Time
	}
	arr := []*Item{
		{Date: time.Date(2022, time.April, 23, 0, 0, 0, 0, time.UTC)},
		{Date: time.Date(2022, time.April, 21, 0, 0, 0, 0, time.UTC)},
		{Date: time.Date(2022, time.April, 20, 0, 0, 0, 0, time.UTC)},
		{Date: time.Date(2022, time.April, 22, 0, 0, 0, 0, time.UTC)},
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Date.After(arr[j].Date)
	})
	for _, v := range arr {
		fmt.Println(v.Date.GoString())
	}
}

// Shortcut for sort.Slice for []int.
func sortInts() {
	nums := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(nums)
	fmt.Println(nums)
}

// Binary search for []int. Only works on ascending sorted slices.
func searchInts() {
	nums := []int{1, 2, 3, 4, 5, 6}
	target := 4
	found := sort.SearchInts(nums, target)
	fmt.Println("Found at index:", found)
	notfound := sort.SearchInts(nums, 10)
	fmt.Println("Not found, returned len(nums):", notfound)
}

func main() {
	sortSlice()
	sortInts()
	searchInts()
}
