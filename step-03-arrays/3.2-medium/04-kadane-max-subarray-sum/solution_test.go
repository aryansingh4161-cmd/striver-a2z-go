package kadane

import "testing"

func TestMaxSubarray(t *testing.T) {
	tests := []struct {
		a                []int
		best, start, end int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6, 3, 6},
		{[]int{1}, 1, 0, 0},
		{[]int{5, 4, -1, 7, 8}, 23, 0, 4},
		{[]int{-3, -1, -2}, -1, 1, 1},
	}
	for _, tc := range tests {
		b, s, e := MaxSubarray(tc.a)
		if b != tc.best || s != tc.start || e != tc.end {
			t.Errorf("MaxSubarray(%v) = %d [%d,%d], want %d [%d,%d]", tc.a, b, s, e, tc.best, tc.start, tc.end)
		}
		if got := MaxSubarrayBrute(tc.a); got != tc.best {
			t.Errorf("MaxSubarrayBrute(%v) = %d", tc.a, got)
		}
	}
}
