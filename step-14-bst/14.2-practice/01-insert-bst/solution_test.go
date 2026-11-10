package insertbst

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func TestInsert(t *testing.T) {
	r := bst.Build([]int{4, 2, 7, 1, 3})
	r = Insert(r, 5)
	if got := bst.Inorder(r); !slices.Equal(got, []int{1, 2, 3, 4, 5, 7}) {
		t.Errorf("got %v", got)
	}
	if Insert(nil, 9).Val != 9 {
		t.Error("empty")
	}
}
