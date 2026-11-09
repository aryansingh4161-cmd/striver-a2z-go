// Package list holds the node types and small helpers shared by every
// linked-list problem in Step 6. Helpers convert between slices and lists so
// tests stay short and readable.
package list

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// FromSlice builds a singly linked list from a and returns its head (nil if a is empty).
func FromSlice(a []int) *Node {
	dummy := &Node{}
	tail := dummy
	for _, v := range a {
		tail.Next = &Node{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

// ToSlice returns the values of the list in order. It stops after 10_000
// nodes so a buggy cycle can't hang a test.
func ToSlice(head *Node) []int {
	out := []int{}
	for cur := head; cur != nil && len(out) < 10_000; cur = cur.Next {
		out = append(out, cur.Val)
	}
	return out
}

// DNode is a doubly linked list node.
type DNode struct {
	Val        int
	Prev, Next *DNode
}

// DFromSlice builds a doubly linked list from a.
func DFromSlice(a []int) *DNode {
	var head, tail *DNode
	for _, v := range a {
		n := &DNode{Val: v, Prev: tail}
		if tail == nil {
			head = n
		} else {
			tail.Next = n
		}
		tail = n
	}
	return head
}

// DToSlice returns the values in order. It also checks that every Prev
// pointer is correct and returns ok=false if one is wrong.
func DToSlice(head *DNode) (vals []int, ok bool) {
	vals = []int{}
	var prev *DNode
	for cur := head; cur != nil && len(vals) < 10_000; cur = cur.Next {
		if cur.Prev != prev {
			return vals, false
		}
		vals = append(vals, cur.Val)
		prev = cur
	}
	return vals, true
}

// WithCycle builds a list from a and links the last node back to the node at
// index pos (0-based). pos = -1 means no cycle. Used to test loop problems.
func WithCycle(a []int, pos int) *Node {
	head := FromSlice(a)
	if pos < 0 || head == nil {
		return head
	}
	var target, tail *Node
	i := 0
	for cur := head; cur != nil; cur = cur.Next {
		if i == pos {
			target = cur
		}
		tail = cur
		i++
	}
	tail.Next = target
	return head
}

// NodeAt returns the node at index i (0-based), following Next pointers.
func NodeAt(head *Node, i int) *Node {
	cur := head
	for ; i > 0 && cur != nil; i-- {
		cur = cur.Next
	}
	return cur
}
