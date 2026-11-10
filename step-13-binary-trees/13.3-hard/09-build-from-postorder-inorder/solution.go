package buildfrompostorderinorder

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// BuildTree rebuilds a tree from inorder and postorder (unique values). O(n).
func BuildTree(in, post []int) *Node {
	idx := make(map[int]int, len(in))
	for i, v := range in {
		idx[v] = i
	}
	var build func(ps, pe, is, ie int) *Node
	build = func(ps, pe, is, ie int) *Node {
		if ps > pe {
			return nil
		}
		root := &Node{Val: post[pe]}
		k := idx[post[pe]]
		left := k - is
		root.Left = build(ps, ps+left-1, is, k-1)
		root.Right = build(ps+left, pe-1, k+1, ie)
		return root
	}
	return build(0, len(post)-1, 0, len(in)-1)
}
