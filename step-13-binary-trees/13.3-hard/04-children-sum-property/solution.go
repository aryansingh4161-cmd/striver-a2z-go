package childrensumproperty

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// Fix makes every non-leaf equal the sum of its children, only increasing
// values. O(n).
func Fix(n *Node) {
	if n == nil {
		return
	}
	childSum := 0
	if n.Left != nil {
		childSum += n.Left.Val
	}
	if n.Right != nil {
		childSum += n.Right.Val
	}
	if childSum < n.Val {
		if n.Left != nil {
			n.Left.Val = n.Val
		}
		if n.Right != nil {
			n.Right.Val = n.Val
		}
	}
	Fix(n.Left)
	Fix(n.Right)
	if n.Left != nil || n.Right != nil {
		total := 0
		if n.Left != nil {
			total += n.Left.Val
		}
		if n.Right != nil {
			total += n.Right.Val
		}
		n.Val = total
	}
}
