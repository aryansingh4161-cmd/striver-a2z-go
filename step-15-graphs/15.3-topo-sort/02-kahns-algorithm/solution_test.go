package kahnsalgorithm

import (
	"math/rand"
	"testing"

	toposortdfs "github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.3-topo-sort/01-topo-sort-dfs"
)

func TestTopoSort(t *testing.T) {
	adj := [][]int{{}, {}, {3}, {1}, {0, 1}, {0, 2}} // 5→0 listed before 5→2, as in the README
	got := TopoSort(adj)
	want := []int{4, 5, 0, 2, 3, 1}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("got %v want %v", got, want)
		}
	}
	r := rand.New(rand.NewSource(2))
	for it := 0; it < 200; it++ {
		n := 1 + r.Intn(9)
		perm := r.Perm(n)
		g := make([][]int, n)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if r.Intn(3) == 0 {
					g[perm[i]] = append(g[perm[i]], perm[j])
				}
			}
		}
		if o := TopoSort(g); !toposortdfs.IsTopoOrder(g, o) {
			t.Fatalf("invalid %v", o)
		}
	}
}

func TestCycleGivesShortOrder(t *testing.T) {
	if len(TopoSort([][]int{{1}, {2}, {0}})) == 3 {
		t.Error("cycle should leave nodes out")
	}
}
