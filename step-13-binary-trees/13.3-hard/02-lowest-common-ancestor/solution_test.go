package lowestcommonancestor

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestLCA(t *testing.T) {
	N := tree.Null
	r := tree.FromLevel([]int{3, 5, 1, 6, 2, 0, 8, N, N, 7, 4})
	tests := [][3]int{{5, 1, 3}, {5, 4, 5}, {7, 4, 2}, {6, 4, 5}, {0, 8, 1}}
	for _, tc := range tests {
		got := LCA(r, tree.Find(r, tc[0]), tree.Find(r, tc[1]))
		if got == nil || got.Val != tc[2] {
			t.Errorf("LCA(%d,%d) wrong", tc[0], tc[1])
		}
	}
}
