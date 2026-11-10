package postordertwostacks

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// Postorder returns Left, Right, Root using two stacks. O(n).
func Postorder(root *Node) []int {
	out := []int{}
	if root == nil {
		return out
	}
	s1, s2 := []*Node{root}, []*Node{}
	for len(s1) > 0 {
		n := s1[len(s1)-1]
		s1 = s1[:len(s1)-1]
		s2 = append(s2, n)
		if n.Left != nil {
			s1 = append(s1, n.Left)
		}
		if n.Right != nil {
			s1 = append(s1, n.Right)
		}
	}
	for i := len(s2) - 1; i >= 0; i-- {
		out = append(out, s2[i].Val)
	}
	return out
}
