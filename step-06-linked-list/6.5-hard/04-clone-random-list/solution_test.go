package clonerandomlist

import "testing"

// build makes nodes with values vals; rand[i] is the index of node i's Random (-1 = nil).
func build(vals, rand []int) []*Node {
	nodes := make([]*Node, len(vals))
	for i, v := range vals {
		nodes[i] = &Node{Val: v}
	}
	for i := range nodes {
		if i+1 < len(nodes) {
			nodes[i].Next = nodes[i+1]
		}
		if rand[i] >= 0 {
			nodes[i].Random = nodes[rand[i]]
		}
	}
	return nodes
}

func check(t *testing.T, name string, orig []*Node, vals, rand []int, clone *Node) {
	t.Helper()
	old := map[*Node]bool{}
	for _, n := range orig {
		old[n] = true
	}
	var copies []*Node
	for c := clone; c != nil; c = c.Next {
		if old[c] {
			t.Fatalf("%s: clone shares a node with the original", name)
		}
		copies = append(copies, c)
	}
	if len(copies) != len(vals) {
		t.Fatalf("%s: length %d", name, len(copies))
	}
	for i, c := range copies {
		if c.Val != vals[i] {
			t.Errorf("%s: val[%d] = %d", name, i, c.Val)
		}
		if rand[i] < 0 && c.Random != nil || rand[i] >= 0 && c.Random != copies[rand[i]] {
			t.Errorf("%s: random[%d] wrong", name, i)
		}
	}
	// original must be untouched
	for i, n := range orig {
		if i+1 < len(orig) && n.Next != orig[i+1] || i+1 == len(orig) && n.Next != nil {
			t.Errorf("%s: original Next changed at %d", name, i)
		}
	}
}

func TestClone(t *testing.T) {
	vals := []int{7, 13, 11, 10, 1}
	rand := []int{-1, 0, 4, 2, 0}
	for name, fn := range map[string]func(*Node) *Node{"map": CloneMap, "interleave": Clone} {
		orig := build(vals, rand)
		check(t, name, orig, vals, rand, fn(orig[0]))
	}
	if Clone(nil) != nil || CloneMap(nil) != nil {
		t.Error("nil")
	}
}
