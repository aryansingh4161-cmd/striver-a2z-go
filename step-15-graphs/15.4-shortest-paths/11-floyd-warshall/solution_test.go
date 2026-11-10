package floydwarshall

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/aryansingh4161-cmd/striver-a2z-go/step-15-graphs/15.4-shortest-paths/03-dijkstra"
)

func TestFloyd(t *testing.T) {
	m := [][]int{{0, 2, -1, -1}, {1, 0, 3, -1}, {-1, -1, 0, -1}, {3, 5, 4, 0}}
	if ShortestDistances(m) {
		t.Error("no negative cycle here")
	}
	want := [][]int{{0, 2, 5, -1}, {1, 0, 3, -1}, {-1, -1, 0, -1}, {3, 5, 4, 0}}
	for i := range want {
		if !slices.Equal(m[i], want[i]) {
			t.Fatalf("row %d got %v want %v", i, m[i], want[i])
		}
	}
	neg := [][]int{{0, 1, -1}, {-1, 0, -1}, {-3, -1, 0}}
	neg[1][2] = 1 // 0→1→2→0 costs 1 + 1 − 3 = −1
	if !ShortestDistances(neg) {
		t.Error("negative cycle not found")
	}
}

// Must agree with Dijkstra from every node when weights are non-negative.
func TestAgainstDijkstra(t *testing.T) {
	r := rand.New(rand.NewSource(12))
	for it := 0; it < 200; it++ {
		n := 1 + r.Intn(7)
		m := make([][]int, n)
		adj := make([][]dijkstra.Edge, n)
		for i := range m {
			m[i] = make([]int, n)
			for j := range m[i] {
				m[i][j] = -1
				if i == j {
					m[i][j] = 0
				} else if r.Intn(3) == 0 {
					w := r.Intn(9)
					m[i][j] = w
					adj[i] = append(adj[i], dijkstra.Edge{To: j, W: w})
				}
			}
		}
		ShortestDistances(m)
		for s := 0; s < n; s++ {
			d, _ := dijkstra.Dijkstra(adj, s)
			for j := range d {
				if d[j] == dijkstra.Inf {
					d[j] = -1
				}
			}
			if !slices.Equal(d, m[s]) {
				t.Fatalf("source %d: floyd %v dijkstra %v", s, m[s], d)
			}
		}
	}
}
