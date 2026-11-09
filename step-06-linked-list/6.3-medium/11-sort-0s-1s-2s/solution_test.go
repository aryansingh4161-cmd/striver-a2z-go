package sort012list

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func TestSort(t *testing.T) {
	cases := [][]int{{1, 0, 2, 1, 0}, {2, 2, 0}, {1, 1}, {}, {0}, {2, 1, 0}, {2, 0, 2, 0}}
	for _, a := range cases {
		want := append([]int{}, a...)
		slices.Sort(want)
		if got := list.ToSlice(Sort(list.FromSlice(a))); !slices.Equal(got, want) {
			t.Errorf("Sort(%v) = %v", a, got)
		}
		if got := list.ToSlice(SortCount(list.FromSlice(a))); !slices.Equal(got, want) {
			t.Errorf("SortCount(%v) = %v", a, got)
		}
	}
}
