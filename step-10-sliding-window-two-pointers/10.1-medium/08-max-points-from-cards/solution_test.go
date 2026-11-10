package maxpointsfromcards

import "testing"

func TestMaxScore(t *testing.T) {
	tests := []struct {
		c       []int
		k, want int
	}{{[]int{1, 2, 3, 4, 5, 6, 1}, 3, 12}, {[]int{2, 2, 2}, 2, 4}, {[]int{9, 7, 7, 9, 7, 7, 9}, 7, 55}, {[]int{1, 1000, 1}, 1, 1}, {[]int{6, 2, 3, 4, 7, 2, 1, 7, 1}, 4, 16}}
	for _, tc := range tests {
		if got := MaxScore(tc.c, tc.k); got != tc.want {
			t.Errorf("MaxScore(%v,%d) = %d", tc.c, tc.k, got)
		}
	}
}
