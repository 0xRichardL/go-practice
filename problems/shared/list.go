package shared

type ListNode struct {
	Val  int
	Next *ListNode
}

func array_to_list(arr []int) *ListNode {
	dummy := &ListNode{}
	node := dummy
	for _, v := range arr {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return dummy.Next
}

func list_to_array(list *ListNode) []int {
	arr := make([]int, 0)
	for node := list; node != nil; node = node.Next {
		arr = append(arr, node.Val)
	}
	return arr
}
