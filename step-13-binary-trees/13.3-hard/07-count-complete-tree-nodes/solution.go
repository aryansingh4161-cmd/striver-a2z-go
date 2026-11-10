package countcompletetreenodes

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// CountNodes counts nodes of a complete tree in O(log^2 n).
func CountNodes(n *Node) int {
	if n == nil {
		return 0
	}
	lh, rh := 0, 0
	for c := n; c != nil; c = c.Left {
		lh++
	}
	for c := n; c != nil; c = c.Right {
		rh++
	}
	if lh == rh {
		return 1<<lh - 1 // perfect subtree
	}
	return 1 + CountNodes(n.Left) + CountNodes(n.Right)
}
