package iterativeinorder

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// Inorder returns Left, Root, Right using an explicit stack. O(n).
func Inorder(root *Node) []int {
	out := []int{}
	st := []*Node{}
	cur := root
	for {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		if len(st) == 0 {
			break
		}
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		out = append(out, cur.Val)
		cur = cur.Right
	}
	return out
}
