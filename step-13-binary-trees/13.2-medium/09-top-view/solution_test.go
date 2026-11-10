package topview

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestTopView(t *testing.T) {
	N := tree.Null
	if got := TopView(tree.FromLevel([]int{1, 2, 3, N, 4, N, N, N, 5})); !slices.Equal(got, []int{2, 1, 3}) {
		t.Errorf("got %v", got)
	}
	if got := TopView(tree.FromLevel([]int{1, 2, 3, 4, 5, 6, 7})); !slices.Equal(got, []int{4, 2, 1, 3, 7}) {
		t.Errorf("got %v", got)
	}
}
