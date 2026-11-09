package searchinsertposition

import "testing"

func TestSearchInsert(t *testing.T) {
	a := []int{1, 3, 5, 6}
	for _, tc := range []struct{ x, want int }{{5, 2}, {2, 1}, {7, 4}, {0, 0}} {
		if got := SearchInsert(a, tc.x); got != tc.want {
			t.Errorf("SearchInsert(%d) = %d", tc.x, got)
		}
	}
}
