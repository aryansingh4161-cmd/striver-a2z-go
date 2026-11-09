package lowerbound

import (
	"sort"
	"testing"
)

func TestLowerBound(t *testing.T) {
	a := []int{3, 5, 8, 15, 19}
	for _, tc := range []struct{ x, want int }{{9, 3}, {8, 2}, {20, 5}, {1, 0}, {3, 0}} {
		if got := LowerBound(a, tc.x); got != tc.want {
			t.Errorf("LowerBound(%d) = %d", tc.x, got)
		}
	}
	b := []int{1, 2, 2, 2, 3, 7, 7}
	for x := 0; x <= 8; x++ {
		if got, want := LowerBound(b, x), sort.SearchInts(b, x); got != want {
			t.Errorf("x=%d got %d want %d", x, got, want)
		}
	}
}
