package lengthofloop

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// LoopLengthMap records the step at which each node was first visited. Space O(n).
func LoopLengthMap(head *Node) int {
	first := map[*Node]int{}
	step := 0
	for cur := head; cur != nil; cur = cur.Next {
		if s, ok := first[cur]; ok {
			return step - s
		}
		first[cur] = step
		step++
	}
	return 0
}

// LoopLength finds the meeting point, then counts one lap. Time O(n), space O(1).
func LoopLength(head *Node) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			count := 1
			for cur := slow.Next; cur != slow; cur = cur.Next {
				count++
			}
			return count
		}
	}
	return 0
}
