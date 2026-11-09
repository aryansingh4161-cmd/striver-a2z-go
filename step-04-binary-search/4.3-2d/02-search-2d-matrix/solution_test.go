package search2dmatrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	m := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	for _, tc := range []struct {
		x    int
		want bool
	}{{3, true}, {13, false}, {60, true}, {1, true}, {0, false}, {61, false}} {
		if got := SearchMatrix(m, tc.x); got != tc.want {
			t.Errorf("SearchMatrix(%d) = %v", tc.x, got)
		}
	}
}
