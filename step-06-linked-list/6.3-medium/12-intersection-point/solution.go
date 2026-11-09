package intersectionpoint

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// IntersectionSet stores A's nodes in a set. Space O(n).
func IntersectionSet(a, b *Node) *Node {
	seen := map[*Node]bool{}
	for cur := a; cur != nil; cur = cur.Next {
		seen[cur] = true
	}
	for cur := b; cur != nil; cur = cur.Next {
		if seen[cur] {
			return cur
		}
	}
	return nil
}

// Intersection switches heads at the ends. Time O(n+m), space O(1).
func Intersection(headA, headB *Node) *Node {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}
