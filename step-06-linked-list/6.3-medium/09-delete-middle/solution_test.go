package deletemiddle

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestDeleteMiddle(t *testing.T) {
	tests := []struct{ a, want []int }{
		{[]int{1, 3, 4, 7, 1, 2, 6}, []int{1, 3, 4, 1, 2, 6}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 4}},
		{[]int{2, 1}, []int{2}},
		{[]int{1}, []int{}},
	}
	for _, tc := range tests {
		if got := list.ToSlice(DeleteMiddle(list.FromSlice(tc.a))); !slices.Equal(got, tc.want) {
			t.Errorf("DeleteMiddle(%v) = %v", tc.a, got)
		}
	}
}
