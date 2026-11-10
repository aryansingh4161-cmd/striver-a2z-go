package twosinglenumbers

import "testing"

func TestTwoSingles(t *testing.T) {
	tests := []struct {
		a    []int
		x, y int
	}{{[]int{1, 2, 1, 3, 2, 5}, 3, 5}, {[]int{-1, 0}, -1, 0}, {[]int{0, 1}, 0, 1}, {[]int{4, 7, 4, 9, 9, -2}, -2, 7}}
	for _, tc := range tests {
		x, y := TwoSingles(tc.a)
		if x != tc.x || y != tc.y {
			t.Errorf("TwoSingles(%v) = %d,%d", tc.a, x, y)
		}
	}
}
