package roottonodepath

import (
	"strconv"
	"strings"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

type Node = tree.Node

// PathTo returns the values from root to the node with value x (nil if absent). O(n).
func PathTo(root *Node, x int) []int {
	var path []int
	var dfs func(*Node) bool
	dfs = func(n *Node) bool {
		if n == nil {
			return false
		}
		path = append(path, n.Val)
		if n.Val == x || dfs(n.Left) || dfs(n.Right) {
			return true
		}
		path = path[:len(path)-1]
		return false
	}
	if dfs(root) {
		return path
	}
	return nil
}

// AllRootToLeafPaths returns every root-to-leaf path as "1->2->5". O(n * h).
func AllRootToLeafPaths(root *Node) []string {
	res := []string{}
	var path []string
	var dfs func(*Node)
	dfs = func(n *Node) {
		if n == nil {
			return
		}
		path = append(path, strconv.Itoa(n.Val))
		if n.Left == nil && n.Right == nil {
			res = append(res, strings.Join(path, "->"))
		}
		dfs(n.Left)
		dfs(n.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return res
}
