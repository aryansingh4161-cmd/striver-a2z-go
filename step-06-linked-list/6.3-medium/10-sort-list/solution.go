package sortlist

import (
	"sort"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"
)

type Node = list.Node

// SortCopy copies values to a slice, sorts, writes back. Space O(n).
func SortCopy(head *Node) *Node {
	var vals []int
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	sort.Ints(vals)
	i := 0
	for cur := head; cur != nil; cur = cur.Next {
		cur.Val = vals[i]
		i++
	}
	return head
}

// Sort runs merge sort on the list. Time O(n log n), space O(log n).
func Sort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	// first middle
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := slow.Next
	slow.Next = nil // cut
	return Merge(Sort(head), Sort(right))
}

// Merge merges two sorted lists (LeetCode 21). Time O(n+m), space O(1).
func Merge(a, b *Node) *Node {
	dummy := &Node{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next, a = a, a.Next
		} else {
			tail.Next, b = b, b.Next
		}
		tail = tail.Next
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}
