package reversenodesinkgroup

import (
	"slices"
	"testing"
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

func Test_reverseKGroup(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			name: "Case 1",
			head: array_to_list([]int{1, 2, 3, 4, 5}),
			k:    2,
			want: array_to_list([]int{2, 1, 4, 3, 5}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.head, tt.k)
			if !slices.Equal(list_to_array(got), list_to_array(tt.want)) {
				t.Errorf("reverseKGroup() = %v, want %v", list_to_array(got), list_to_array(tt.want))
			}
		})
	}
}
