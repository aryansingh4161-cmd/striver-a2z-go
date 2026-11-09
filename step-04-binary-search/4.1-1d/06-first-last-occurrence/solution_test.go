package firstlastoccurrence

import "testing"

func TestSearchRange(t *testing.T) {
	a := []int{5, 7, 7, 8, 8, 10}
	tests := []struct {
		x    int
		want [2]int
	}{{8, [2]int{3, 4}}, {6, [2]int{-1, -1}}, {5, [2]int{0, 0}}, {10, [2]int{5, 5}}}
	for _, tc := range tests {
		if got := SearchRange(a, tc.x); got != tc.want {
			t.Errorf("SearchRange(%d) = %v", tc.x, got)
		}
	}
	if got := SearchRange(nil, 1); got != [2]int{-1, -1} {
		t.Errorf("empty = %v", got)
	}
}
