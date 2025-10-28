package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{
		Val:  0,
		Next: nil,
	}
	l := r
	car := 0
	i := 1
	j := 1
	for true {
		s := l1.Val*i + l2.Val*j + car
		car = s / 10
		l.Val = s % 10
		if l1.Next == nil && l2.Next == nil && s < 10 {
			break
		}
		l.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		l = l.Next
		if l1.Next == nil {
			i = 0
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			j = 0
		} else {
			l2 = l2.Next
		}
	}
	return r
}
