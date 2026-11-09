package detectloop

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// HasCycleSet remembers visited nodes. Time O(n), space O(n).
func HasCycleSet(head *Node) bool {
	seen := map[*Node]bool{}
	for cur := head; cur != nil; cur = cur.Next {
		if seen[cur] {
			return true
		}
		seen[cur] = true
	}
	return false
}

// HasCycle uses Floyd's slow/fast pointers. Time O(n), space O(1).
func HasCycle(head *Node) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
