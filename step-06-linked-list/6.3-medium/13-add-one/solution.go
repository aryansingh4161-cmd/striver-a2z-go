package addone

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// AddOneReverse reverses, adds, and reverses back. Time O(n), space O(1).
func AddOneReverse(head *Node) *Node {
	head = reverse(head)
	carry := 1
	cur := head
	var last *Node
	for cur != nil && carry > 0 {
		v := cur.Val + carry
		cur.Val, carry = v%10, v/10
		last, cur = cur, cur.Next
	}
	if carry > 0 {
		last.Next = &Node{Val: carry}
	}
	return reverse(head)
}

// AddOneRecursive carries back on the way up. Time O(n), space O(n) stack.
func AddOneRecursive(head *Node) *Node {
	var f func(n *Node) int
	f = func(n *Node) int {
		if n == nil {
			return 1 // the +1
		}
		v := n.Val + f(n.Next)
		n.Val = v % 10
		return v / 10
	}
	if c := f(head); c > 0 {
		return &Node{Val: c, Next: head}
	}
	return head
}

func reverse(h *Node) *Node {
	var prev *Node
	for h != nil {
		next := h.Next
		h.Next = prev
		prev, h = h, next
	}
	return prev
}
