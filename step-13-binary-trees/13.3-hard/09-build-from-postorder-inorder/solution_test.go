package buildfrompostorderinorder

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestBuildTree(t *testing.T) {
	N := tree.Null
	r := BuildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	if got := tree.ToLevel(r); !slices.Equal(got, []int{3, 9, 20, N, N, 15, 7}) {
		t.Errorf("got %v", got)
	}
}
