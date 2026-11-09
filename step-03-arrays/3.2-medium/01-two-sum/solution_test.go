package twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		a      []int
		target int
		want   []int
	}{
		{[]int{2, 6, 5, 8, 11}, 14, []int{1, 3}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{1, 2}, 10, nil},
	}
	for _, tc := range tests {
		if got := TwoSum(tc.a, tc.target); !slices.Equal(got, tc.want) {
			t.Errorf("TwoSum(%v,%d) = %v", tc.a, tc.target, got)
		}
		if got := HasPair(tc.a, tc.target); got != (tc.want != nil) {
			t.Errorf("HasPair(%v,%d) = %v", tc.a, tc.target, got)
		}
	}
}
