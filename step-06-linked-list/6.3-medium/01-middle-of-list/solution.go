package middleoflist

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// MiddleCount counts first, then walks n/2 steps. Two passes.
func MiddleCount(head *Node) *Node {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	cur := head
	for i := 0; i < n/2; i++ {
		cur = cur.Next
	}
	return cur
}

// Middle uses slow/fast pointers. Returns the second middle for even length.
// Time O(n), space O(1).
func Middle(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
