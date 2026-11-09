package search2dmatrixii

import "testing"

func TestSearchMatrix(t *testing.T) {
	m := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	for _, tc := range []struct {
		x    int
		want bool
	}{{5, true}, {20, false}, {30, true}, {18, true}, {0, false}} {
		if got := SearchMatrix(m, tc.x); got != tc.want {
			t.Errorf("SearchMatrix(%d) = %v", tc.x, got)
		}
	}
}
