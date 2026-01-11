package algorithms

import "testing"

func Test_binarySearch(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Found Target",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "Target Not Found",
			nums:   []int{10, 20, 30, 40, 50},
			target: 25,
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binarySearch(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
