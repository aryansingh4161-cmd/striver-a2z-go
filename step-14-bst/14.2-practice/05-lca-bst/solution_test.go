package lcabst

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func TestLCA(t *testing.T) {
	r := bst.Build([]int{6, 2, 8, 0, 4, 7, 9, 3, 5})
	for _, tc := range [][3]int{{2, 8, 6}, {2, 4, 2}, {3, 5, 4}, {7, 9, 8}, {0, 5, 2}} {
		if got := LCA(r, tc[0], tc[1]); got == nil || got.Val != tc[2] {
			t.Errorf("LCA(%d,%d)", tc[0], tc[1])
		}
	}
}
