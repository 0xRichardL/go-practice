package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}
	nodes := make([]*TreeNode, 0)
	for _, v := range preorder {
		nodes = append(nodes, &TreeNode{Val: v})
	}
}
