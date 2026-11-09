package mininrotatedsortedarray

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		a    []int
		want int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{11, 13, 15, 17}, 11},
		{[]int{2, 1}, 1},
		{[]int{5}, 5},
	}
	for _, tc := range tests {
		if got := FindMin(tc.a); got != tc.want {
			t.Errorf("FindMin(%v) = %d", tc.a, got)
		}
	}
}
