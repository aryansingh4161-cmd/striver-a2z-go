package buildfrompreorderinorder

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestBuildTree(t *testing.T) {
	N := tree.Null
	r := BuildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	if got := tree.ToLevel(r); !slices.Equal(got, []int{3, 9, 20, N, N, 15, 7}) {
		t.Errorf("got %v", got)
	}
	if BuildTree(nil, nil) != nil {
		t.Error("empty")
	}
	r2 := BuildTree([]int{1, 2, 3}, []int{3, 2, 1}) // left chain
	if got := tree.ToLevel(r2); !slices.Equal(got, []int{1, 2, N, 3}) {
		t.Errorf("chain: %v", got)
	}
}
