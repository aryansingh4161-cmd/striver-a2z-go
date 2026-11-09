package flattenlist

import (
	"slices"
	"testing"
)

func build(cols [][]int) *Node {
	var head, prevTop *Node
	for _, col := range cols {
		var top, last *Node
		for _, v := range col {
			n := &Node{Val: v}
			if top == nil {
				top = n
			} else {
				last.Child = n
			}
			last = n
		}
		if head == nil {
			head = top
		} else {
			prevTop.Next = top
		}
		prevTop = top
	}
	return head
}

func vals(n *Node) []int {
	out := []int{}
	for ; n != nil; n = n.Child {
		if n.Next != nil {
			return nil // Next must be cleared
		}
		out = append(out, n.Val)
	}
	return out
}

func TestFlatten(t *testing.T) {
	cols := [][]int{{5, 7, 8, 30}, {10, 20}, {19, 22, 50}, {28, 35, 40, 45}}
	want := []int{5, 7, 8, 10, 19, 20, 22, 28, 30, 35, 40, 45, 50}
	if got := vals(Flatten(build(cols))); !slices.Equal(got, want) {
		t.Errorf("Flatten = %v", got)
	}
	if got := vals(FlattenSort(build(cols))); !slices.Equal(got, want) {
		t.Errorf("FlattenSort = %v", got)
	}
	if Flatten(nil) != nil {
		t.Error("nil")
	}
}
