package buildfrompreorderinorder

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// BuildTree rebuilds a tree from preorder and inorder (unique values). O(n).
func BuildTree(pre, in []int) *Node {
	idx := make(map[int]int, len(in))
	for i, v := range in {
		idx[v] = i
	}
	var build func(ps, pe, is, ie int) *Node
	build = func(ps, pe, is, ie int) *Node {
		if ps > pe {
			return nil
		}
		root := &Node{Val: pre[ps]}
		k := idx[pre[ps]]
		left := k - is
		root.Left = build(ps+1, ps+left, is, k-1)
		root.Right = build(ps+left+1, pe, k+1, ie)
		return root
	}
	return build(0, len(pre)-1, 0, len(in)-1)
}
