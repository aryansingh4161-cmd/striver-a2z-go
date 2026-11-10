package burntree

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestBurnTime(t *testing.T) {
	N := tree.Null
	r := tree.FromLevel([]int{1, 2, 3, 4, N, 5, 6})
	for _, tc := range [][2]int{{1, 2}, {4, 4}, {6, 4}, {3, 3}} {
		if got := BurnTime(r, tree.Find(r, tc[0])); got != tc[1] {
			t.Errorf("start %d: %d", tc[0], got)
		}
	}
	single := tree.FromLevel([]int{9})
	if BurnTime(single, single) != 0 {
		t.Error("single node")
	}
}
