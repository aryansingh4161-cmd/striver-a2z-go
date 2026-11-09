package reversekgroup

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		a    []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1, 2, 3, 4}, 4, []int{4, 3, 2, 1}},
		{[]int{1, 2}, 1, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5, 6}, 3, []int{3, 2, 1, 6, 5, 4}},
	}
	for _, tc := range tests {
		if got := list.ToSlice(ReverseKGroup(list.FromSlice(tc.a), tc.k)); !slices.Equal(got, tc.want) {
			t.Errorf("ReverseKGroup(%v,%d) = %v", tc.a, tc.k, got)
		}
	}
}
