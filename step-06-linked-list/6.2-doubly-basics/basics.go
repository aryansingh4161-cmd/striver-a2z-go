// Package doublybasics implements basic doubly linked list operations.
package doublybasics

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type DNode = list.DNode

// InsertHead adds v at the front. O(1).
func InsertHead(head *DNode, v int) *DNode {
	n := &DNode{Val: v, Next: head}
	if head != nil {
		head.Prev = n
	}
	return n
}

// InsertTail adds v at the end. O(n).
func InsertTail(head *DNode, v int) *DNode {
	if head == nil {
		return &DNode{Val: v}
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &DNode{Val: v, Prev: cur}
	return head
}

// InsertBeforeK inserts v before the k-th node (1-based). O(k).
func InsertBeforeK(head *DNode, v, k int) *DNode {
	if k == 1 {
		return InsertHead(head, v)
	}
	cur := head
	for i := 1; cur != nil && i < k; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return head
	}
	prev := cur.Prev
	n := &DNode{Val: v, Prev: prev, Next: cur}
	prev.Next = n
	cur.Prev = n
	return head
}

// DeleteHead removes the first node. O(1).
func DeleteHead(head *DNode) *DNode {
	if head == nil || head.Next == nil {
		return nil
	}
	newHead := head.Next
	newHead.Prev = nil
	head.Next = nil
	return newHead
}

// DeleteTail removes the last node. O(n).
func DeleteTail(head *DNode) *DNode {
	if head == nil || head.Next == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Prev.Next = nil
	cur.Prev = nil
	return head
}

// DeleteK removes the k-th node (1-based). O(k).
func DeleteK(head *DNode, k int) *DNode {
	cur := head
	for i := 1; cur != nil && i < k; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return head
	}
	return deleteNode(head, cur)
}

// deleteNode unlinks n from the list and returns the (possibly new) head. O(1).
func deleteNode(head, n *DNode) *DNode {
	if n.Prev != nil {
		n.Prev.Next = n.Next
	} else {
		head = n.Next
	}
	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
	n.Prev, n.Next = nil, nil
	return head
}

// Reverse reverses the list by swapping Prev and Next of every node.
// Time O(n), space O(1).
func Reverse(head *DNode) *DNode {
	var last *DNode
	for cur := head; cur != nil; cur = cur.Prev {
		cur.Prev, cur.Next = cur.Next, cur.Prev
		last = cur
	}
	return last
}
