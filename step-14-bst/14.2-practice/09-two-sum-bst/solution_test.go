package twosumbst

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func TestFindTarget(t *testing.T) {
	r := bst.Build([]int{5, 3, 6, 2, 4, 7})
	if !FindTarget(r, 9) || FindTarget(r, 28) || !FindTarget(r, 13) || FindTarget(r, 4) {
		t.Error("examples")
	}
	if FindTarget(bst.Build([]int{1}), 2) {
		t.Error("same node twice")
	}
}
