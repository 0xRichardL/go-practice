package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0ms Beats 100.00%
// 4.21MB Beats 24.24%
func sumNumbers_string_path(root *TreeNode) int {
	sum := 0
	var dfs func(node *TreeNode, path string)
	dfs = func(node *TreeNode, prepPath string) {
		path := prepPath + string(rune(node.Val)+'0')
		if node.Left == nil && node.Right == nil {
			i, _ := strconv.Atoi(path)
			sum += i
		}
		if node.Left != nil {
			dfs(node.Left, path)
		}
		if node.Right != nil {
			dfs(node.Right, path)
		}
	}
	dfs(root, "")
	return sum
}

// 0ms Beats 100.00%
// 4.19 MB Beats 81.21%
func sumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(node *TreeNode, path int)
	dfs = func(node *TreeNode, prepPath int) {
		path := prepPath*10 + node.Val
		if node.Left == nil && node.Right == nil {
			sum += path
		}
		if node.Left != nil {
			dfs(node.Left, path)
		}
		if node.Right != nil {
			dfs(node.Right, path)
		}
	}
	dfs(root, 0)
	return sum
}
