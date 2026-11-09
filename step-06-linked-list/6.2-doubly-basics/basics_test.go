package doublybasics

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func check(t *testing.T, name string, got *DNode, want []int) {
	t.Helper()
	g, ok := list.DToSlice(got)
	if !ok || !slices.Equal(g, want) {
		t.Errorf("%s = %v (prev ok=%v), want %v", name, g, ok, want)
	}
}

func TestDoubly(t *testing.T) {
	l := func() *DNode { return list.DFromSlice([]int{1, 2, 3}) }
	check(t, "InsertHead", InsertHead(l(), 0), []int{0, 1, 2, 3})
	check(t, "InsertTail", InsertTail(l(), 4), []int{1, 2, 3, 4})
	check(t, "InsertBeforeK 2", InsertBeforeK(l(), 9, 2), []int{1, 9, 2, 3})
	check(t, "InsertBeforeK 3", InsertBeforeK(l(), 9, 3), []int{1, 2, 9, 3})
	check(t, "DeleteHead", DeleteHead(l()), []int{2, 3})
	check(t, "DeleteTail", DeleteTail(l()), []int{1, 2})
	check(t, "DeleteK 2", DeleteK(l(), 2), []int{1, 3})
	check(t, "DeleteK 1", DeleteK(l(), 1), []int{2, 3})
	check(t, "DeleteK 3", DeleteK(l(), 3), []int{1, 2})
	check(t, "Reverse", Reverse(l()), []int{3, 2, 1})
	check(t, "Reverse one", Reverse(list.DFromSlice([]int{7})), []int{7})
	if Reverse(nil) != nil {
		t.Error("Reverse(nil)")
	}
}
