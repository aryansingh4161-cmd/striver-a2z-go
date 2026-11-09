package reversekgroup

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// ReverseKGroup reverses every full group of k nodes. Time O(n), space O(1).
func ReverseKGroup(head *Node, k int) *Node {
	dummy := &Node{Next: head}
	groupPrev := dummy
	for {
		kth := groupPrev
		for i := 0; i < k && kth != nil; i++ {
			kth = kth.Next
		}
		if kth == nil {
			break // fewer than k nodes left
		}
		groupNext := kth.Next
		// reverse the group; its new tail should point to groupNext
		prev, cur := groupNext, groupPrev.Next
		for cur != groupNext {
			next := cur.Next
			cur.Next = prev
			prev, cur = cur, next
		}
		oldFirst := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = oldFirst
	}
	return dummy.Next
}
