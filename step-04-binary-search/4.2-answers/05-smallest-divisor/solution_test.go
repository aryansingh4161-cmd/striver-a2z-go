package smallestdivisor

import "testing"

func TestSmallestDivisor(t *testing.T) {
	tests := []struct {
		a        []int
		th, want int
	}{
		{[]int{1, 2, 5, 9}, 6, 5},
		{[]int{44, 22, 33, 11, 1}, 5, 44},
		{[]int{1, 2, 3, 4, 5}, 8, 3},
	}
	for _, tc := range tests {
		if got := SmallestDivisor(tc.a, tc.th); got != tc.want {
			t.Errorf("SmallestDivisor(%v,%d) = %d", tc.a, tc.th, got)
		}
	}
}
