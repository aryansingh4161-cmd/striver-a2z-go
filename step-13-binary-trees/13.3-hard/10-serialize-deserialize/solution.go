package serializedeserialize

import (
	"strconv"
	"strings"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

type Node = tree.Node

// Serialize writes the tree in level order with "#" for nil. O(n).
func Serialize(root *Node) string {
	if root == nil {
		return ""
	}
	var parts []string
	q := []*Node{root}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n == nil {
			parts = append(parts, "#")
			continue
		}
		parts = append(parts, strconv.Itoa(n.Val))
		q = append(q, n.Left, n.Right)
	}
	return strings.Join(parts, ",")
}

// Deserialize rebuilds the tree written by Serialize. O(n).
func Deserialize(s string) *Node {
	if s == "" {
		return nil
	}
	tok := strings.Split(s, ",")
	v, _ := strconv.Atoi(tok[0])
	root := &Node{Val: v}
	q := []*Node{root}
	i := 1
	for len(q) > 0 && i < len(tok) {
		n := q[0]
		q = q[1:]
		if tok[i] != "#" {
			v, _ := strconv.Atoi(tok[i])
			n.Left = &Node{Val: v}
			q = append(q, n.Left)
		}
		i++
		if i < len(tok) && tok[i] != "#" {
			v, _ := strconv.Atoi(tok[i])
			n.Right = &Node{Val: v}
			q = append(q, n.Right)
		}
		i++
	}
	return root
}

// SerializePre writes preorder with "#" for nil. O(n).
func SerializePre(root *Node) string {
	var parts []string
	var f func(*Node)
	f = func(n *Node) {
		if n == nil {
			parts = append(parts, "#")
			return
		}
		parts = append(parts, strconv.Itoa(n.Val))
		f(n.Left)
		f(n.Right)
	}
	f(root)
	return strings.Join(parts, ",")
}

// DeserializePre rebuilds a tree from SerializePre output. O(n).
func DeserializePre(s string) *Node {
	tok := strings.Split(s, ",")
	i := 0
	var f func() *Node
	f = func() *Node {
		t := tok[i]
		i++
		if t == "#" {
			return nil
		}
		v, _ := strconv.Atoi(t)
		n := &Node{Val: v}
		n.Left = f()
		n.Right = f()
		return n
	}
	return f()
}
