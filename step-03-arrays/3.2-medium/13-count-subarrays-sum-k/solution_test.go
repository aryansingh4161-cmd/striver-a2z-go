package countsubarrayssumk

import "testing"

func TestCount(t *testing.T) {
	tests := []struct {
		a       []int
		k, want int
	}{
		{[]int{3, 1, 2, 4}, 6, 2},
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1, -1, 0}, 0, 3},
		{[]int{}, 0, 0},
	}
	for _, tc := range tests {
		if got := Count(tc.a, tc.k); got != tc.want {
			t.Errorf("Count(%v,%d) = %d", tc.a, tc.k, got)
		}
		if got := CountBrute(tc.a, tc.k); got != tc.want {
			t.Errorf("CountBrute(%v,%d) = %d", tc.a, tc.k, got)
		}
	}
}
