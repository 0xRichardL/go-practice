package reversenodesinkgroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	dummy := &ListNode{Next: head}
	grpPrev := dummy
	count := 1

	for node := head; node != nil; count++ {
		if count == k {
			//preserve grpNext node
			grpNext := node.Next
			grpHead := grpPrev.Next

			prev := grpNext
			curr := grpHead

			for curr != grpNext {
				next := curr.Next
				curr.Next = prev
				prev = curr
				curr = next
			}
			grpPrev.Next = prev
			grpPrev = grpHead
			node = grpNext
			count = 0
		} else {
			node = node.Next
		}
	}
	return dummy.Next
}
