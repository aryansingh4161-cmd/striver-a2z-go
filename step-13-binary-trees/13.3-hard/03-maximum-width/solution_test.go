package maximumwidth

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestWidth(t *testing.T) {
	N := tree.Null
	tests := []struct {
		lvl  []int
		want int
	}{{[]int{1, 3, 2, 5, 3, N, 9}, 4}, {[]int{1, 3, 2, 5, N, N, 9, 6, N, 7}, 7}, {[]int{1, 3, 2, 5}, 2}, {[]int{1}, 1}}
	for _, tc := range tests {
		if got := WidthOfBinaryTree(tree.FromLevel(tc.lvl)); got != tc.want {
			t.Errorf("Width(%v) = %d", tc.lvl, got)
		}
	}
	// very deep zigzag tree: indexes would overflow without normalising
	root := &Node{Val: 0}
	cur := root
	for i := 0; i < 200; i++ {
		if i%2 == 0 {
			cur.Right = &Node{Val: i}
			cur = cur.Right
		} else {
			cur.Left = &Node{Val: i}
			cur = cur.Left
		}
	}
	if WidthOfBinaryTree(root) != 1 {
		t.Error("deep tree")
	}
}
