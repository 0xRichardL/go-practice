package main

import (
	"fmt"
)

// 4ms | Beats 18.64%
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0, len(nums1)+len(nums2))
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			merged = append(merged, nums1[i])
			i++
		} else {
			merged = append(merged, nums2[j])
			j++
		}
	}
	if i < len(nums1) {
		merged = append(merged, nums1[i:]...)
	}
	if j < len(nums2) {
		merged = append(merged, nums2[j:]...)
	}

	mid := len(merged) / 2
	if len(merged)%2 == 1 {
		return float64(merged[mid])
	}
	return (float64(merged[mid]) + float64(merged[mid-1])) / 2
}

func main() {
	fmt.Println("[1,3] [2] \t\t2.0 =>", findMedianSortedArrays([]int{1, 3}, []int{2}))                             // 2.0
	fmt.Println("[1,2] [3,4] \t\t2.5 =>", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))                        // 2.5
	fmt.Println("[0,0] [0,0] \t\t0.0 =>", findMedianSortedArrays([]int{0, 0}, []int{0, 0}))                        // 0.0
	fmt.Println("[] [1] \t\t\t1.0 =>", findMedianSortedArrays([]int{}, []int{1}))                                  // 1.0
	fmt.Println("[2] [] \t\t\t2.0 =>", findMedianSortedArrays([]int{2}, []int{}))                                  // 2.0
	fmt.Println("[1] [2] \t\t1.5 =>", findMedianSortedArrays([]int{1}, []int{2}))                                  // 1.5
	fmt.Println("[2,2,4,4] [2,2,2,4,4] \t2.0 =>", findMedianSortedArrays([]int{2, 2, 4, 4}, []int{2, 2, 2, 4, 4})) // 3.0
}
