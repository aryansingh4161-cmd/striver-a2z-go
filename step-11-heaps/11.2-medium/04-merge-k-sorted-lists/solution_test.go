package mergeksortedlists

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestMergeKLists(t *testing.T) {
	lists := []*Node{list.FromSlice([]int{1, 4, 5}), list.FromSlice([]int{1, 3, 4}), list.FromSlice([]int{2, 6})}
	if got := list.ToSlice(MergeKLists(lists)); !slices.Equal(got, []int{1, 1, 2, 3, 4, 4, 5, 6}) {
		t.Errorf("got %v", got)
	}
	if MergeKLists(nil) != nil || MergeKLists([]*Node{nil}) != nil {
		t.Error("empty")
	}
}
