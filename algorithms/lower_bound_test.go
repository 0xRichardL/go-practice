package algorithms

import "testing"

func Test_lowerBound(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Lower Bound Found",
			nums:   []int{1, 3, 5, 7},
			target: 5,
			want:   2,
		},
		{
			name:   "Lower Bound Between Elements",
			nums:   []int{1, 3, 5, 7},
			target: 6,
			want:   3,
		},
		{
			name:   "Lower Bound Before All Elements",
			nums:   []int{1, 3, 5, 7},
			target: 0,
			want:   0,
		},
		{
			name:   "Lower Bound After All Elements",
			nums:   []int{1, 3, 5, 7},
			target: 8,
			want:   4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lowerBound(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("lowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
