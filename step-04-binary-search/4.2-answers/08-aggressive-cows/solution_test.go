package aggressivecows

import "testing"

func TestMaxMinDistance(t *testing.T) {
	tests := []struct {
		stalls  []int
		k, want int
	}{
		{[]int{0, 3, 4, 7, 10, 9}, 4, 3},
		{[]int{4, 2, 1, 3, 6}, 2, 5},
		{[]int{1, 2, 8, 4, 9}, 3, 3},
	}
	for _, tc := range tests {
		if got := MaxMinDistance(tc.stalls, tc.k); got != tc.want {
			t.Errorf("MaxMinDistance(%v,%d) = %d", tc.stalls, tc.k, got)
		}
	}
}
