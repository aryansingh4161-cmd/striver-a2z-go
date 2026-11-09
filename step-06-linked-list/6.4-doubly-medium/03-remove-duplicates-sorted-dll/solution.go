package removeduplicatessorteddll

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type DNode = list.DNode

// RemoveDuplicates keeps one node per value in a sorted DLL. Time O(n), space O(1).
func RemoveDuplicates(head *DNode) *DNode {
	for cur := head; cur != nil; cur = cur.Next {
		nxt := cur.Next
		for nxt != nil && nxt.Val == cur.Val {
			nxt = nxt.Next
		}
		cur.Next = nxt
		if nxt != nil {
			nxt.Prev = cur
		}
	}
	return head
}
