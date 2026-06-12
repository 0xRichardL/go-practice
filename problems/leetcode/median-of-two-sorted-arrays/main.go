// findMedianSortedArrays finds the median of two sorted arrays by merging them partially.
// It uses an optimized approach that only merges elements up to the median position,
// avoiding the need to merge the entire arrays.
//
// Algorithm:
// 1. Calculate the total length and determine if it's odd or even
// 2. Merge arrays up to the median position (expLen + 1 elements)
// 3. For odd length: return the middle element
// 4. For even length: return the average of two middle elements
//
// Time Complexity: O((m+n)/2) where m and n are lengths of input arrays
// Space Complexity: O((m+n)/2) for the partial merged array
//
// Note: This solution achieves 0ms runtime by avoiding full array merging,
// only processing elements needed to find the median.
//
// Leetcode: 0ms Beats 100.00%
package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	totalLen := len(nums1) + len(nums2)
	expLen := totalLen / 2
	odd := totalLen%2 == 1
	if odd {
		expLen++
	}

	for len(merged) < expLen+1 && (i < len(nums1) || j < len(nums2)) {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				merged = append(merged, nums1[i])
				i++
			} else {
				merged = append(merged, nums2[j])
				j++
			}
		} else if i < len(nums1) {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	if odd {
		return float64(merged[expLen-1])
	}
	return (float64(merged[expLen]) + float64(merged[expLen-1])) / 2
}

func main() {
	testCases := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.0},
		{[]int{}, []int{1}, 1.0},
		{[]int{2}, []int{}, 2.0},
		{[]int{1}, []int{2}, 1.5},
		{[]int{2, 2, 4, 4}, []int{2, 2, 4, 4}, 3.0},
		{[]int{2, 2, 4, 4}, []int{2, 2, 2, 4, 4}, 2.0},
		{[]int{0, 0, 0, 0, 0}, []int{-1, 0, 0, 0, 0, 0, 1}, 0.0},
	}

	for _, tc := range testCases {
		fmt.Println(tc.nums1, tc.nums2)
		fmt.Println(tc.expected, " -> ", findMedianSortedArrays(tc.nums1, tc.nums2))
		fmt.Println("- - - - - - - -")
	}
}
