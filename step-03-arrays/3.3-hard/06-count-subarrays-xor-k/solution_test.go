package countsubarraysxork

import "testing"

func TestCountXOR(t *testing.T) {
	tests := []struct {
		a       []int
		k, want int
	}{
		{[]int{4, 2, 2, 6, 4}, 6, 4},
		{[]int{5, 6, 7, 8, 9}, 5, 2},
		{[]int{1, 1}, 0, 1},
	}
	for _, tc := range tests {
		if got := CountXOR(tc.a, tc.k); got != tc.want {
			t.Errorf("CountXOR(%v,%d) = %d", tc.a, tc.k, got)
		}
	}
}
