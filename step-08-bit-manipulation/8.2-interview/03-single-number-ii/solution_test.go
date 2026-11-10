package singlenumberii

import "testing"

func TestSingle(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{{[]int{2, 2, 3, 2}, 3}, {[]int{0, 1, 0, 1, 0, 1, 99}, 99}, {[]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}, -4}, {[]int{5}, 5}}
	for _, tc := range tests {
		if got := Single(tc.a); got != tc.want {
			t.Errorf("Single(%v) = %d", tc.a, got)
		}
		if got := SingleBitCount(tc.a); got != tc.want {
			t.Errorf("SingleBitCount(%v) = %d", tc.a, got)
		}
	}
}
