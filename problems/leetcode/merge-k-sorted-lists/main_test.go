package mergeksortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func Test_mergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists []*ListNode
		want  *ListNode
	}{
		{
			name: "Case 1",
			lists: []*ListNode{
				array_to_list([]int{1, 4, 5}),
				array_to_list([]int{1, 3, 4}),
				array_to_list([]int{2, 6}),
			},
			want: array_to_list([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name:  "Case 2",
			lists: []*ListNode{},
			want:  array_to_list([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeKLists(tt.lists)
			assert.EqualValues(t, list_to_array(tt.want), list_to_array(got))
		})
	}
}
