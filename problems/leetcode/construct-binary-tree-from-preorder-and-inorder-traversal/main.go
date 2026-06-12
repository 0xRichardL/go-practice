package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 462ms Beats 5.05%
// 78.52MB Beats 5.73%
func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}
	subPreorder := make([]int, 0)
	for _, v := range preorder {
		if _, ok := m[v]; ok {
			subPreorder = append(subPreorder, v)
		}
	}
	if len(subPreorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	root.Val = subPreorder[0]
	idx := m[root.Val]
	root.Left = buildTree(subPreorder, inorder[:idx])
	root.Right = buildTree(subPreorder, inorder[idx+1:])
	return root
}
