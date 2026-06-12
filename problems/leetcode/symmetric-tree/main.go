package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0ms Beats 100%
// 4.79MB Beats 78.91%
// Double tree traveling solution.
func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return check(root.Left, root.Right)
}
