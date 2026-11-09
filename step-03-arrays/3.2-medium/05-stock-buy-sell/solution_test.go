package stockbuysell

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		p    []int
		want int
	}{{[]int{7, 1, 5, 3, 6, 4}, 5}, {[]int{7, 6, 4, 3, 1}, 0}, {[]int{2, 4, 1}, 2}, {[]int{}, 0}}
	for _, tc := range tests {
		if got := MaxProfit(tc.p); got != tc.want {
			t.Errorf("MaxProfit(%v) = %d", tc.p, got)
		}
	}
}
