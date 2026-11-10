package bstfrompreorder

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestBstFromPreorder(t *testing.T) {
	N := tree.Null
	if got := tree.ToLevel(BstFromPreorder([]int{8, 5, 1, 7, 10, 12})); !slices.Equal(got, []int{8, 5, 10, 1, 7, N, 12}) {
		t.Errorf("got %v", got)
	}
	if got := tree.ToLevel(BstFromPreorder([]int{1, 3})); !slices.Equal(got, []int{1, N, 3}) {
		t.Errorf("got %v", got)
	}
}
