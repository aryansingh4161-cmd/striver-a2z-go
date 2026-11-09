// Package singlybasics implements the basic singly linked list operations.
package singlybasics

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// Length counts the nodes. O(n).
func Length(head *Node) int {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	return n
}

// Search reports whether x is in the list. O(n).
func Search(head *Node, x int) bool {
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val == x {
			return true
		}
	}
	return false
}

// InsertHead adds v at the front. O(1).
func InsertHead(head *Node, v int) *Node {
	return &Node{Val: v, Next: head}
}

// InsertTail adds v at the end. O(n).
func InsertTail(head *Node, v int) *Node {
	n := &Node{Val: v}
	if head == nil {
		return n
	}
	cur := head
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = n
	return head
}

// InsertAt inserts v so that it becomes the k-th node (1-based).
// k = length+1 appends. Larger k leaves the list unchanged. O(k).
func InsertAt(head *Node, v, k int) *Node {
	if k == 1 {
		return InsertHead(head, v)
	}
	cur := head
	for i := 1; cur != nil && i < k-1; i++ {
		cur = cur.Next
	}
	if cur == nil {
		return head // k too large
	}
	cur.Next = &Node{Val: v, Next: cur.Next}
	return head
}

// InsertBefore inserts v before the first node with value x (if any). O(n).
func InsertBefore(head *Node, v, x int) *Node {
	if head == nil {
		return nil
	}
	if head.Val == x {
		return InsertHead(head, v)
	}
	for cur := head; cur.Next != nil; cur = cur.Next {
		if cur.Next.Val == x {
			cur.Next = &Node{Val: v, Next: cur.Next}
			break
		}
	}
	return head
}

// DeleteHead removes the first node. O(1).
func DeleteHead(head *Node) *Node {
	if head == nil {
		return nil
	}
	return head.Next
}

// DeleteTail removes the last node. O(n).
func DeleteTail(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	cur := head
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	cur.Next = nil
	return head
}

// DeleteAt removes the k-th node (1-based). O(k).
func DeleteAt(head *Node, k int) *Node {
	if head == nil {
		return nil
	}
	if k == 1 {
		return head.Next
	}
	cur := head
	for i := 1; cur.Next != nil && i < k-1; i++ {
		cur = cur.Next
	}
	if cur.Next != nil {
		cur.Next = cur.Next.Next
	}
	return head
}

// DeleteValue removes the first node with value x. O(n).
func DeleteValue(head *Node, x int) *Node {
	dummy := &Node{Next: head}
	for cur := dummy; cur.Next != nil; cur = cur.Next {
		if cur.Next.Val == x {
			cur.Next = cur.Next.Next
			break
		}
	}
	return dummy.Next
}
