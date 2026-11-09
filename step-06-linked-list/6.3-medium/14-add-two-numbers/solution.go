package addtwonumbers

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// AddTwo adds two reversed-digit lists. Time O(max(n,m)).
func AddTwo(a, b *Node) *Node {
	dummy := &Node{}
	tail := dummy
	carry := 0
	for a != nil || b != nil || carry != 0 {
		sum := carry
		if a != nil {
			sum += a.Val
			a = a.Next
		}
		if b != nil {
			sum += b.Val
			b = b.Next
		}
		tail.Next = &Node{Val: sum % 10}
		tail = tail.Next
		carry = sum / 10
	}
	return dummy.Next
}
