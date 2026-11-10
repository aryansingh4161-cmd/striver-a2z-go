package connectropes

import "testing"

func TestMinCost(t *testing.T) {
	tests := []struct {
		r    []int
		want int
	}{{[]int{4, 3, 2, 6}, 29}, {[]int{4, 2, 7, 6, 9}, 62}, {[]int{10}, 0}, {[]int{}, 0}}
	for _, tc := range tests {
		if got := MinCost(tc.r); got != tc.want {
			t.Errorf("MinCost(%v) = %d", tc.r, got)
		}
	}
}
