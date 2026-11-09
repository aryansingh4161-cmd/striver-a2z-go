package deletealloccurrences

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type DNode = list.DNode

// DeleteAll removes every node with value key. Time O(n), space O(1).
func DeleteAll(head *DNode, key int) *DNode {
	cur := head
	for cur != nil {
		next := cur.Next
		if cur.Val == key {
			if cur == head {
				head = next
			}
			if cur.Prev != nil {
				cur.Prev.Next = next
			}
			if next != nil {
				next.Prev = cur.Prev
			}
			cur.Prev, cur.Next = nil, nil
		}
		cur = next
	}
	return head
}
