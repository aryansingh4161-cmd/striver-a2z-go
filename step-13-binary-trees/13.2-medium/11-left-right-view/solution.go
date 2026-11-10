package leftrightview

import "github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"

type Node = tree.Node

// RightView returns the rightmost node of each level. DFS right-first. O(n).
func RightView(root *Node) []int {
	res := []int{}
	var dfs func(n *Node, depth int)
	dfs = func(n *Node, depth int) {
		if n == nil {
			return
		}
		if depth == len(res) {
			res = append(res, n.Val)
		}
		dfs(n.Right, depth+1)
		dfs(n.Left, depth+1)
	}
	dfs(root, 0)
	return res
}

// LeftView returns the leftmost node of each level. DFS left-first. O(n).
func LeftView(root *Node) []int {
	res := []int{}
	var dfs func(n *Node, depth int)
	dfs = func(n *Node, depth int) {
		if n == nil {
			return
		}
		if depth == len(res) {
			res = append(res, n.Val)
		}
		dfs(n.Left, depth+1)
		dfs(n.Right, depth+1)
	}
	dfs(root, 0)
	return res
}
