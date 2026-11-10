package bellmanford

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.4-shortest-paths/03-dijkstra"
)

func TestBellmanFord(t *testing.T) {
	edges := [][3]int{{3, 2, 6}, {5, 3, 1}, {0, 1, 5}, {1, 5, -3}, {1, 2, -2}, {3, 4, -2}, {2, 4, 3}}
	dist, neg := BellmanFord(6, edges, 0)
	if neg || !slices.Equal(dist, []int{0, 5, 3, 3, 1, 2}) {
		t.Errorf("got %v %v", dist, neg)
	}
	if _, neg := BellmanFord(3, [][3]int{{0, 1, 1}, {1, 2, -1}, {2, 0, -1}}, 0); !neg {
		t.Error("negative cycle not found")
	}
	// a negative cycle that src can't reach doesn't matter
	if d, neg := BellmanFord(3, [][3]int{{1, 2, -1}, {2, 1, -1}}, 0); neg || d[1] != Inf {
		t.Error("unreachable cycle should be ignored")
	}
}

// With non-negative weights it must agree with Dijkstra.
func TestAgainstDijkstra(t *testing.T) {
	r := rand.New(rand.NewSource(11))
	for it := 0; it < 300; it++ {
		n := 1 + r.Intn(8)
		var edges [][3]int
		adj := make([][]dijkstra.Edge, n)
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if u != v && r.Intn(3) == 0 {
					w := r.Intn(10)
					edges = append(edges, [3]int{u, v, w})
					adj[u] = append(adj[u], dijkstra.Edge{To: v, W: w})
				}
			}
		}
		got, _ := BellmanFord(n, edges, 0)
		want, _ := dijkstra.Dijkstra(adj, 0)
		if !slices.Equal(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	}
}
