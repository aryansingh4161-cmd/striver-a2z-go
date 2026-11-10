package nonoverlappingintervals

import "testing"

func TestEraseOverlap(t *testing.T) {
	tests := []struct {
		iv   [][]int
		want int
	}{{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1}, {[][]int{{1, 2}, {1, 2}, {1, 2}}, 2}, {[][]int{{1, 2}, {2, 3}}, 0}, {[][]int{}, 0}}
	for _, tc := range tests {
		if got := EraseOverlap(tc.iv); got != tc.want {
			t.Errorf("EraseOverlap(%v) = %d", tc.iv, got)
		}
	}
}
