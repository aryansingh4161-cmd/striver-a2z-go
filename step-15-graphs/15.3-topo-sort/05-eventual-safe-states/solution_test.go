package eventualsafestates

import (
	"math/rand"
	"slices"
	"testing"
)

func TestSafeNodes(t *testing.T) {
	g := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	want := []int{2, 4, 5, 6}
	if got := SafeNodesDFS(g); !slices.Equal(got, want) {
		t.Errorf("dfs %v", got)
	}
	if got := SafeNodesKahn(g); !slices.Equal(got, want) {
		t.Errorf("kahn %v", got)
	}
	r := rand.New(rand.NewSource(5))
	for it := 0; it < 500; it++ {
		n := 1 + r.Intn(7)
		g := make([][]int, n)
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if r.Intn(4) == 0 {
					g[u] = append(g[u], v)
				}
			}
		}
		if a, b := SafeNodesDFS(g), SafeNodesKahn(g); !slices.Equal(a, b) {
			t.Fatalf("%v: dfs %v kahn %v", g, a, b)
		}
	}
}
