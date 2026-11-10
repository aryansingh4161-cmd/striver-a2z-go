package nodesatdistancek

import (
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-13-binary-trees/tree"
)

func TestDistanceK(t *testing.T) {
	N := tree.Null
	r := tree.FromLevel([]int{3, 5, 1, 6, 2, 0, 8, N, N, 7, 4})
	if got := DistanceK(r, tree.Find(r, 5), 2); !slices.Equal(got, []int{1, 4, 7}) {
		t.Errorf("got %v", got)
	}
	if got := DistanceK(r, tree.Find(r, 5), 0); !slices.Equal(got, []int{5}) {
		t.Errorf("k=0: %v", got)
	}
	if got := DistanceK(r, tree.Find(r, 7), 3); !slices.Equal(got, []int{3, 6}) {
		t.Errorf("from 7: %v", got)
	}
}
