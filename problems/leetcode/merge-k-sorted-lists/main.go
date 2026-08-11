package mergeksortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func replica(v, c int) (*ListNode, *ListNode) {
	node := &ListNode{Val: v}
	first := node
	for i := 0; i < c-1; i++ {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return first, node
}

func getIndex(v int) int {
	return v + 100000
}

func getValue(i int) int {
	return i - 100000
}

func mergeKLists(lists []*ListNode) *ListNode {
	marker := make([]int, 2*100000)
	for _, list := range lists {
		for node := list; node != nil; node = node.Next {
			marker[getIndex(node.Val)]++
		}
	}
	dummy := &ListNode{}
	node := dummy
	for i, c := range marker {
		if c > 0 {
			f, l := replica(getValue(i), c)
			node.Next = f
			node = l
		}
	}
	return dummy.Next
}
