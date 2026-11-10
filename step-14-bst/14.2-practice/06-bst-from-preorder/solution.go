package bstfrompreorder

import (
	"math"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

type Node = bst.Node

// BstFromPreorder builds the BST in O(n) using an upper bound.
func BstFromPreorder(pre []int) *Node {
	i := 0
	var build func(bound int) *Node
	build = func(bound int) *Node {
		if i == len(pre) || pre[i] > bound {
			return nil
		}
		root := &Node{Val: pre[i]}
		i++
		root.Left = build(root.Val)
		root.Right = build(bound)
		return root
	}
	return build(math.MaxInt)
}
