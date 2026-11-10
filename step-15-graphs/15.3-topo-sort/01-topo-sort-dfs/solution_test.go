package toposortdfs

import (
	"math/rand"
	"testing"
)

// randomDAG only adds edges from a smaller to a larger label of a shuffled
// permutation, so it can never contain a cycle.
func randomDAG(r *rand.Rand, n int) [][]int {
	perm := r.Perm(n)
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if r.Intn(3) == 0 {
				adj[perm[i]] = append(adj[perm[i]], perm[j])
			}
		}
	}
	return adj
}

func TestTopoSort(t *testing.T) {
	adj := [][]int{{}, {}, {3}, {1}, {0, 1}, {2, 0}}
	if got := TopoSort(adj); !IsTopoOrder(adj, got) {
		t.Errorf("invalid order %v", got)
	}
	if got := TopoSort([][]int{{}}); !IsTopoOrder([][]int{{}}, got) {
		t.Error("single node")
	}
	r := rand.New(rand.NewSource(1))
	for it := 0; it < 200; it++ {
		g := randomDAG(r, 1+r.Intn(9))
		if got := TopoSort(g); !IsTopoOrder(g, got) {
			t.Fatalf("invalid order %v for %v", got, g)
		}
	}
}

func TestIsTopoOrder(t *testing.T) {
	adj := [][]int{{1}, {}}
	if IsTopoOrder(adj, []int{1, 0}) {
		t.Error("1 before 0 breaks edge 0→1")
	}
}
