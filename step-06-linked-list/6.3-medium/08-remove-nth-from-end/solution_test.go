package removenthfromend

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestRemoveNth(t *testing.T) {
	tests := []struct {
		a    []int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{[]int{1}, 1, []int{}},
		{[]int{1, 2}, 1, []int{1}},
		{[]int{1, 2}, 2, []int{2}},
	}
	for _, tc := range tests {
		if got := list.ToSlice(RemoveNth(list.FromSlice(tc.a), tc.n)); !slices.Equal(got, tc.want) {
			t.Errorf("RemoveNth(%v,%d) = %v", tc.a, tc.n, got)
		}
		if got := list.ToSlice(RemoveNthCount(list.FromSlice(tc.a), tc.n)); !slices.Equal(got, tc.want) {
			t.Errorf("RemoveNthCount(%v,%d) = %v", tc.a, tc.n, got)
		}
	}
}
