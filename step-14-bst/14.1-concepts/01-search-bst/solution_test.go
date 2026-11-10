package searchbst

import (
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func TestSearch(t *testing.T) {
	r := bst.Build([]int{4, 2, 7, 1, 3})
	if n := Search(r, 2); n == nil || n.Val != 2 || n.Left.Val != 1 {
		t.Error("found 2")
	}
	if Search(r, 5) != nil {
		t.Error("5 missing")
	}
}
