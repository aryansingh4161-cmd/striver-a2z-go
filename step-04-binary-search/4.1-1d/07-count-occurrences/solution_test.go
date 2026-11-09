package countoccurrences

import "testing"

func TestCount(t *testing.T) {
	a := []int{2, 2, 3, 3, 3, 3, 4}
	for _, tc := range []struct{ x, want int }{{3, 4}, {2, 2}, {4, 1}, {5, 0}, {1, 0}} {
		if got := Count(a, tc.x); got != tc.want {
			t.Errorf("Count(%d) = %d", tc.x, got)
		}
	}
}
