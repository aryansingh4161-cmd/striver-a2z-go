package rottingoranges

import "testing"

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		g    [][]int
		want int
	}{{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4}, {[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1}, {[][]int{{0, 2}}, 0}, {[][]int{{0}}, 0}, {[][]int{{2, 1, 1, 1, 2}}, 2}}
	for _, tc := range tests {
		if got := OrangesRotting(tc.g); got != tc.want {
			t.Errorf("OrangesRotting(%v) = %d", tc.g, got)
		}
	}
}
