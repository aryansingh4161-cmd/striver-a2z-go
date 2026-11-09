package removenthfromend

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// RemoveNthCount counts the length first (two passes).
func RemoveNthCount(head *Node, n int) *Node {
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	dummy := &Node{Next: head}
	prev := dummy
	for i := 0; i < length-n; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	return dummy.Next
}

// RemoveNth uses two pointers n apart (one pass). Time O(n), space O(1).
func RemoveNth(head *Node, n int) *Node {
	dummy := &Node{Next: head}
	slow, fast := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
