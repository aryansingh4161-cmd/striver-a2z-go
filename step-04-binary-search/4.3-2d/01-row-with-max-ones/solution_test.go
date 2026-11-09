package rowwithmaxones

import "testing"

func TestRowMaxOnes(t *testing.T) {
	tests := []struct {
		m    [][]int
		want int
	}{
		{[][]int{{0, 0, 1, 1, 1}, {0, 0, 0, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 1, 1, 1}}, 2},
		{[][]int{{0, 0}, {0, 0}}, -1},
		{[][]int{{0, 1}, {0, 1}}, 0},
		{[][]int{{1, 1, 1}, {0, 1, 1}}, 0},
	}
	for _, tc := range tests {
		if got := RowMaxOnes(tc.m); got != tc.want {
			t.Errorf("RowMaxOnes(%v) = %d", tc.m, got)
		}
		if got := RowMaxOnesStaircase(tc.m); got != tc.want {
			t.Errorf("RowMaxOnesStaircase(%v) = %d", tc.m, got)
		}
	}
}
