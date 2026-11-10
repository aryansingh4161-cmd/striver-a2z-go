package deletebst

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func TestDelete(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 100; it++ {
		vals := r.Perm(15)
		root := bst.Build(vals)
		key := r.Intn(18) // sometimes missing
		root = Delete(root, key)
		want := []int{}
		for v := 0; v < 15; v++ {
			if v != key {
				want = append(want, v)
			}
		}
		if got := bst.Inorder(root); !slices.Equal(got, want) {
			t.Fatalf("delete %d from %v: %v", key, vals, got)
		}
	}
	if Delete(bst.Build([]int{1}), 1) != nil {
		t.Error("delete only node")
	}
}
