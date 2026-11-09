package singlenumber

import "testing"

func TestSingle(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{{[]int{4, 1, 2, 1, 2}, 4}, {[]int{2, 2, 1}, 1}, {[]int{1}, 1}, {[]int{-3, 5, 5}, -3}}
	for _, tc := range tests {
		if got := SingleXOR(tc.a); got != tc.want {
			t.Errorf("SingleXOR(%v) = %d", tc.a, got)
		}
		if got := SingleMap(tc.a); got != tc.want {
			t.Errorf("SingleMap(%v) = %d", tc.a, got)
		}
	}
}
