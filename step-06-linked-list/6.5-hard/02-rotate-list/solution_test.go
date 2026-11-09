package rotatelist

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		a    []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{0, 1, 2}, 4, []int{2, 0, 1}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{}, 3, []int{}},
	}
	for _, tc := range tests {
		if got := list.ToSlice(RotateRight(list.FromSlice(tc.a), tc.k)); !slices.Equal(got, tc.want) {
			t.Errorf("RotateRight(%v,%d) = %v", tc.a, tc.k, got)
		}
	}
}
