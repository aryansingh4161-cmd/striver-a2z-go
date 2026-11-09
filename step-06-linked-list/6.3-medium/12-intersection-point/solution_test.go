package intersectionpoint

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

// join builds A = a + shared and B = b + shared, sharing the same tail nodes.
func join(a, b, shared []int) (*Node, *Node, *Node) {
	common := list.FromSlice(shared)
	attach := func(x []int) *Node {
		h := list.FromSlice(x)
		if h == nil {
			return common
		}
		cur := h
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = common
		return h
	}
	return attach(a), attach(b), common
}

func TestIntersection(t *testing.T) {
	tests := []struct{ a, b, shared []int }{
		{[]int{4, 1}, []int{5, 6, 1}, []int{8, 4, 5}},
		{[]int{1, 9, 1}, []int{3}, []int{2, 4}},
		{[]int{2, 6, 4}, []int{1, 5}, nil},
		{nil, []int{1}, []int{7}},
	}
	for _, tc := range tests {
		ha, hb, want := join(tc.a, tc.b, tc.shared)
		if got := Intersection(ha, hb); got != want {
			t.Errorf("Intersection %v %v wrong", tc.a, tc.b)
		}
		if got := IntersectionSet(ha, hb); got != want {
			t.Errorf("IntersectionSet %v %v wrong", tc.a, tc.b)
		}
	}
}
