package flattenlist

import "sort"

// Node has a Next pointer (top level) and a Child pointer (vertical sorted list).
type Node struct {
	Val         int
	Next, Child *Node
}

// FlattenSort collects all values, sorts and rebuilds. Time O(N log N), space O(N).
func FlattenSort(head *Node) *Node {
	var vals []int
	for col := head; col != nil; col = col.Next {
		for n := col; n != nil; n = n.Child {
			vals = append(vals, n.Val)
		}
	}
	sort.Ints(vals)
	dummy := &Node{}
	tail := dummy
	for _, v := range vals {
		tail.Child = &Node{Val: v}
		tail = tail.Child
	}
	return dummy.Child
}

// Flatten merges vertical lists from the right. Time O(N*k), space O(k).
func Flatten(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	right := Flatten(head.Next)
	head.Next = nil
	return merge(head, right)
}

func merge(a, b *Node) *Node {
	dummy := &Node{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Child, a = a, a.Child
		} else {
			tail.Child, b = b, b.Child
		}
		tail = tail.Child
		tail.Next = nil
	}
	if a != nil {
		tail.Child = a
	} else {
		tail.Child = b
	}
	return dummy.Child
}
