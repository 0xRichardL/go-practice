package main

import (
	"fmt"
	"sort"
)

type SortedSet struct {
	M   map[int]bool
	Arr []int
}

func NewSortedSet(cap int) *SortedSet {
	return &SortedSet{
		M:   make(map[int]bool),
		Arr: make([]int, 0),
	}
}

func (s *SortedSet) Add(value int) {
	if s.M[value] {
		return
	}
	s.M[value] = true
	l, r := 0, len(s.Arr)
	for l < r {
		mid := (l + r) / 2
		if s.Arr[mid] < value {
			l = mid + 1
		} else {
			r = mid
		}
	}
	// insert at position l
	s.Arr = append(s.Arr, 0)
	copy(s.Arr[l+1:], s.Arr[l:])
	s.Arr[l] = value
}

// 1080ms Beats 5.01%
func longestConsecutive_sorted_set(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sortedSet := NewSortedSet(len(nums))
	for _, n := range nums {
		sortedSet.Add(n)
	}
	dp := sortedSet.Arr
	ln := 1
	max := 0
	for i := 1; i < len(dp); i++ {
		if dp[i] == dp[i-1] {
			continue
		}
		if dp[i] == dp[i-1]+1 {
			ln++
		} else {
			if max < ln {
				max = ln
			}
			ln = 1
		}
	}
	if max < ln {
		return ln
	}
	return max
}

// 55ms Beats 40.96%
func longestConsecutive_map_sort(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}
	dp := make([]int, 0, len(nums))
	for n := range m {
		dp = append(dp, n)
	}
	sort.Ints(dp)
	ln := 1
	max := 0
	for i := 1; i < len(dp); i++ {
		if dp[i] == dp[i-1] {
			continue
		}
		if dp[i] == dp[i-1]+1 {
			ln++
		} else {
			if max < ln {
				max = ln
			}
			ln = 1
		}
	}
	if max < ln {
		return ln
	}
	return max
}

// 6ms Beats 97.31%
func longestConsecutive_sort_nums(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	ln := 1
	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if nums[i] == nums[i-1]+1 {
			ln++
		} else {
			if max < ln {
				max = ln
			}
			ln = 1
		}
	}
	if max < ln {
		return ln
	}
	return max
}

func main() {
	type Test struct {
		nums     []int
		expected int
	}
	tests := []Test{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 2, 6, 7, 8}, 3},
	}
	for _, test := range tests {
		fmt.Println(test)
		res := longestConsecutive(test.nums)
		fmt.Println(test.expected, " - - - > ", res, res == test.expected)
		fmt.Println("- - - - -")
	}
}
