package countsubsequencessumk

import "testing"

func TestCount(t *testing.T) {
	tests := []struct {
		a       []int
		k, want int
	}{{[]int{1, 2, 1}, 2, 2}, {[]int{1, 2, 3}, 3, 2}, {[]int{5}, 1, 0}, {[]int{}, 0, 1}, {[]int{0, 0}, 0, 4}}
	for _, tc := range tests {
		if got := Count(tc.a, tc.k); got != tc.want {
			t.Errorf("Count(%v,%d) = %d", tc.a, tc.k, got)
		}
	}
}
