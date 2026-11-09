package leadersinarray

import (
	"slices"
	"testing"
)

func TestLeaders(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{10, 22, 12, 3, 0, 6}, []int{22, 12, 6}},
		{[]int{1, 2, 3}, []int{3}},
		{[]int{3, 2, 1}, []int{3, 2, 1}},
		{[]int{5, 5}, []int{5}},
	}
	for _, tc := range tests {
		if got := Leaders(tc.a); !slices.Equal(got, tc.want) {
			t.Errorf("Leaders(%v) = %v", tc.a, got)
		}
	}
}
