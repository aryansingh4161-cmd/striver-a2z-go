package oddevenlist

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// OddEven groups odd-position nodes before even-position nodes in place.
// Time O(n), space O(1).
func OddEven(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	odd, even := head, head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
