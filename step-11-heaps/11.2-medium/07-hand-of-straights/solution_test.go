package handofstraights

import "testing"

func TestIsNStraightHand(t *testing.T) {
	tests := []struct {
		hand []int
		g    int
		want bool
	}{{[]int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3, true}, {[]int{1, 2, 3, 4, 5}, 4, false}, {[]int{1, 1, 2, 2, 3, 3}, 3, true}, {[]int{1, 2, 4, 5, 6, 7}, 3, false}, {[]int{5}, 1, true}}
	for _, tc := range tests {
		if got := IsNStraightHand(tc.hand, tc.g); got != tc.want {
			t.Errorf("IsNStraightHand(%v,%d) = %v", tc.hand, tc.g, got)
		}
	}
}
