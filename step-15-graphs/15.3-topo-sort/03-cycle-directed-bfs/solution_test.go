package cycledirectedbfs

import (
	"math/rand"
	"testing"

	cycledirecteddfs "github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.2-bfs-dfs/13-cycle-directed-dfs"
)

func TestHasCycle(t *testing.T) {
	if !HasCycle([][]int{{1}, {2}, {0, 3}, {}}) {
		t.Error("0→1→2→0 is a cycle")
	}
	if HasCycle([][]int{{1, 2}, {}, {1}}) {
		t.Error("diamond has no cycle")
	}
	if !HasCycle([][]int{{0}}) {
		t.Error("self loop is a cycle")
	}
}

// Compare BFS and DFS versions on random graphs.
func TestAgainstDFS(t *testing.T) {
	r := rand.New(rand.NewSource(3))
	for it := 0; it < 500; it++ {
		n := 1 + r.Intn(7)
		g := make([][]int, n)
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if r.Intn(5) == 0 {
					g[u] = append(g[u], v)
				}
			}
		}
		if HasCycle(g) != cycledirecteddfs.HasCycle(g) {
			t.Fatalf("disagree on %v", g)
		}
	}
}
