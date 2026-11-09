package rotatelist

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// RotateRight rotates the list right by k. Time O(n), space O(1).
func RotateRight(head *Node, k int) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	n, tail := 1, head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	k %= n
	if k == 0 {
		return head
	}
	tail.Next = head // make a ring
	newTail := head
	for i := 1; i < n-k; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil // cut
	return newHead
}
