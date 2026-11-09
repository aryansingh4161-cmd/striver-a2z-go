package maxconsecutiveones

import "testing"

func TestMaxOnes(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{{[]int{1, 1, 0, 1, 1, 1}, 3}, {[]int{1, 0, 1, 1, 0, 1}, 2}, {[]int{0, 0}, 0}, {[]int{1}, 1}}
	for _, tc := range tests {
		if got := MaxOnes(tc.a); got != tc.want {
			t.Errorf("MaxOnes(%v) = %d", tc.a, got)
		}
	}
}
