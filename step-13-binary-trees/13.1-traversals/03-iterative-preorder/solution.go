package iterativepreorder

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// Preorder returns Root, Left, Right using an explicit stack. O(n).
func Preorder(root *Node) []int {
	out := []int{}
	if root == nil {
		return out
	}
	st := []*Node{root}
	for len(st) > 0 {
		n := st[len(st)-1]
		st = st[:len(st)-1]
		out = append(out, n.Val)
		if n.Right != nil {
			st = append(st, n.Right)
		}
		if n.Left != nil {
			st = append(st, n.Left)
		}
	}
	return out
}
