package recoverbst

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-14-bst/bst"
)

func nodes(n *Node, out *[]*Node) {
	if n == nil {
		return
	}
	nodes(n.Left, out)
	*out = append(*out, n)
	nodes(n.Right, out)
}

func TestRecover(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	for it := 0; it < 200; it++ {
		root := bst.Build(r.Perm(12))
		var ns []*Node
		nodes(root, &ns)
		i, j := r.Intn(12), r.Intn(12)
		for i == j {
			j = r.Intn(12)
		}
		ns[i].Val, ns[j].Val = ns[j].Val, ns[i].Val // break it
		Recover(root)
		want := make([]int, 12)
		for k := range want {
			want[k] = k
		}
		if got := bst.Inorder(root); !slices.Equal(got, want) {
			t.Fatalf("swap %d,%d: %v", i, j, got)
		}
	}
}
