package postorderonestack

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// Postorder returns Left, Right, Root using one stack. O(n), space O(h).
func Postorder(root *Node) []int {
	out := []int{}
	st := []*Node{}
	var last *Node
	cur := root
	for cur != nil || len(st) > 0 {
		if cur != nil {
			st = append(st, cur)
			cur = cur.Left
			continue
		}
		top := st[len(st)-1]
		if top.Right != nil && top.Right != last {
			cur = top.Right
		} else {
			out = append(out, top.Val)
			last = top
			st = st[:len(st)-1]
		}
	}
	return out
}
