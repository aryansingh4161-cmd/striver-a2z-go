package sort012list

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-06-linked-list/list"

type Node = list.Node

// SortCount counts and overwrites values. Time O(2n), space O(1).
func SortCount(head *Node) *Node {
	var c [3]int
	for cur := head; cur != nil; cur = cur.Next {
		c[cur.Val]++
	}
	v := 0
	for cur := head; cur != nil; cur = cur.Next {
		for c[v] == 0 {
			v++
		}
		cur.Val = v
		c[v]--
	}
	return head
}

// Sort re-links nodes into three chains. Time O(n), space O(1).
func Sort(head *Node) *Node {
	zeroD, oneD, twoD := &Node{}, &Node{}, &Node{}
	z, o, w := zeroD, oneD, twoD
	for cur := head; cur != nil; cur = cur.Next {
		switch cur.Val {
		case 0:
			z.Next, z = cur, cur
		case 1:
			o.Next, o = cur, cur
		default:
			w.Next, w = cur, cur
		}
	}
	w.Next = nil
	o.Next = twoD.Next
	if oneD.Next != nil {
		z.Next = oneD.Next
	} else {
		z.Next = twoD.Next
	}
	return zeroD.Next
}
