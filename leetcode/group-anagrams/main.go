package main

import (
	"sort"
	"strings"
)

// 6ms Beats 88.79%
func groupAnagrams_1(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		chars := []byte(strings.Clone(str))
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		m[string(chars)] = append(m[string(chars)], str)
	}
	res := make([][]string, 0)
	for _, gr := range m {
		res = append(res, gr)
	}
	return res
}

// 5ms Beats 92.89%
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int8][]string)
	for _, str := range strs {
		var chars [26]int8
		for _, r := range str {
			chars[r-'a']++
		}
		m[chars] = append(m[chars], str)
	}
	res := make([][]string, 0)
	for _, gr := range m {
		res = append(res, gr)
	}
	return res
}
