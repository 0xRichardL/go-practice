package main

import "sort"

// 1ms Beats 77.90%
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
			continue
		}
		last := res[len(res)-1]
		if interval[0] > last[1] {
			res = append(res, interval)
		}
		if interval[0] <= last[1] && interval[1] > last[1] {
			last[1] = interval[1]
		}
	}
	return res
}
