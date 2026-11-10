package alltraversalsonepass

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

type frame struct {
	n     *Node
	state int
}

// AllTraversals returns preorder, inorder and postorder in one pass. O(n).
func AllTraversals(root *Node) (pre, in, post []int) {
	pre, in, post = []int{}, []int{}, []int{}
	if root == nil {
		return
	}
	st := []frame{{root, 1}}
	for len(st) > 0 {
		top := &st[len(st)-1]
		switch top.state {
		case 1:
			pre = append(pre, top.n.Val)
			top.state = 2
			if top.n.Left != nil {
				st = append(st, frame{top.n.Left, 1})
			}
		case 2:
			in = append(in, top.n.Val)
			top.state = 3
			if top.n.Right != nil {
				st = append(st, frame{top.n.Right, 1})
			}
		default:
			post = append(post, top.n.Val)
			st = st[:len(st)-1]
		}
	}
	return
}
