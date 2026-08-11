package sumclosest

import "testing"

func Test_threeSumClosest(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Case 1",
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		{
			name:   "Case 2",
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0,
		},
		{
			name:   "Case 3",
			nums:   []int{4, 0, 5, -5, 3, 3, 0, -4, -5},
			target: -2,
			want:   -2,
		},
		{
			name:   "Case 4",
			nums:   []int{2, 3, 8, 9, 10},
			target: 16,
			want:   15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumClosest(tt.nums, tt.target)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
