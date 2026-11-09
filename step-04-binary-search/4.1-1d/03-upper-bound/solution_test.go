package upperbound

import (
	"sort"
	"testing"
)

func TestUpperBound(t *testing.T) {
	a := []int{2, 3, 6, 7, 8, 8, 11, 11, 11, 12}
	for _, tc := range []struct{ x, want int }{{6, 3}, {11, 9}, {12, 10}, {0, 0}, {8, 6}} {
		if got := UpperBound(a, tc.x); got != tc.want {
			t.Errorf("UpperBound(%d) = %d", tc.x, got)
		}
	}
	for x := 0; x <= 13; x++ {
		want := sort.Search(len(a), func(i int) bool { return a[i] > x })
		if got := UpperBound(a, x); got != want {
			t.Errorf("x=%d got %d want %d", x, got, want)
		}
	}
}
