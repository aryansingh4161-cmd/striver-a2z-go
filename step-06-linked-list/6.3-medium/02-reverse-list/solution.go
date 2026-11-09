package reverselist

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// ReverseStack copies the values to a stack and writes them back. Space O(n).
func ReverseStack(head *Node) *Node {
	var st []int
	for cur := head; cur != nil; cur = cur.Next {
		st = append(st, cur.Val)
	}
	for cur := head; cur != nil; cur = cur.Next {
		cur.Val = st[len(st)-1]
		st = st[:len(st)-1]
	}
	return head
}

// Reverse flips every Next pointer. Time O(n), space O(1).
func Reverse(head *Node) *Node {
	var prev *Node
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}

// ReverseRecursive reverses the rest first, then fixes one link. Space O(n).
func ReverseRecursive(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
