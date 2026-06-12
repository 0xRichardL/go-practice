package main

import "fmt"

func dpf(height []int) int {
	n := len(height)
	type Level struct {
		L int
		R int
	}
	m := make(map[int]Level)
	for _, h := range height {
		_, ok := m[h]
		if !ok {
			l := 0
			for ; l < n; l++ {
				if height[l] >= h {
					break
				}
			}
			for r := n - 1; r > l; r-- {
				if height[r] >= h {
					m[h] = Level{
						L: l,
						R: r,
					}
					break
				}
			}
		}
	}
	c := 0
	for h, lvl := range m {
		v := h * (lvl.R - lvl.L)
		// fmt.Println(lvl, v)
		if v > c {
			c = v
		}
	}
	return c
}

func twoPtrs(height []int) int {
	i, j := 0, len(height)-1
	c := 0
	for i < j {
		min := height[i]
		if min > height[j] {
			min = height[j]
		}
		v := min * (j - i)
		if v > c {
			c = v
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return c
}

func maxArea(height []int) int {
	return twoPtrs(height)
}

func main() {
	type Test struct {
		Input    []int
		Expected int
	}
	tests := []Test{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{1, 2, 1}, 2},
		{[]int{1, 2, 4, 3}, 4},
		{[]int{1, 2, 1, 99, 1, 99, 1}, 198},
	}
	for _, test := range tests {
		output := maxArea(test.Input)
		fmt.Println(test.Input)
		fmt.Println(test.Expected, ":", output, " ---> ", output == test.Expected)
	}
}
