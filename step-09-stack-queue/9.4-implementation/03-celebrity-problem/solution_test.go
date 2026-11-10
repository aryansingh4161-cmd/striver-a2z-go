package celebrityproblem

import "testing"

func TestCelebrity(t *testing.T) {
	tests := []struct {
		m    [][]int
		want int
	}{
		{[][]int{{0, 1, 0}, {0, 0, 0}, {0, 1, 0}}, 1},
		{[][]int{{0, 1}, {1, 0}}, -1},
		{[][]int{{0}}, 0},
		{[][]int{{0, 1, 1, 0}, {0, 0, 0, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}, 1},
		{[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, -1},
	}
	for _, tc := range tests {
		if got := Celebrity(tc.m); got != tc.want {
			t.Errorf("Celebrity(%v) = %d", tc.m, got)
		}
		if got := CelebrityCount(tc.m); got != tc.want {
			t.Errorf("CelebrityCount(%v) = %d", tc.m, got)
		}
	}
}
