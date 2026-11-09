package deletemiddle

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// DeleteMiddle removes node n/2 (0-based). Time O(n), space O(1).
func DeleteMiddle(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return head
}
