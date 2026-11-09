package mindaysbouquets

import "testing"

func TestMinDays(t *testing.T) {
	tests := []struct {
		bloom      []int
		m, k, want int
	}{
		{[]int{7, 7, 7, 7, 13, 11, 12, 7}, 2, 3, 12},
		{[]int{1, 10, 3, 10, 2}, 3, 1, 3},
		{[]int{1, 10, 3, 10, 2}, 3, 2, -1},
		{[]int{1000000000, 1000000000}, 1, 1, 1000000000},
	}
	for _, tc := range tests {
		if got := MinDays(tc.bloom, tc.m, tc.k); got != tc.want {
			t.Errorf("MinDays(%v,%d,%d) = %d", tc.bloom, tc.m, tc.k, got)
		}
	}
}
