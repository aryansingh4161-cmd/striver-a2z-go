package singleelementsortedarray

import "testing"

func TestSingle(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
		{[]int{1}, 1},
		{[]int{1, 2, 2}, 1},
		{[]int{1, 1, 2}, 2},
	}
	for _, tc := range tests {
		if got := Single(tc.a); got != tc.want {
			t.Errorf("Single(%v) = %d", tc.a, got)
		}
		if got := SingleXOR(tc.a); got != tc.want {
			t.Errorf("SingleXOR(%v) = %d", tc.a, got)
		}
	}
}
