package singlybasics

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

func check(t *testing.T, name string, got *Node, want []int) {
	t.Helper()
	if g := list.ToSlice(got); !slices.Equal(g, want) {
		t.Errorf("%s = %v, want %v", name, g, want)
	}
}

func TestBasics(t *testing.T) {
	l := func() *Node { return list.FromSlice([]int{1, 2, 3}) }
	if Length(l()) != 3 || Length(nil) != 0 {
		t.Error("Length")
	}
	if !Search(l(), 2) || Search(l(), 9) {
		t.Error("Search")
	}
	check(t, "InsertHead", InsertHead(l(), 0), []int{0, 1, 2, 3})
	check(t, "InsertTail", InsertTail(l(), 4), []int{1, 2, 3, 4})
	check(t, "InsertTail empty", InsertTail(nil, 4), []int{4})
	check(t, "InsertAt 2", InsertAt(l(), 9, 2), []int{1, 9, 2, 3})
	check(t, "InsertAt end", InsertAt(l(), 9, 4), []int{1, 2, 3, 9})
	check(t, "InsertAt too far", InsertAt(l(), 9, 7), []int{1, 2, 3})
	check(t, "InsertBefore 3", InsertBefore(l(), 9, 3), []int{1, 2, 9, 3})
	check(t, "InsertBefore head", InsertBefore(l(), 9, 1), []int{9, 1, 2, 3})
	check(t, "DeleteHead", DeleteHead(l()), []int{2, 3})
	check(t, "DeleteTail", DeleteTail(l()), []int{1, 2})
	check(t, "DeleteTail one", DeleteTail(list.FromSlice([]int{5})), []int{})
	check(t, "DeleteAt 2", DeleteAt(l(), 2), []int{1, 3})
	check(t, "DeleteAt 1", DeleteAt(l(), 1), []int{2, 3})
	check(t, "DeleteAt 3", DeleteAt(l(), 3), []int{1, 2})
	check(t, "DeleteValue 3", DeleteValue(l(), 3), []int{1, 2})
	check(t, "DeleteValue missing", DeleteValue(l(), 8), []int{1, 2, 3})
}
