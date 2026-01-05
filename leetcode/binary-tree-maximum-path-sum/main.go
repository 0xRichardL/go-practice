package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0ms Beats 100.00%
func maxPathSum(root *TreeNode) int {
	// single node case.
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	max := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		left := math.MinInt
		right := math.MinInt

		if node.Left != nil {
			left = dfs(node.Left)
			if max < left {
				max = left
			}
		}
		if node.Right != nil {
			right = dfs(node.Right)
			if max < right {
				max = right
			}
		}
		// Cut parent node
		pSum := node.Val
		if left > 0 {
			pSum += left
		}
		if right > 0 {
			pSum += right
		}
		if pSum > max {
			max = pSum
		}
		// branch case: can pass LN || NR
		sum := node.Val
		if left > 0 || right > 0 {
			if left > right {
				sum += left
			} else {
				sum += right
			}
		}
		return sum
	}
	rootSum := dfs(root)
	if rootSum > max {
		return rootSum
	}
	return max
}
