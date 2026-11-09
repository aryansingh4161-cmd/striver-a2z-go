package pairswithsum

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestPairs(t *testing.T) {
	tests := []struct {
		a      []int
		target int
		want   [][2]int
	}{
		{[]int{1, 2, 4, 5, 6, 8, 9}, 7, [][2]int{{1, 6}, {2, 5}}},
		{[]int{1, 5, 6}, 6, [][2]int{{1, 5}}},
		{[]int{1, 2}, 10, [][2]int{}},
		{[]int{1, 2, 3, 4}, 5, [][2]int{{1, 4}, {2, 3}}},
	}
	for _, tc := range tests {
		if got := Pairs(list.DFromSlice(tc.a), tc.target); !slices.Equal(got, tc.want) {
			t.Errorf("Pairs(%v,%d) = %v", tc.a, tc.target, got)
		}
	}
}
