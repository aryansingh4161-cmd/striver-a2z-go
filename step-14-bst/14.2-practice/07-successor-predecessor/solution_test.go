package successorpredecessor

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func TestSuccPred(t *testing.T) {
	r := bst.Build([]int{5, 3, 8, 1, 4, 7, 9, 2, 6})
	for x := 0; x <= 10; x++ {
		ws, wp := -1, -1
		for v := 1; v <= 9; v++ {
			if v > x && ws == -1 {
				ws = v
			}
			if v < x {
				wp = v
			}
		}
		if Successor(r, x) != ws || Predecessor(r, x) != wp {
			t.Errorf("x=%d", x)
		}
	}
}
