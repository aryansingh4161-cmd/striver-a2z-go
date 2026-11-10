package jumpgame

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		a    []int
		want bool
	}{{[]int{2, 3, 1, 1, 4}, true}, {[]int{3, 2, 1, 0, 4}, false}, {[]int{0}, true}, {[]int{0, 1}, false}, {[]int{1, 0, 1}, false}}
	for _, tc := range tests {
		if got := CanJump(tc.a); got != tc.want {
			t.Errorf("CanJump(%v) = %v", tc.a, got)
		}
	}
}
