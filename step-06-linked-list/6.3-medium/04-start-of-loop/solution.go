package startofloop

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// DetectCycleSet returns the first repeated node. Space O(n).
func DetectCycleSet(head *Node) *Node {
	seen := map[*Node]bool{}
	for cur := head; cur != nil; cur = cur.Next {
		if seen[cur] {
			return cur
		}
		seen[cur] = true
	}
	return nil
}

// DetectCycle uses Floyd's algorithm. Time O(n), space O(1).
func DetectCycle(head *Node) *Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // phase 2
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
