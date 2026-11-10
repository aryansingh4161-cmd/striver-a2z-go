package jumpgameii

import "testing"

func TestMinJumps(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{{[]int{2, 3, 1, 1, 4}, 2}, {[]int{2, 3, 0, 1, 4}, 2}, {[]int{0}, 0}, {[]int{1, 1, 1, 1}, 3}, {[]int{10, 1, 1}, 1}}
	for _, tc := range tests {
		if got := MinJumps(tc.a); got != tc.want {
			t.Errorf("MinJumps(%v) = %d", tc.a, got)
		}
	}
}
